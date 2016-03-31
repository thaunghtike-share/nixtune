/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"log"

	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Profile struct {
	// Name of the profile
	Name string
	// Description of the service that is being updated.
	Description string
	// Documentation is the documentation for this profile
	Documentation string
	// ProcFS contains the kernel key values that will be changed.
	ProcFS map[string]struct {
		Value       string
		Description string
	} `yaml:"procfs"`

	SysFS map[string]struct {
		Value       string
		Description string
	} `yaml:"sysfs"`

	// Env is the environment variables that should be changed for
	// maximum performance.
	Env map[string]struct {
		Value       string
		Description string
	}

	Vars map[string]string

	// Deps of other profiles.
	Deps []string
}

func ParseProfile(s []byte) (p Profile) {
	p = Profile{}

	log.Println(string(s))

	err := yaml.Unmarshal(s, &p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return
}
