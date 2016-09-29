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

type MachineStats struct {
	Username  string
	Name      string
	Stats     []byte
	Timestamp int64
}

// /<username>/machine/<name>/<timestamp>
func (s *MachineStats) ID() string {
	return fmt.Sprintf("%s/machine/%s/%d", s.Username, s.Name, s.Timestamp)
}

func (s *MachineStats) s3Key() string {
	return fmt.Sprintf(filepath.Join(s.ID(), "stats.json"))
}

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

	return nil
}

func (s *MachineStats) GetURL() string {
	req, _ := s3svc().GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(acksinBucket()),
		Key:    aws.String(s.s3Key()),
	})

	urlStr, err := req.Presign(60 * time.Minute)
	if err != nil {
		log.Println("Failed to sign request", err)
		return ""
	}

	return urlStr
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

	s := &MachineStats{
		Username: username,
		Name:     machineName,
		Stats:    b,
	}

	if err = s.Create(); err != nil {
		log.Println(err)
		respondJSON(w, errorResponse{"Failed to save stats.", 500})
		return
	}

	go s.RunModels()

	respondJSON(w, struct{}{})
}
