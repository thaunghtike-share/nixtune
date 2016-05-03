/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"log"
)

type SysFS struct {
}

func (k *SysFS) Synopsis() string {
	return "Show the SysFS variables for signature"
}

func (k *SysFS) Help() string {
	return "Show the SysFS variables for signature"
}

func (k *SysFS) Run(args []string) int {
	profile, err := profiles.Get(args[0], false)
	if err != nil {
		log.Println(err)
		return -1
	}
	profile.PrintSysFS()

	return 0
}
