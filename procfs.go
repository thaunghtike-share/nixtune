/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

type ProcFS struct {
}

func (k *ProcFS) Synopsis() string {
	return "Show the ProcFSironment variables for signature"
}

func (k *ProcFS) Help() string {
	return "Show the ProcFSironment variables for signature"
}

func (k *ProcFS) Run(args []string) int {
	profile := profiles.Get(args[0], false)
	profile.PrintProcFS()

	return 0
}
