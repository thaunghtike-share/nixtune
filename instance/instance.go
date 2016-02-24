/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package instance

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// CloudType represents ia cloud provider.
type CloudType int

// Supported Cloud Providers.
const (
	Aws CloudType = iota
	Azure
	Google
	DigitalOcean
)

type Instance struct {
	// CmdName is the subcommand used to access this feature.
	CmdName string
	// ApiKey is the Fugue key to send metrics.
	ApiKey string
	// MachineName represents how to find the machine on Fugue.
	MachineName string
	// Type is the current cloud provider.
	Type CloudType
	// Frequency that metrics are sent to Fugue
	Every time.Duration
}

func (n *Instance) ParseArgs(args []string) {
	var err error

	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
	flags.StringVar(&n.ApiKey, "fugue-api-key", "", "API key to authenticate with Fugue.")
	flags.StringVar(&n.MachineName, "machine-name", "", "Machine name as to be found in Fugue.")
	every := flags.String("every", "1m", "Send metrics [every] duration.")

	if err = flags.Parse(args); err != nil {
		os.Exit(-1)
	}

	n.Every, err = time.ParseDuration(*every)
	if err != nil {
		os.Exit(-1)
	}
}

func (n *Instance) sendStats(c2 chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		log.Println("sending autotune metrics")
		select {
		case <-time.After(time.Minute * 1):
			continue
		case <-c:
			c2 <- struct{}{}
		}
	}

}

func (n *Instance) Run() error {
	aws := NewAws()
	// TODO: Support other than AWS.
	if aws == nil {
		return fmt.Errorf("not an aws instance.")
	}

	c := make(chan struct{})
	go n.sendStats(c)

	<-c

	return nil
}

func New(cmdName string) *Instance {
	return &Instance{
		CmdName: cmdName,
		Type:    Aws,
	}
}
