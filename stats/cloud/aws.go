/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package cloud

import (
	"reflect"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/acksin/autotune/shared"
)

// AWSStats returns the relevant AWS information about the current
// instance via the machine's EC2 Metadata IP interface.
type AWSStats struct {
	AmiID           string `metadata:"ami-id" `
	AmiLaunchIndex  string `metadata:"ami-launch-index"`
	AmiManifestPath string `metadata:"ami-manifest-path"`
	Hostname        string `metadata:"hostname"`
	InstanceAction  string `metadata:"instance-action"`
	InstanceID      string `metadata:"instance-id"`
	InstanceType    string `metadata:"instance-type"`
	LocalHostname   string `metadata:"local-hostname"`
	LocalIpv4       string `metadata:"local-ipv4"`
	MAC             string `metadata:"mac"`
	Profile         string `metadata:"profile"`
	PublicHostname  string `metadata:"public-hostname"`
	PublicIpv4      string `metadata:"public-ipv4"`
	ReservationID   string `metadata:"reservation-id"`
	SecurityGroups  string `metadata:"security-groups"`

	Placement struct {
		AvailabilityZone string `metadata:"availability-zone"`
	} `metadata:"placement"`

	Spot struct {
		Termination string `metadata:"termination" json:",omitempty"`
	} `metadata:"spot" json:",omitempty"`

	// TODO: These need to be converted into iteration tasks.
	// block-device-mapping/
	// metrics/
	// network/
	// public-keys/
	// services/

	// Full
}

func (a *AWSStats) parseMetadata(m *ec2metadata.EC2Metadata) {
	st := reflect.TypeOf(*a)

	for i := 0; i < st.NumField(); i++ {
		metadataTag := st.Field(i).Tag.Get("metadata")
		if metadataTag != "" {
			data, err := m.GetMetadata(metadataTag)
			if err != nil {
				continue
			}

			reflect.ValueOf(a).Elem().Field(i).SetString(data)
		}
	}
}

func (i *AWSStats) getFull() {
	svc := ec2.New(session.New(), &aws.Config{
		Region: aws.String("us-west-2"),
	})

	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
	for idx, res := range resp.Reservations {
		fmt.Println("  > Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println("    - Instance ID: ", *inst.InstanceId)
		}
	}
}

// NewAWS returns an AWSStats if the current machine is an AWS
// instance otherwise it returns nil.
func NewAWS(a *shared.Config) (i *AWSStats) {
	metadata := ec2metadata.New(session.New())

	// Verify that this is in fact an AWS machine.
	if !metadata.Available() {
		return nil
	}

	i = &AWSStats{}
	i.parseMetadata(metadata)

	if data, err := metadata.GetMetadata("placement/availability-zone"); err == nil {
		i.Placement.AvailabilityZone = data
	}

	if data, err := metadata.GetMetadata("spot/termination"); err == nil {
		i.Spot.Termination = data
	}

	// if a.Cloud != nil && a.Cloud.AWS != nil {
	// 	i.getFull()
	// }

	return
}
