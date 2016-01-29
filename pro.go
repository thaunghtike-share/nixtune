/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

// +build pro

package main

import (
	"os"

	"github.com/anatma/fugue/autotune/instance"
)

func subCommandsPro() (handled bool, err error) {
	switch os.Args[1] {
	case "instance":
		sig := instance.New(SubCmd("instance"))
		sig.ParseArgs(os.Args[2:])
		err = sig.Run()
		handled = true
	}

	return handled, err
}
