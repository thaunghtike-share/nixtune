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
	"log"
	"os"
	"os/signal"
	"time"

	"fmt"
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
	return "Autotune Pro feature to recommend Instance sizes."
}

func (n *Instance) Help() string {
	return ""
}

func (n *Instance) Run(args []string) int {
	var err error

	flags := flag.NewFlagSet(n.CmdName, flag.ContinueOnError)
	flags.StringVar(&n.APIKey, "fugue-api-key", "", "API key to authenticate with Fugue.")
	flags.StringVar(&n.MachineName, "machine-name", "", "Machine name as to be found in Fugue.")
	every := flags.String("every", "1m", "Send metrics [every] duration.")

	if err = flags.Parse(args); err != nil {
		return -1
	}

	n.Every, err = time.ParseDuration(*every)
	if err != nil {
		return -1
	}

	aws := NewAws(n.APIKey)
	// TODO: Support other than AWS.
	if aws == nil {
		return -1 // fmt.Errorf("not an aws instance.")
	}

	c := make(chan struct{})
	go n.sendStats(c)

	<-c

	return 0
}

// invokeLambda calls a AWS Lambda function which stores the Metrics.
func (n *Instance) invokeLambda() {
	const (
		functionName = "arn:aws:lambda:us-west-2:451305228097:function:autotune_fugueKey_instance_metrics_POST"
	)

	awsInstance := NewAws(n.APIKey)
	if awsInstance == nil {
		return
	}

	config := aws.NewConfig().WithCredentials(credentials.NewStaticCredentials(n.aws.APIKey, n.aws.SecretKey, "")).WithRegion(n.aws.Region)
	svc := lambda.New(session.New(config))

	fmt.Println(awsInstance.Json())

	params := &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      awsInstance.Json(),
	}
	resp, err := svc.Invoke(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		log.Println(err)
		return
	}

	// Pretty-print the response data.
	log.Println(resp)
	log.Println(string(resp.Payload))
}

func (n *Instance) sendStats(c2 chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		log.Println("sending autotune metrics")
		select {
		case <-time.After(n.Every):
			n.invokeLambda()
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
