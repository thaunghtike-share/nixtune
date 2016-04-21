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
	"flag"
	"fmt"
	"log"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"

	"github.com/acksin/strum/stats"
)

var (
	ProfileFuncMaps = template.FuncMap{
		"multiply": func(param1 int, param2 int) int {
			return param1 * param2
		},
		"divide": func(param1 int, param2 int) int {
			return param1 / param2
		},
	}
)

type Profile struct {
	// Name of the profile
	Name string
	// Description of the service that is being updated.
	Description string `json:",omitempty"`
	// Documentation is the documentation for this profile
	Documentation string `json:",omitempty"`
	// References are places we got this information.
	References []string `json:",omitempty"`

	// ProcFS contains the /proc filesystem variables.
	ProcFS map[string]ProfileKV `yaml:"procfs" json:",omitempty"`
	// SysFS contains the /sys filesystem variables.
	SysFS map[string]ProfileKV `yaml:"sysfs" json:",omitempty"`
	// Env is the environment variables that will be changed
	Env map[string]ProfileKV `json:",omitempty"`
	// Files that need to be modified for specific tuning.
	Files map[string]map[string]ProfileKV `json:",omitempty"`
	// Cron jobs that should be run to optimize performance.
	Cron map[string]ProfileKV `json:",omitempty"`
	// Flags are values that are passed from the command line to
	// be used by the Profile.
	Flags map[string]*ProfileKV `json:",omitempty"`

	// Vars are the variables passed to modify the signature
	// templates. These can be used to pass values to ProcFS,
	// SysFS and Env.
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

func (p *Profile) ParseFlags(args []string) error {
	if p.Flags == nil {
		return nil
	}

	flags := flag.NewFlagSet(p.Name, flag.ContinueOnError)
	for k, v := range p.Flags {
		flags.StringVar(&v.Value, k, v.Default, v.Description)
	}

	if err := flags.Parse(args); err != nil {
		return err
	}

	return nil
}

func (p *Profile) parseValueTemplates() {
	var (
		s = stats.New([]int{})
	)

	for _, valueMap := range []map[string]ProfileKV{
		p.ProcFS,
		p.SysFS,
		p.Env,
	} {
		for k, v := range valueMap {
			varStruct := struct {
				Vars  map[string]interface{}
				Stats *stats.Stats
			}{
				Vars:  p.Vars,
				Stats: s,
			}

			tmpl, err := template.New(p.Name + k).
				Funcs(ProfileFuncMaps).
				Parse(v.Value)
			if err != nil {
				panic(err)
			}
			var b bytes.Buffer

			err = tmpl.Execute(&b, &varStruct)
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

func (p Profiles) Get(s string, withDeps bool) (profile *Profile) {
	if withDeps {
		return p.getWithDeps(s)
	}

	return p.getWithoutDeps(s)
}

func (p Profiles) getWithDeps(s string) (profile *Profile) {
	for _, k := range p {
		if k.Name == s {
			profile = k
			break
		}
	}

	profile.parseValueTemplates()

	return
}

func (p Profiles) getWithoutDeps(s string) *Profile {
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
		depProfiles = append(depProfiles, p.getWithDeps(dep))
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

func loadProfiles() {
	for _, i := range AssetNames() {
		ymlData, err := Asset(i)
		if err != nil {
			log.Fatal(err)
		}

		p := ParseProfile(ymlData)

		switch {
		case strings.HasPrefix(i, "signatures/open"):
			profiles = append(profiles, p)
		case strings.HasPrefix(i, "signatures/pro") && currentSubscription == ProSubscription || currentSubscription == PremiumSubscription:
			profiles = append(profiles, p)
		case strings.HasPrefix(i, "signatures/premium") && currentSubscription == PremiumSubscription:
			profiles = append(profiles, p)
		}
	}
}
