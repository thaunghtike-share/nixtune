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
	"os"
)

// Signature is the command used to update the system settings based
// on the profile specified by the user.
type Signature struct {
}

func (k *Signature) Synopsis() string {
	return "Show the Signatureironment variables for signature"
}

func (k *Signature) Help() string {
	return "Show the Signatureironment variables for signature"
}

// Run gets the configuration for the profile and updates the system
// settings with the new values.
func (k *Signature) Run(args []string) int {
	var (
		showDeps bool = false
	)

	profile := profiles.Get(args[0], showDeps)

	e, err := json.MarshalIndent(&profile, "", "  ")
	if err != nil {
		return -1
	}
	os.Stdout.Write(e)

	return 0
}
