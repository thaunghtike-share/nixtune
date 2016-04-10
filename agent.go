/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

// import (
// 	"encoding/json"
// 	"flag"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"github.com/acksin/fugue/client"
// 	"github.com/acksin/strum/stats"
// )

// type Agent struct {
// 	// CmdName is the subcommand used to access this feature.
// 	CmdName string `json:"-"`
// 	// APIKey is the Fugue key to send metrics.
// 	APIKey string `json:"-"`
// 	// MachineName represents how to find the machine on Fugue.
// 	MachineName string
// 	// Frequency that metrics are sent to Fugue
// 	Every time.Duration `json:"-"`
// 	// APIToken is the token used to send stats.
// 	APIToken string
// 	// Stats is the stats that are sent to the server.
// 	Stats *stats.Stats
// }

// func (n *Agent) Synopsis() string {
// 	return "Pro feature to recommend Instance sizes."
// }

// func (n *Agent) Help() string {
// 	return ""
// }

// func (n *Agent) verifyMachineName() (err error) {
// 	if n.MachineName == "" {
// 		n.MachineName, err = os.Hostname()
// 		if err != nil {
// 			log.Println("can't get hostname")
// 			return err
// 		}

// 		log.Println("no machine-name passed. using hostname", n.MachineName)
// 	}

// 	return nil
// }

// func (n *Agent) Run(args []string) int {
// 	var err error

// 	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
// 	flags.StringVar(&n.APIKey, "fugue-api-key", "", "API key to authenticate with Fugue.")
// 	flags.StringVar(&n.MachineName, "machine-name", "", "Machine name as to be found in Fugue.")

// 	if err = flags.Parse(args); err != nil {
// 		return -1
// 	}

// 	if n.APIKey == "" {
// 		log.Println("need a Acksin Fugue API Key. Get it at http://www.acksin.com/fugue")
// 		return -1
// 	}

// 	n.Every = time.Minute
// 	log.Println("sending metrics every minute.")

// 	err = n.verifyMachineName()
// 	if err != nil {
// 		return -1
// 	}

// 	n.APIToken, err = n.getAPIToken()
// 	if err != nil {
// 		log.Println("invalid API token.")
// 		return -1
// 	}

// 	if n.APIToken == "" {
// 		log.Println("invalid API token.")
// 		return -1
// 	}

// 	c := make(chan struct{})
// 	go n.sendStats(c)
// 	<-c

// 	return 0
// }

// func (n *Agent) sendStats(c2 chan struct{}) {
// 	const (
// 		functionName = "arn:aws:lambda:us-west-2:451305228097:function:autotune_instance_metrics_POST-dev"
// 	)

// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)

// 	for {
// 		log.Println("sending autotune metrics")

// 		n.Stats = stats.New([]int{})
// 		payload := n.JSON()
// 		if len(payload) > 0 {
// 			// _, err := invokeLambda(functionName, payload)
// 			// if err != nil {
// 			// 	log.Println(err)
// 			// }
// 		}

// 		select {
// 		case <-time.After(n.Every):
// 			continue
// 		case <-c:
// 			c2 <- struct{}{}
// 		}
// 	}

// }

// // JSON returns JSON string of Stats
// func (n *Agent) JSON() []byte {
// 	js, err := json.MarshalIndent(n, "", "  ")
// 	if err != nil {
// 		return []byte("")
// 	}

// 	return js
// }

// func NewAgent(cmdName string) *Agent {
// 	return &Agent{
// 		CmdName: cmdName,
// 	}
// }
