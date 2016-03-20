/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package instance

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
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

type awsCreds struct {
	APIKey    string
	SecretKey string
	Region    string
}

type Instance struct {
	// CmdName is the subcommand used to access this feature.
	CmdName string
	// APIKey is the Fugue key to send metrics.
	APIKey string
	// MachineName represents how to find the machine on Fugue.
	MachineName string
	// Type is the current cloud provider.
	Type CloudType
	// Frequency that metrics are sent to Fugue
	Every time.Duration
	// aws contains the credentials for the Lambda function.
	aws awsCreds
}

func (n *Instance) Synopsis() string {
	return "Pro feature to recommend Instance sizes."
}

func (n *Instance) Help() string {
	return ""
}

func (n *Instance) Run(args []string) int {
	var err error

	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
	flags.StringVar(&n.APIKey, "fugue-api-key", "", "API key to authenticate with Fugue.")
	flags.StringVar(&n.MachineName, "machine-name", "", "Machine name as to be found in Fugue.")

	if err = flags.Parse(args); err != nil {
		return -1
	}

	n.Every = time.Minute
	log.Println("sending metrics every minute.")

	if n.MachineName == "" {
		n.MachineName, err = os.Hostname()
		if err != nil {
			log.Println("can't get hostname")
			return -1
		}

		log.Println("no machine-name passed. using hostname", n.MachineName)
	}

	if !n.validAPIKey() {
		log.Println("invalid API key.")
		return -1
	}

	// TODO: Support other than AWS.
	if aws := NewAws(n.APIKey); aws == nil {
		log.Println("only AWS supported currently.")
		return -1
	}

	c := make(chan struct{})
	go n.sendStats(c)
	<-c

	return 0
}

// invokeLambda calls a AWS Lambda functions.
func (n *Instance) invokeLambda(functionName string, payload []byte) ([]byte, error) {
	config := aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(n.aws.APIKey, n.aws.SecretKey, "")).WithRegion(n.aws.Region)
	svc := lambda.New(session.New(config))

	params := &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	}
	resp, err := svc.Invoke(params)

	if err != nil {
		return []byte(""), err
	}

	// Pretty-print the response data.
	log.Println("lambda response", string(resp.Payload))

	return resp.Payload, nil
}

func (n *Instance) validAPIKey() bool {
	const (
		functionName = "arn:aws:lambda:us-west-2:451305228097:function:auth_validate_apikey_POST"
	)

	var apiKey = struct {
		APIKey string
	}{APIKey: n.APIKey}

	js, err := json.Marshal(apiKey)
	if err != nil {
		log.Println("failed to marshall API data", err)
		return false
	}

	payload, err := n.invokeLambda(functionName, js)
	if err != nil {
		log.Println("failed to call API to check API validity.", err)
		return false
	}

	var validity struct {
		Retry int
		Valid bool
	}
	err = json.Unmarshal(payload, &validity)
	if err != nil {
		log.Println("failed to parse api payload", err)
		return false
	}

	return validity.Valid
}

func (n *Instance) sendStats(c2 chan struct{}) {
	const (
		functionName = "arn:aws:lambda:us-west-2:451305228097:function:autotune_fugueKey_instance_metrics_POST"
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		log.Println("sending autotune metrics")
		select {
		case <-time.After(n.Every):
			awsInstance := NewAws(n.APIKey)
			if awsInstance == nil {
				return
			}

			_, err := n.invokeLambda(functionName, awsInstance.Json())
			if err != nil {
				log.Println(err)
			}
		case <-c:
			c2 <- struct{}{}
		}
	}

}

func New(cmdName, apiKey, secretKey, region string) *Instance {
	return &Instance{
		CmdName: cmdName,
		Type:    Aws,
		aws:     awsCreds{apiKey, secretKey, region},
	}
}
