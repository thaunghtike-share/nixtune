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

	"github.com/gorilla/mux"
)

func GetMachinesHandler(w http.ResponseWriter, r *http.Request) {
	type autotuneNode struct {
		ID          string
		CreatedAt   string
		MachineName string
	}

	var (
		nodes    []autotuneNode
		username = authUsername(r)
	)

	rows, err := userDB().Query(`SELECT 
                      id, 
                      created_at, 
                      machine_name 
                 FROM autotune_stats s1 WHERE s1.username = $1 
                  AND s1.created_at = (SELECT MAX(s2.created_at) 
                                         FROM autotune_stats s2 
                                        WHERE s2.machine_name = s1.machine_name)
             ORDER BY created_at desc;`, username)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var n autotuneNode

		err := rows.Scan(&n.ID, &n.CreatedAt, &n.MachineName)
		if err != nil {
			log.Println(err)
			continue
		}

		nodes = append(nodes, n)
	}

	respondJSON(w, nodes)
}

// TODO: This logic is completely broken
func GetMachineDiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		username = authUsername(r)
	)

	s := &MachineStats{
		Username: username,
		Name:     id,
	}

	url := s.GetURL()
	if url == "" {
		respondJSON(w, errorResponse{"Failed to find that id", 404})
		return
	}

	respondJSON(w, struct{ URL string }{url})
}

func GetMachineTuningHandler(w http.ResponseWriter, r *http.Request) {
	type autotune struct {
		ID     string
		ProcFS JSONB
		SysFS  JSONB
	}

	var (
		vars     = mux.Vars(r)
		a        autotune
		username = authUsername(r)
	)

	err := userDB().QueryRow(`SELECT id, procfs_features, sysfs_features FROM autotune_stats WHERE username = $1 AND id = $2;`, username, vars["id"]).Scan(&a.ID, &a.ProcFS, &a.SysFS)
	if err != nil {
		log.Println(err)
	}

	respondJSON(w, a)
}
