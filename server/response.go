/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch v := i.(type) {
	case errorResponse:
		http.Error(w, string(jsonResponse), v.Code)
		return
	case []byte:
		w.Write(i.([]byte))
		return
	}

	w.Write(jsonResponse)
}

type errorResponse struct {
	Message string
	Code    int
}
