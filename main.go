/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	//	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/cli"
)

var (
	cmdName = "autotune"
	version = "v0.0.0"
)

var (
	subscription Subscription

	profiles Profiles
)

func subCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", cmdName, strings.Join(cmds, " "))
}

func copyright() string {
	return fmt.Sprintf(`Acksin Autotune %s.
Copyright (c) 2016. Acksin.
https://www.acksin.com/autotune`, version)
}

func loadProfiles() {
	for _, i := range AssetNames() {
		ymlData, err := Asset(i)
		if err != nil {
			log.Fatal(err)
		}
		p := ParseProfile(ymlData)
		profiles = append(profiles, p)
	}
}

func main() {
	loadProfiles()

	subscription = setSubscription(os.Getenv("ACKSIN_FUGUE_API_KEY"))

	c := cli.NewCLI("autotune", version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return NewList(), nil
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
