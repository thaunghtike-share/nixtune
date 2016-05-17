/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/cli"

	"github.com/acksin/autotune/signatures"
	"github.com/acksin/gmeasure"
)

var (
	cmdName = "autotune"
	version = "v0.0.0"
)

var (
	profiles signatures.Profiles
	ga       *gmeasure.GA

	cid = "bff73d42-1c83-11e6-b6ba-3e1d05defe78"
)

func gaCid() string {
	return fmt.Sprintf("%s-autotune%s", cid, version)
}

func gaInvokeEvent(action, label string) {
	ga.Event(gmeasure.Event{
		Category: "AutotuneInvoke",
		Action:   action,
		Label:    label,
	})
}

func subCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", cmdName, strings.Join(cmds, " "))
}

func copyright() string {
	return fmt.Sprintf(`Acksin Autotune %s.
Copyright (c) 2016. Acksin.
https://www.acksin.com/autotune`, version)
}

func main() {
	ga = gmeasure.New("UA-75403807-1", gaCid())
	profiles = signatures.Load()

	c := cli.NewCLI(cmdName, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &List{}, nil
		},
		"sig": func() (cli.Command, error) {
			return &Signature{}, nil
		},
		"procfs": func() (cli.Command, error) {
			return &ProcFS{}, nil
		},
		"sysfs": func() (cli.Command, error) {
			return &SysFS{}, nil
		},
		"files": func() (cli.Command, error) {
			return &Files{}, nil
		},
		"app": func() (cli.Command, error) {
			return &App{}, nil
		},
		"env": func() (cli.Command, error) {
			return &Env{}, nil
		},
	}
	c.HelpFunc = func(commands map[string]cli.CommandFactory) string {
		return fmt.Sprintf("%s\n%s", cli.BasicHelpFunc(cmdName)(commands), copyright())
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
