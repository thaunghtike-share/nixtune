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

	"github.com/mitchellh/cli"

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

func copyright() string {
	return fmt.Sprintf(`Acksin Autotune %s.
Copyright (c) 2015-2016. Abhi Yerra.
https://acksin.com/autotune
`, version)
}

func main() {
	c := cli.NewCLI(cmdName, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"signature": func() (cli.Command, error) {
			return signatures.New(subCmd("signature")), nil
		},
		"stats": func() (cli.Command, error) {
			return stats.New(subCmd("stats")), nil
		},
		"instance": func() (cli.Command, error) {
			return instance.New(subCmd("instance")), nil
		},
	}

	c.HelpFunc = func(commands map[string]cli.CommandFactory) string {
		return fmt.Sprintf("%s\n%s", cli.BasicHelpFunc(cmdName)(commands), copyright())
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)

}
