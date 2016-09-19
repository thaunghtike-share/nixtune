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
)

var (
	version string
)

type Server struct {
	Version string
}

func (s *Server) Synopsis() string {
	return "Acksin Server"
}

func (s *Server) Help() string {
	return ""
}

func (s *Server) Run(args []string) int {
	version = s.Version

	setup()

	log.Fatal(http.ListenAndServe(":8080", commonRouter(router())))

	return 0
}
