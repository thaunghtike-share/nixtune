/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/acksin/fugue/commons"
	"github.com/acksin/fugue/config"
)

var (
	c *config.Config
)

func userDB() *sql.DB {
	return c.UserDB()
}

func acksinRowBucket() string {
	return c.Acksin.Bucket
}

func mentalModelFunction() string {
	return c.Acksin.MentalModelFunction
}

func authMiddleware(handler http.HandlerFunc) http.Handler {
	return commons.AuthMiddleware(http.HandlerFunc(handler))
}

func authUsername(r *http.Request) string {
	return commons.UsernameValue(r)
}

func s3svc() *s3.S3 {
	return c.S3()
}

func lambdasvc() *lambda.Lambda {
	return c.Lambda()
}

func commonRouter(r *mux.Router) http.Handler {
	return commons.CommonRouter(router())
}

func setup() {
	c = config.NewConfig(os.Getenv("ACKSIN_ENV"), config.AcksinApp)
	commons.Setup(c)
}
