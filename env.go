/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

type Env struct {
}

func (k *Env) Synopsis() string {
	return "Show the Environment variables for signature"
}

func (k *Env) Help() string {
	return "Show the Environment variables for signature"
}

func (k *Env) Run(args []string) int {
	profile := profiles.Get(args[0], false)
	if profile == nil {
		return -1
	}

	profile.PrintEnv()

	return 0
}
