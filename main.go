/* Acksin Autotune - Kernel Autotuning
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

	"github.com/acksin/autotune/instance"
	"github.com/acksin/autotune/signatures"
	"github.com/acksin/autotune/stats"
)

var (
	cmdName = "autotune"
	version = "v0.0.0"
)

func subCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", cmdName, strings.Join(cmds, " "))
}

func usage() {
	usage := `Usage: %s [command]

Available commands:
    signature [profile]     Update settings based on signature of man application.
    stats                   Gives a quick diagnostics about the state of the machine.
    instance [api_key]      PRO. Recommended instance size for this machine.

Acksin Autotune %s.
Copyright (c) 2015-2016. Abhi Yerra.
https://acksin.com/autotune
`
	fmt.Printf(usage, cmdName, version)
}

func main() {
	var (
		err error
	)

	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "signature":
		sig := signatures.New(subCmd("signature"))
		sig.ParseArgs(os.Args[2:])
		err = sig.Run()
	case "stats":
		stats := stats.New(subCmd("stats"))
		stats.ParseArgs(os.Args[2:])
		err = stats.Run()
	case "instance":
		ins := instance.New(subCmd("instance"))
		ins.ParseArgs(os.Args[2:])
		err = ins.Run()
	default:
		usage()
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
