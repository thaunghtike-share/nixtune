/*
 * Anatma Autotune - Kernel Autotuning
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

	sig "github.com/anatma/autotune/signatures"
)

type Agent struct {
	Signature string
}

func (k *Agent) ParseArgs(args []string) {
	flags := flag.NewFlagSet(CmdName, flag.ContinueOnError)
	flags.StringVar(&k.Signature, "signature", "", "The signature to use.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (k *Agent) Run() error {
	sc := &SystemConfig{}
	sc.Update(sig.Configs(k.Signature))

	return nil
}

func NewAgent() *Agent {
	return &Agent{}
}
