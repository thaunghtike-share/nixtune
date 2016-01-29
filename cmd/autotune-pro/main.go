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
	"fmt"
	"os"

	. "github.com/anatma/autotune/cmd"
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

func main() {
	var (
		handled bool
		err     error
	)

	handled, err = SubCommands()
	if !handled {
		handled, err = subCommandsPro()
		if !handled {
			Usage()
			os.Exit(-1)
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
