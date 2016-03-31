/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"strings"
)

// List is the command used to update the system settings based
// on the profile specified by the user.
type List struct {
	CmdName string
}

func (k *List) Synopsis() string {
	return "List the signatures."
}

func (k *List) Help() string {
	return ""
}

func (k *List) Run(args []string) int {
	var (
		freeList []string
		proList  []string
	)

	for _, i := range AssetNames() {
		if strings.HasPrefix(i, "signatures/pro") {
			proList = append(proList, strings.TrimSuffix(strings.TrimPrefix(i, "signatures/pro/"), ".json"))
		} else {
			freeList = append(freeList, strings.TrimSuffix(strings.TrimPrefix(i, "signatures/"), ".json"))
		}

	}

	fmt.Println("Free Signatures")
	for _, i := range freeList {
		fmt.Println(i)
	}

	fmt.Println("\nPro Signatures")
	for _, i := range proList {
		fmt.Println(i)
	}

	return 0
}

// NewList returns a new List object
func NewList(cmdName string) *List {
	// New returns a map of profile and system configurations that are
	// currently supported.
	s := &List{
		CmdName: cmdName,
	}

	return s
}
