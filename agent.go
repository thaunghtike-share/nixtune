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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/acksin/strum/stats"
)

// Agent runs a STRUM Cloud agent.
type agent struct {
	Config struct {
		APIKey   string
		URL      string
		Duration time.Duration
	}
}

func (a *agent) Synopsis() string {
	return "Run a STRUM Cloud agent."
}

func (a *agent) Help() string {
	return "Run a STRUM Cloud agent."
}

func (a *agent) Post() error {
	s := stats.New([]int{})
	jsonStr, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
		return err
	}

	u := statsURL
	if os.Getenv("ACKSIN_DEBUG") != "" {
		u = statsDebugURL
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Acksin-API-Key", a.Config.APIKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var respForm struct {
		ID string
	}

	if err = json.Unmarshal(body, &respForm); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *agent) Run(args []string) int {
	if a.Config.APIKey = os.Getenv("ACKSIN_API_KEY"); a.Config.APIKey == "" {
		fmt.Fprintln(os.Stderr, "Set the ACKSIN_API_KEY.")
		fmt.Fprintln(os.Stderr, "The API Key can be gathered at https://www.acksin.com/console/credentials")
		return -1
	}

	for {
		if err := a.Post(); err != nil {
			return -1
		}

		select {
		case <-time.After(1 * time.Hour):
			log.Println("Ping")
		}
	}
}
