/*
 * Anatma Autotune - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"fmt"
	"os"
)

const (
	CmdName = "autotune"
)

func usage() {
	usage := `
%s [cmd]

    network stats Get network utilization over a period of time.
    network pid   Figure out the profile of the machine based on
                  network processes that are running on the machine.

    memory stats  FUTURE
    memory pid    FUTURE

    io stats      FUTURE
    io pid        FUTURE

    profile       FUTURE: Guess signature of the machine based on memory,
                  network and IO usage.

    server        Update settings based on profile of server.
`

	fmt.Printf(usage, CmdName)
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
	case "network":
		switch os.Args[2] {
		case "stats":
			network := NewNetworkMean()
			network.ParseArgs(os.Args[3:])
			err = network.Run()
		case "pid":
			network := NewNetworkPid()
			network.ParseArgs(os.Args[3:])
			err = network.Run()
		}
	case "server":
		agent := NewAgent()
		agent.ParseArgs(os.Args[2:])
		err = agent.Run()
	}

	if err != nil {
		os.Exit(-1)
	}
}
