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

	"gopkg.in/yaml.v2"
)

type ProfileKV struct {
	Value       string
	Description string
}

type Profile struct {
	// Name of the profile
	Name string
	// Description of the service that is being updated.
	Description string
	// Documentation is the documentation for this profile
	Documentation string
	// ProcFS contains the kernel key values that will be changed.
	ProcFS map[string]ProfileKV `yaml:"procfs"`

	SysFS map[string]ProfileKV `yaml:"sysfs"`

	// Env is the environment variables that should be changed for
	// maximum performance.
	Env map[string]ProfileKV

	Vars map[string]string

	// Deps of other profiles.
	Deps []string
}

func ParseProfile(s []byte) (p *Profile) {
	p = &Profile{}

	err := yaml.Unmarshal(s, p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return
}

type Profiles []*Profile

func (p Profiles) Get(s string) *Profile {
	var profile *Profile

	for _, k := range p {
		if k.Name == s {
			profile = k
			break
		}
	}

	if profile == nil {
		return nil
	}

	// Get the dependency profiles as we will construct a new
	// profile with everything as one.
	var depProfiles Profiles
	for _, dep := range profile.Deps {
		depProfiles = append(depProfiles, p.Get(dep))
	}

	if profile.ProcFS == nil {
		profile.ProcFS = make(map[string]ProfileKV)
	}

	if profile.SysFS == nil {
		profile.SysFS = make(map[string]ProfileKV)
	}

	if profile.Env == nil {
		profile.Env = make(map[string]ProfileKV)
	}

	for _, dep := range depProfiles {
		for k, v := range dep.ProcFS {
			_, ok := profile.ProcFS[k]
			if !ok {
				profile.ProcFS[k] = v
			}
		}

		for k, v := range dep.SysFS {
			_, ok := profile.SysFS[k]
			if !ok {
				profile.SysFS[k] = v
			}
		}

		for k, v := range dep.Env {
			_, ok := profile.Env[k]
			if !ok {
				profile.Env[k] = v
			}
		}
	}

	profile.Deps = []string{}

	return profile
}
