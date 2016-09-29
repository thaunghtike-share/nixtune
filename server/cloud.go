/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"log"
	"net/http"
	//	"github.com/aws/aws-sdk-go/service/ec2"
)

func CloudUpdater() {

}

func GetAWSCloudHandler(w http.ResponseWriter, r *http.Request) {
	type autotuneNode struct {
		ID           string
		CreatedAt    string
		InstanceID   string
		InstanceType string
	}

	var (
		nodes    []autotuneNode
		username = authUsername(r)
	)

	rows, err := userDB().Query(`SELECT 
                      id, 
                      created_at, 
                      data->'Cloud'->'AWS'->>'InstanceID' as instance_id, 
                      data->'Cloud'->'AWS'->>'InstanceType' as instance_type 
                 FROM autotune_stats s1 WHERE s1.username = $1 
                  AND s1.created_at = (SELECT MAX(s2.created_at) 
                                         FROM autotune_stats s2 
                                        WHERE s2.data->'Cloud'->'AWS'->'InstanceID' = s1.data->'Cloud'->'AWS'->'InstanceID')
             ORDER BY created_at desc;`, username)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var n autotuneNode

		err := rows.Scan(&n.ID, &n.CreatedAt, &n.InstanceID, &n.InstanceType)
		if err != nil {
			log.Println(err)
			continue
		}

		nodes = append(nodes, n)
	}

	respondJSON(w, nodes)
}
