/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"encoding/json"
	"log"
	"os"
)

// Signature is the command used to update the system settings based
// on the profile specified by the user.
type Signature struct {
}

func (k *Signature) Synopsis() string {
	return "Show JSON output of changes"
}

func (k *Signature) Help() string {
	return "Show JSON output of changes"
}

// Run gets the configuration for the profile and updates the system
// settings with the new values.
func (k *Signature) Run(args []string) int {
	var (
		showDeps bool = false
	)

	if len(args) < 1 {
		log.Println("need to pass the signature")
		return -1
	}

	gaInvokeEvent("sig", args[0])

	profile, err := profiles.Get(args[0], showDeps)
	if err != nil {
		log.Println(err)
		return -1
	}

	err = profile.ParseFlags(args[1:])
	if err != nil {
		log.Println(err)
		return -1
	}

	e, err := json.MarshalIndent(&profile, "", "  ")
	if err != nil {
		return -1
	}
	os.Stdout.Write(e)

	return 0
}
