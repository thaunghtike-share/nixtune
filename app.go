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

type App struct{}

func (k *App) Synopsis() string {
	return "Show app changes needed for the signature"
}

func (k *App) Help() string {
	return "Show app changes needed for the signature"
}

func (k *App) Run(args []string) int {
	gaInvokeEvent("app", args[0])

	profile, err := profiles.Get(args[0], false)
	if err != nil {
		log.Println(err)
		return -1
	}

	profile.PrintApp()

	return 0
}
