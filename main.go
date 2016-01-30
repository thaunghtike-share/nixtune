/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

// +build open

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/anatma/autotune/instance"
	"github.com/anatma/autotune/signatures"
	"github.com/anatma/autotune/stats"
)

var (
	CmdName = "autotune"
	version = "v0.0.0"
)

func SubCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", CmdName, strings.Join(cmds, " "))
}

func Usage() {
	usage := `Usage: %s [command]

Available commands:
    signature [profile]     Update settings based on signature of man application.
    stats                   Gives a quick diagnostics about the state of the machine.
    instance [api_key]      PRO. Recommended instance size for this machine.

Autotune %s by Anatma.
Copyright (c) 2015-2016. Abhi Yerra.
https://anatma.co/autotune
`

	fmt.Printf(usage, CmdName, version)
}

func main() {
	var (
		err error
	)

	if len(os.Args) < 2 {
		Usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "signature":
		sig := signatures.New(SubCmd("signature"))
		sig.ParseArgs(os.Args[2:])
		err = sig.Run()
	case "stats":
		stats := stats.New(SubCmd("stats"))
		stats.ParseArgs(os.Args[2:])
		err = stats.Run()
	case "instance":
		ins := instance.New(SubCmd("instance"))
		ins.ParseArgs(os.Args[2:])
		err = ins.Run()
	default:
		Usage()
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
