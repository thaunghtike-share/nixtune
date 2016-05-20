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

	"bytes"
	"encoding/json"
	"github.com/acksin/strum/stats"
	"io/ioutil"
	"net/http"
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

	flag.StringVar(&conf.output, "output", "json", "Formatted outputs available. Available: json, flat, fugue")
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
}
