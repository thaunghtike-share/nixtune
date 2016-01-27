/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package stats

import (
	"encoding/json"
	"flag"
	"os"

	"fmt"
	"github.com/anatma/autotune/stats/memory"
)

type Stats struct {
	CmdName string

	//network  *Network
	Duration int
}

func (n *Stats) ParseArgs(args []string) {
	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
	flags.IntVar(&n.Duration, "duration", 60, "Duration to monitor in seconds. Defaults to 60 seconds.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *Stats) Run() error {
	type statsResponse struct {
		Memory struct {
			Swapping bool
			// SwappingProcesses map[string]bool
		}
	}

	var s statsResponse

	mem := memory.New(nil)
	s.Memory.Swapping = mem.Swapping()
	// s.Memory.SwappingProcesses = mem.SwappingProcesses()

	return printJson(s)
}

func printJson(s interface{}) error {
	js, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(js))

	return nil
}

func New(cmdName string) *Stats {
	return &Stats{
		CmdName: cmdName,
	}
}
