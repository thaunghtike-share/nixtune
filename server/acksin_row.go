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
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AcksinRow struct {
	ID          string
	Username    string
	MachineName string
	Stats       []byte
}

func (s *AcksinRow) s3Key() string {
	return fmt.Sprintf("stats/%s/%s", s.Username, s.ID)
}

func (s *AcksinRow) Create() error {
	userDB().QueryRow("INSERT INTO autotune_stats(username, data, machine_name) VALUES ($1, $2, $3) RETURNING id", s.Username, s.Stats, s.MachineName).Scan(&s.ID)

	// Backup Write to S3
	params := &s3.PutObjectInput{
		Bucket: aws.String(acksinRowBucket()),
		Key:    aws.String(s.s3Key()),
		Body:   bytes.NewReader(s.Stats),
	}

	_, err := s3svc().PutObject(params)
	if err != nil {
		return err
	}

	return nil
}

func (s *AcksinRow) GetURL() string {
	req, _ := s3svc().GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(acksinRowBucket()),
		Key:    aws.String(s.s3Key()),
	})

	urlStr, err := req.Presign(60 * time.Minute)
	if err != nil {
		log.Println("Failed to sign request", err)
		return ""
	}

	return urlStr
}

func (s *AcksinRow) RunModels(r struct{ ID string }) {
	b, _ := json.Marshal(r)

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
