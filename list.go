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
	var listOutput struct {
		Open []string
		Pro  []string
	}

	for _, i := range AssetNames() {
		if strings.HasPrefix(i, "signatures/pro") {
			listOutput.Pro = append(listOutput.Pro, strings.TrimSuffix(strings.TrimPrefix(i, "signatures/pro/"), ".yml"))
		} else {
			listOutput.Open = append(listOutput.Open, strings.TrimSuffix(strings.TrimPrefix(i, "signatures/"), ".yml"))
		}

	}

	sort.Strings(listOutput.Pro)
	sort.Strings(listOutput.Open)

	e, err := json.MarshalIndent(&listOutput, "", "  ")
	if err != nil {
		return -1
	}
	os.Stdout.Write(e)

	return 0
}

// NewList returns a new List object
func NewList(cmdName string) *List {
	s := &List{
		CmdName: cmdName,
	}

	return s
}
