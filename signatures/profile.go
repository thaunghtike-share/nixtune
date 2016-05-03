/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

import (
	"bytes"
	"flag"
	"fmt"
	"sort"
	"text/template"

	"github.com/acksin/strum/stats"
)

var (
	ProfileFuncMaps = template.FuncMap{
		"multiply": func(param1 int64, param2 int64) int64 {
			return param1 * param2
		},
		"multiplyInt": func(param1 int, param2 int) int {
			return param1 * param2
		},
		"divide": func(param1 int, param2 int) int {
			return param1 / param2
		},
	}
)

type Profiler interface {
	GetProfile() *Profile
}

type Profile struct {
	// Name of the profile
	Name string
	// Plan that this signature belongs to.
	Subscription Subscription
	// Description of the service that is being updated.
	Description string `json:",omitempty"`
	// Documentation is the documentation for this profile
	Documentation string `json:",omitempty"`
	// References are places we got this information.
	References []string `json:",omitempty"`

	// ProcFS contains the /proc filesystem variables.
	ProcFS map[string]*ProfileKV `yaml:"procfs" json:",omitempty"`
	// SysFS contains the /sys filesystem variables.
	SysFS map[string]*ProfileKV `yaml:"sysfs" json:",omitempty"`
	// Env is the environment variables that will be changed
	Env map[string]*ProfileKV `json:",omitempty"`
	// Files that need to be modified for specific tuning.
	Files map[string]*ProfileKV `json:",omitempty"`
	// Cron jobs that should be run to optimize performance.
	Cron map[string]*ProfileKV `json:",omitempty"`
	// Flags are values that are passed from the command line to
	// be used by the Profile.
	Flags map[string]*ProfileKV `json:",omitempty"`
	// Conf looks at the application config for changes that need
	// to be done.
	Conf map[string]*ProfileKV `json:",omitempty"`

	// Vars are the variables passed to modify the signature
	// templates. These can be used to pass values to ProcFS,
	// SysFS and Env.
	Vars map[string]interface{} `json:",omitempty"`
	// Deps of other profiles.
	Deps []Profiler `json:",omitempty"`
}

func (p *Profile) GetFlag(flagKey string) string {
	k, ok := p.Flags[flagKey]
	if !ok {
		return ""
	}

	if k.Value != "" {
		return k.Value
	}

	return k.Default
}

func (p *Profile) printMap(m map[string]*ProfileKV) {
	if m == nil {
		return
	}

	var s []string

	for k := range m {
		s = append(s, k)
	}

	sort.Strings(s)

	for _, k := range s {
		fmt.Printf("%s=%s\n", k, m[k].Value)
	}
}

func (p *Profile) PrintFiles() {
	if p.Files == nil {
		return
	}

	var s []string
	for k := range p.Files {
		s = append(s, k)
	}

	sort.Strings(s)

	for _, k := range s {
		fmt.Printf("%s:\n", k)
		fmt.Printf("%s\n", p.Files[k].Value)
	}
}

func (p *Profile) PrintConf() {
	p.printMap(p.Conf)
}

func (p *Profile) PrintEnv() {
	p.printMap(p.Env)
}

func (p *Profile) PrintProcFS() {
	p.printMap(p.ProcFS)
}

func (p *Profile) PrintSysFS() {
	p.printMap(p.SysFS)
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

	for _, valueMap := range []map[string]*ProfileKV{
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

func (p *Profile) HasSubscription(currentSub Subscription) bool {
	switch p.Subscription {
	case OpenSubscription:
		if currentSub == OpenSubscription {
			return true
		}
	case StartupSubscription:
		if currentSubscription == StartupSubscription || currentSubscription == ProSubscription || currentSubscription == PremiumSubscription {
			return true
		}
	case ProSubscription:
		if currentSubscription == ProSubscription || currentSubscription == PremiumSubscription {
			return true
		}
	case PremiumSubscription:
		if currentSubscription == PremiumSubscription {
			return true
		}
	}

	return false
}
