/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func init() {
// 	isTesting = true
// 	config = NewConfig("./dev.json")
// }

// func TestPostAutotuneStatsHandler(t *testing.T) {
// 	jsonStr := []byte(`{"Stats":{"Cloud":{"AWS":{"Name":"12345"}}},"Machine":"barfoo"}`)

// 	req, _ := http.NewRequest("POST", "/v1/autotune/stats", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth("foobar", "")

// 	w := httptest.NewRecorder()

// 	router().ServeHTTP(w, req)

// 	var c struct {
// 		ID string
// 	}

// 	err := json.Unmarshal(w.Body.Bytes(), &c)
// 	if err != nil {
// 		t.Errorf("Failed to unmarshall %s", w.Body.String())
// 	}

// 	if c.ID == "" {
// 		t.Errorf("Failed to create a new API Key")
// 	}

// 	if w.Code != http.StatusOK {
// 		t.Errorf("Home page didn't return %v", http.StatusOK)
// 	}
// }

// func TestListCredentialsHandlerFail(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/credentials", nil)
// 	w := httptest.NewRecorder()

// 	router().ServeHTTP(w, req)

// 	var c struct {
// 		Message string
// 	}
// 	err := json.Unmarshal(w.Body.Bytes(), &c)
// 	if err != nil {
// 		t.Errorf("Failed to unmarshall %s", w.Body.String())
// 	}

// 	if c.Message == "" {
// 		t.Errorf("Failed to unmarshall %s", w.Body.String())
// 	}
// }
