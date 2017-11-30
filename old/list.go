/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"encoding/json"
	"os"
	"sort"

	"github.com/acksin/autotune/signatures"
)

// List is the command used to update the system settings based
// on the profile specified by the user.
type List struct {
	Open    []string
	Startup []string
	Pro     []string `json:"-"`
	Premium []string `json:"-"`
}

func (k *List) UpdateProfiles() {
	for _, p := range profiles {
		p2 := p.GetProfile()

		switch p2.Subscription {
		case signatures.OpenSubscription:
			k.Open = append(k.Open, p2.Name)
		case signatures.StartupSubscription:
			k.Startup = append(k.Startup, p2.Name)
		case signatures.ProSubscription:
			k.Pro = append(k.Pro, p2.Name)
		case signatures.PremiumSubscription:
			k.Premium = append(k.Premium, p2.Name)
		}
	}

	sort.Strings(k.Open)
	sort.Strings(k.Startup)
	sort.Strings(k.Pro)
	sort.Strings(k.Premium)
}

func (k *List) Synopsis() string {
	return "List all the signatures available"
}

func (k *List) Help() string {
	return "List all the signatures available"
}

func (k *List) Run(args []string) int {
	k.UpdateProfiles()

	gaInvokeEvent("list", "")

	e, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return -1
	}
	os.Stdout.Write(e)

	return 0
}
