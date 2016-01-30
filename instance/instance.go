/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package instance

import (
	"flag"
	"fmt"
	"os"
)

// CloudType represents ia cloud provider.
type CloudType int

// Supported Cloud Providers.
const (
	Aws CloudType = iota
	Azure
	Google
	DigitalOcean
)

type Instance struct {
	CmdName string

	Type CloudType
}

func (n *Instance) ParseArgs(args []string) {
	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *Instance) Run() error {
	aws := NewAws()

	// TODO: Support other than AWS.
	if aws == nil {
		return fmt.Errorf("not an aws instance.")
	}

	return nil
}

func New(cmdName string) *Instance {
	return &Instance{
		CmdName: cmdName,
		Type:    Aws,
	}
}
