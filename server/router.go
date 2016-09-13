/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/v1/autotune/stats", authMiddleware(PostAutotuneStatsHandler)).Methods("POST") // Deprecated.
	r.Handle("/v1/stats", authMiddleware(PostAutotuneStatsHandler)).Methods("POST")

	// Console API
	r.Handle("/v1/autotune/nodes/aws", authMiddleware(GetAutotuneAWSNodesHandler)).Methods("GET")
	r.Handle("/v1/autotune/nodes", authMiddleware(GetAutotuneNodesHandler)).Methods("GET")
	r.Handle("/v1/autotune/stats/{id}", authMiddleware(GetAutotuneStatsHandler)).Methods("GET")
	r.Handle("/v1/autotune/tuning/{id}", authMiddleware(GetAutotuneAutotuneHandler)).Methods("GET")
	r.Handle("/v1/autotune/autotune/{id}", authMiddleware(GetAutotuneAutotuneHandler)).Methods("GET") // Deprecated

	return r
}
