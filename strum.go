/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/acksin/strum/fd"
	"github.com/acksin/strum/memory"
)

func printJson(s interface{}) error {
	js, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(js))

	return nil
}

func main() {
	n := &Stats{}
	s := Response{}

	flag.IntVar(&n.Every, "every", -1, "Run stats [every] seconds and give average.")
	flag.IntVar(&n.Duration, "duration", -1, "Run command for [duration] seconds.")

	flag.Parse()

	s.System.Memory = memory.New()

	for _, proc := range n.processes() {
		exe, err := proc.Executable()
		if err != nil || exe == "" {
			status, _ := proc.NewStatus()
			exe = status.Name
		}

		p := Process{
			Exe:    exe,
			PID:    proc.PID,
			Memory: memory.NewProcess(proc),
			FD:     fd.NewProcess(proc),
		}

		s.Processes = append(s.Processes, p)
	}

	printJson(s)
}
