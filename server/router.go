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

	r.Handle("/v1/autotune/stats", authMiddleware(PostStatsHandler)).Methods("POST") // Deprecated.
	r.Handle("/v1/stats", authMiddleware(PostStatsHandler)).Methods("POST")

	// Console API
	r.Handle("/v1/machines", authMiddleware(GetMachinesHandler)).Methods("GET")
	r.Handle("/v1/machines/{id}/stats", authMiddleware(GetMachineDiagnosticsHandler)).Methods("GET")
	r.Handle("/v1/machines/{id}/tuning/{tuningType}", authMiddleware(GetMachineTuningHandler)).Methods("GET")
	//r.Handle("/v1/machines/{id}/security", authMiddleware(GetAutotuneAutotuneHandler)).Methods("GET")

	r.Handle("/v1/cloud/aws", authMiddleware(GetAWSCloudHandler)).Methods("GET")
	// r.Handle("/v1/security", authMiddleware()).Methods("GET")
	// r.Handle("/v1/cloud", authMiddleware()).Methods("GET")
	// r.Handle("/v1/cloud/stats/{statId}", authMiddleware()).Methods("GET")
	// r.Handle("/v1/cloud/{instanceId}", authMiddleware()).Methods("GET")

	return r
}
