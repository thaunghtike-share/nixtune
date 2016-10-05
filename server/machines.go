/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	// features is the list of available features that we have about the machine.
	features = [...]string{
		"procfs",
		"sysfs",
		"quick",
	}
)

type machine struct {
	ID        string
	Name      string
	CreatedAt string
}

func (s *machine) features(featuresType string) []byte {
	for _, b := range features {
		if b == featuresType {
			return s.s3Content(featuresType)
		}
	}

	return []byte("")
}

func (s *machine) s3Content(key string) []byte {
	req, err := s3svc().GetObject(&s3.GetObjectInput{
		Bucket: aws.String(acksinBucket()),
		Key:    aws.String(filepath.Join(s.ID, key+".json")),
	})

	if err != nil {
		log.Println(err)
		return []byte("")
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return []byte("")
	}

	return b
}

func GetMachinesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		machines []machine
		username = authUsername(r)
	)

	rows, err := userDB().Query(`SELECT id, name, created_at
                 FROM acksin_machines s1 WHERE s1.username = $1 
                  AND s1.created_at = (SELECT MAX(s2.created_at) 
                                         FROM acksin_machines s2 
                                        WHERE s2.name = s1.name)
             ORDER BY created_at desc;`, username)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var n machine

		if err := rows.Scan(&n.ID, &n.Name, &n.CreatedAt); err == nil {
			machines = append(machines, n)
		} else {
			log.Println(err)
		}
	}

	respondJSON(w, machines)
}

func GetMachineFeaturesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		vars        = mux.Vars(r)
		machineName = vars["machineName"]
		features    = vars["features"]
		username    = authUsername(r)
		m           machine
	)

	err := userDB().QueryRow(`SELECT id, name, created_at 
                   FROM acksin_machines s1
                  WHERE s1.username = $1 
                    AND s1.name = $2
                    AND s1.created_at = (SELECT MAX(s2.created_at)
                                          FROM acksin_machines s2
                                         WHERE s2.username = s1.username
                                           AND s2.name = s2.name);`, username, machineName).Scan(&m.ID, &m.Name, &m.CreatedAt)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, m.features(features))
}
