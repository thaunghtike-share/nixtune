/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

// MachineStats is the data that we receive from the agent that we
// will put into S3 and a Metadata table.
type MachineStats struct {
	Username  string
	Name      string
	Stats     []byte
	Timestamp int64
}

// ID represents the prefix where we will store information in the S3 bucket.
//
// <username>/machine/<name>/<timestamp>
func (s *MachineStats) ID() string {
	return fmt.Sprintf("%s/machine/%s/%d", s.Username, s.Name, s.Timestamp)
}

func (s *MachineStats) s3Key() string {
	return fmt.Sprintf(filepath.Join(s.ID(), "stats.json"))
}

// Create write the stats data from the agent to S3 and if that was
// successful to a metadata table in a DB.
func (s *MachineStats) Create() error {
	s.Timestamp = time.Now().UTC().Unix()

	// Backup Write to S3
	params := &s3.PutObjectInput{
		Bucket: aws.String(acksinBucket()),
		Key:    aws.String(s.s3Key()),
		Body:   bytes.NewReader(s.Stats),
	}

	_, err := s3svc().PutObject(params)
	if err != nil {
		return err
	}

	userDB().QueryRow("INSERT INTO acksin_machines(username, name, id) VALUES ($1, $2, $3) RETURNING id", s.Username, s.Name, s.ID())

	return nil
}

func (s *MachineStats) RunModels() {
	b, _ := json.Marshal(struct {
		Machine bool
		ID      string
	}{true, s.ID()})

	params := &lambda.InvokeInput{
		FunctionName: aws.String(mentalModelFunction()),
		Payload:      b,
	}

	resp, err := lambdasvc().Invoke(params)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(resp.Payload))
}

func PostStatsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		username = authUsername(r)
	)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		respondJSON(w, errorResponse{"Failed to parse stats.", 500})
		return
	}

	machineName := r.Header.Get("X-Acksin-MachineName")
	if machineName == "" {
		machineName = "noname"
	}

	s := &MachineStats{
		Username: username,
		Name:     machineName,
		Stats:    b,
	}

	// POST responses should be fast. Offload the processing.
	go func(stats *MachineStats) {
		if err := stats.Create(); err != nil {
			log.Println(err)
			return
		}

		stats.RunModels()
	}(s)

	respondJSON(w, struct{}{})
}
