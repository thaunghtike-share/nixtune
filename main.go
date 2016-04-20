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
)

func subCmd(cmds ...string) string {
	return fmt.Sprintf("%s %s", cmdName, strings.Join(cmds, " "))
}

func copyright() string {
	return fmt.Sprintf(`Acksin Autotune %s.
Copyright (c) 2016. Acksin.
https://acksin.com/autotune
`, version)
}

func main() {
	// listMode := flag.Bool("list", false, "List all the signatures available.")
	// showFlag := flag.String("show", "all", "What settings to show: procfs, sysfs, env, software, all.")
	// flag.String("output", "json", "Type of output.")
	// showDepsFlag := flag.Bool("deps", true, "Merge deps into this signature.")

	// flag.Usage = func() {
	// 	fmt.Fprintln(os.Stderr, "autotune [flags] [profile]")
	// 	fmt.Fprintln(os.Stderr, "")
	// 	flag.PrintDefaults()
	// 	fmt.Fprintf(os.Stderr, "\n%s", copyright())
	// }
	// flag.Parse()

	// subscription = setSubscription(os.Getenv("ACKSIN_FUGUE_API_KEY"))

	// if len(flag.Args()) < 1 {
	// 	flag.Usage()
	// 	os.Exit(-1)
	// }

	// signature := NewSignature(flag.Args()[0], *showFlag, *showDepsFlag)
	// os.Exit(signature.Run())

	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return NewList(), nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
