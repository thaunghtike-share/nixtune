/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/acksin/strum/stats"
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

// Currently available output types.
const (
	JSONOutput   = "json"
	FlatOutput   = "flat"
	AcksinOutput = "cloud"
)

type config struct {
	apiKey    string
	sessionID string

	stats *stats.Stats
}

func main() {
	conf := config{}

	cloudOut := flag.Bool("cloud", false, "Send to STRUM Cloud.")
	jsonOut := flag.Bool("json", true, "Formatted output as JSON.")
	flatOut := flag.Bool("flat", false, "Formatted output as flat key value.")
	flag.StringVar(&conf.apiKey, "api-key", os.Getenv("ACKSIN_API_KEY"), "API Key for Acksin. https://www.acksin.com/console/credentials")

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

	switch {
	case *cloudOut:
		postToAcksin(&conf)
	case *flatOut:
		fmt.Printf("%s", conf.stats.Flat())
	case *jsonOut:
		fmt.Printf("%s", conf.stats.JSON())
	}
}
