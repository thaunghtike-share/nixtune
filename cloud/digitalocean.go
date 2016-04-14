/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package cloud

import (
	"github.com/digitalocean/go-metadata"
)

// AWSStats returns the relevant AWS information about the current
// instance via the machine's EC2 Metadata IP interface.
type DigitalOceanStats struct {
	*metadata.Metadata
}

// NewAWS returns an AWSStats if the current machine is an AWS
// instance otherwise it returns nil.
func NewDigitalOcean() *DigitalOceanStats {
	// Create a client
	client := metadata.NewClient()

	// Request all the metadata about the current droplet
	all, err := client.Metadata()
	if err != nil {
		return nil
	}

	return &DigitalOceanStats{all}
}
