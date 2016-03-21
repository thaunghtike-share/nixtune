/* Acksin STRUM - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"flag"
	"fmt"

	"github.com/acksin/strum/stats"
)

// OutputType is the formatted output of the command.
type OutputType string

// Currently available output types.
const (
	JsonOutput  OutputType = "json"
	FlatOutput             = "flat"
	HumanOutput            = "human"
)

type config struct {
	output string
	stats  *stats.Stats
}

func main() {
	conf := config{}

	flag.StringVar(&conf.output, "output", "json", "Formatted outputs available. Available: json, flat, human")
	flag.Parse()

	conf.stats = stats.New()

	switch OutputType(conf.output) {
	case JsonOutput:
		fmt.Printf("%s", conf.stats.Json())
	}
}
