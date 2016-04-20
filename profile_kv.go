/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

type ProfileKV struct {
	Value       string
	Description string   `json:",omitempty"`
	If          []IfCond `json:",omitempty"`
}

func (p *ProfileKV) HasCondition() bool {
	for _, i := range p.If {
		if i.Match() {
			return true
		}
	}
	return false
}
