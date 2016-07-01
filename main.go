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
	ui "github.com/gizak/termui"
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
	JSONOutput   OutputType = "json"
	FlatOutput              = "flat"
	AcksinOutput            = "acksin"
)

type config struct {
	apiKey    string
	sessionID string

	output string
	stats  *stats.Stats
}

func main() {
	conf := config{}

	flag.StringVar(&conf.output, "output", "ui", "Formatted outputs available. Available: json, flat, acksin")
	flag.StringVar(&conf.apiKey, "api-key", "", "API Key for Acksin. https://www.acksin.com/console")

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
	case AcksinOutput:
		postToAcksin(&conf)
	default:
		startUI()
	}
}

func startUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar(":PRESS q TO QUIT DEMO")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Text Box"
	p.BorderFg = ui.ColorCyan

	g := ui.NewGauge()
	g.Percent = 50
	g.Width = 50
	g.Height = 3
	g.Y = 11
	g.BorderLabel = "Gauge"
	g.BarColor = ui.ColorRed
	g.BorderFg = ui.ColorWhite
	g.BorderLabelFg = ui.ColorCyan

	ui.Render(p, g)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}
