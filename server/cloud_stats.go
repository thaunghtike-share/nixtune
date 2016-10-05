/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"fmt"
)

type CloudStats struct {
	Username  string
	Provider  string
	Timestamp int64
}

// username/cloud/aws/<timestamp>/ec2/<instance>/stats.json
func (s *CloudStats) ID() string {
	return fmt.Sprintf("%s/cloud/%s/%d", s.Username, s.Provider, s.Timestamp)
}
