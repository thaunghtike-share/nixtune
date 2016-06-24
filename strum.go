/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/acksin/strum/stats"
	ui "github.com/gizak/termui"
)

var (
	version = "0.0"
)

const (
	bridgeService = "fugue.strum"
	bridgeAPIURL  = "http://bridge-api.acksin.com/lambda"
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
	FugueOutput            = "fugue"
)

type config struct {
	apiKey    string
	sessionID string

	output string
	stats  *stats.Stats
}

func main() {
	conf := config{}

	flag.StringVar(&conf.output, "output", "ui", "Formatted outputs available. Available: json, flat, fugue")
	flag.StringVar(&conf.apiKey, "api-key", "", "API Key for Fugue. https://www.acksin.com/fugue")

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
	case FugueOutput:
		postToFugue(&conf)
	default:
		startUI()
	}
}

func postToFugue(conf *config) {
	var err error

	if conf.apiKey == "" {
		fmt.Fprintln(os.Stderr, `Provide the -api-key flag or set the ACKSIN_API_KEY.\nThe API Key can be gathered at 
https://www.acksin.com/fugue/console/#/credentials`)
		os.Exit(-1)
	}

	reqForm := struct {
		Machine string
		Stats   *stats.Stats
	}{
		Machine: "",
		Stats:   conf.stats,
	}

	jsonStr, _ := json.Marshal(reqForm)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://bridge-api.acksin.com/v1/strum/stats", bytes.NewBuffer(jsonStr))
	req.SetBasicAuth(conf.apiKey, "")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("An error occured", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var respForm struct {
		ID string
	}

	err = json.Unmarshal(body, &respForm)
	if err != nil {
		log.Println("An error occured", err)
		return
	}

	fmt.Printf("https://www.acksin.com/fugue/console/#/strum/%s\n", respForm.ID)
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
