/* Acksin STRUM - Linux Diagnostics
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
	"strconv"

	"github.com/acksin/strum/stats"
	flag "github.com/ogier/pflag"
)

var (
	version = "0.0"
)

func copyright() string {
	return fmt.Sprintf(`Acksin STRUM %s.
Copyright (c) 2016. Acksin.
https://acksin.com/strum
`, version)
}

// OutputType is the formatted output of the command.
type OutputType string

// Currently available output types.
const (
	JSONOutput  OutputType = "json"
	FlatOutput             = "flat"
	HumanOutput            = "human"
)

type config struct {
	output string
	stats  *stats.Stats
}

func main() {
	conf := config{}

	flag.StringVarP(&conf.output, "output", "o", "json", "Formatted outputs available. Available: json, flat, human")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "strum [flags] [pid]")
		fmt.Fprintln(os.Stderr, "")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n%s", copyright())
	}
	flag.Parse()

	var pids []int
	for _, pid := range flag.Args() {
		i, err := strconv.Atoi(pid)
		if err != nil {
			log.Fatalf("failed to parse %s", pid)
		}

		pids = append(pids, i)
	}

	conf.stats = stats.New(pids)

	switch OutputType(conf.output) {
	case JSONOutput:
		fmt.Printf("%s", conf.stats.JSON())
	case FlatOutput:
		fmt.Printf("%s", conf.stats.Flat())
	case HumanOutput:
		fmt.Printf("%s", conf.stats.Human())
	}
}
