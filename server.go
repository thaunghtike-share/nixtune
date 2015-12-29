/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package main

import (
	"flag"
	"os"

	sig "github.com/anatma/autotune/signatures"
)

type Server struct {
	Signature string
}

func (k *Server) ParseArgs(args []string) {
	flags := flag.NewFlagSet(subCmd("server"), flag.ContinueOnError)
	flags.StringVar(&k.Signature, "signature", "", "The signature to use.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (k *Server) Run() error {
	sc := &SystemConfig{}
	sc.Update(sig.Configs(k.Signature))

	return nil
}

func NewServer() *Server {
	return &Server{}
}
