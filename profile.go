/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"gopkg.in/yaml.v2"
)

type IfCond string

const (
	// SsdIfCond if machine has SSD disks.
	SsdIfCond = "ssd"
)

type ProfileKV struct {
	Value       string
	Description string   `json:",omitempty"`
	If          []string `json:",omitempty"`
}

type Profile struct {
	// Name of the profile
	Name string
	// Description of the service that is being updated.
	Description string `json:",omitempty"`
	// Documentation is the documentation for this profile
	Documentation string `json:",omitempty"`
	// ProcFS contains the kernel key values that will be changed.
	ProcFS map[string]ProfileKV `yaml:"procfs" json:",omitempty"`

	SysFS map[string]ProfileKV `yaml:"sysfs" json:",omitempty"`

	// Env is the environment variables that should be changed for
	// maximum performance.
	Env map[string]ProfileKV `json:",omitempty"`

	Vars map[string]interface{} `json:",omitempty"`

	// Deps of other profiles.
	Deps []string `json:",omitempty"`
}

func (p *Profile) PrintEnv() {
	for k, v := range p.Env {
		fmt.Printf("%s=%s\n", k, v.Value)
	}
}

func (p *Profile) PrintProcFS() {
	for k, v := range p.ProcFS {
		fmt.Printf("%s=%s\n", k, v.Value)
	}
}

func (p *Profile) PrintSysFS() {
	for k, v := range p.SysFS {
		fmt.Printf("%s=%s\n", k, v.Value)
	}
}

func (p *Profile) parseValueTemplates() {
	for _, valueMap := range []map[string]ProfileKV{
		p.ProcFS,
		p.SysFS,
		p.Env,
	} {
		for k, v := range valueMap {
			tmpl, err := template.New(p.Name + k).
				Funcs(template.FuncMap{
					"mul": func(param1 int, param2 int) int {
						return param1 * param2
					},
					"divide": func(param1 int, param2 int) int {
						return param1 / param2
					},
				}).
				Parse(v.Value)
			if err != nil {
				panic(err)
			}
			var b bytes.Buffer

			err = tmpl.Execute(&b, p)
			if err != nil {
				panic(err)
			}

			v.Value = b.String()
			valueMap[k] = v
		}
	}
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

func (p Profiles) GetWithDeps(s string) (profile *Profile) {
	for _, k := range p {
		if k.Name == s {
			profile = k
			break
		}
	}

	return
}

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

	profile.parseValueTemplates()

	return profile
}
