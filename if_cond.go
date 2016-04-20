/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

type IfCond string

const (
	// SSDIfCond if machine has SSD disks.
	SSDIfCond IfCond = "ssd"
)

func (i *IfCond) Match() bool {
	// TODO: Check to see if the condition matches the Hardware.

	return false
}
