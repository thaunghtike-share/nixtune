/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package instance

import (
	"strings"

	"github.com/anatma/autotune/stats"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AwsInstance returns the relevant AWS information about the current
// instance.
type AwsInstance struct {
	// Region is the AWS region that this instance is in.
	Region string
	// InstanceType is the AWS instance that this instance is. Ex. m4.large
	Type string
	// Stats includes current metrics about the machine.
	Stats stats.Response

	metadata *ec2metadata.EC2Metadata
}

func (a *AwsInstance) Family() string {
	if a.Type == "" {
		return ""
	}

	instType := strings.Split(a.Type, ".")
	return instType[0]
}

// NewAws returns an AwsInstance if the current machine is an AWS
// instance otherwise it returns nil.
func NewAws() *AwsInstance {
	sess := session.New()
	i := &AwsInstance{
		metadata: ec2metadata.New(sess),
	}

	// Verify that this is in fact an AWS machine.
	if i.metadata.Available() {
		return i
	}

	return nil
}
