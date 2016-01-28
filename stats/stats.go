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
	"fmt"
	"os"

	"github.com/anatma/autotune/stats/fd"
	"github.com/anatma/autotune/stats/memory"
)

type Stats struct {
	CmdName string

	Duration int
	Every    int
}

func (n *Stats) ParseArgs(args []string) {
	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
	flags.IntVar(&n.Every, "every", -1, "Run stats [every] seconds and give average.")
	flags.IntVar(&n.Duration, "duration", -1, "Run command for [duration] seconds.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}
}

func (n *Stats) Run() error {
	type statsResponse struct {
		Memory struct {
			Swapping          bool
			SwappingProcesses []memory.SwappingProcess
		}

		FD struct {
			FileDescriptors []fd.FileDescriptor
		}
	}

	var s statsResponse

	mem := memory.New(nil)
	s.Memory.Swapping = mem.Swapping()
	s.Memory.SwappingProcesses = mem.SwappingProcesses()

	fdesc := fd.New(nil)
	s.FD.FileDescriptors = fdesc.FDs()

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
