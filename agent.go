/*
 * Anatma Knight - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"flag"
	"os"
	"time"

	sig "github.com/anatma/knight/signatures"
)

type Agent struct {
	Signature string
}

func NewAgent() *Agent {
	return &Agent{}
}

func (k *Agent) ParseArgs(args []string) {
	flags := flag.NewFlagSet(CmdName, flag.ContinueOnError)
	flags.StringVar(&k.Signature, "signature", "", "The signature to use.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}

	// role
	// profile-hints
	// service-name
	// fake run
	// level 0..3
	// interval time between runs. Default is 1 hour.
}

// TODO: Right now it just waits a minute but the goal is to run based
// on various profiles.
func (k *Agent) Profile() {
	// start profiling

	// Make a profile
	time.Sleep(time.Minute)
}

func (k *Agent) Run() {
	var (
		sc *SystemConfig
	)

	for {
		configs := sig.Configs(k.Signature)

		sc.Update(configs)

		k.Profile()
	}
}
