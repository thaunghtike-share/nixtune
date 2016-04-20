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
	"log"
	"os"
)

// Signature is the command used to update the system settings based
// on the profile specified by the user.
type Signature struct {
	Profile string

	Show     string
	ShowDeps bool

	Profiles Profiles

	// write dictates if we want to actually write the settings to
	// kernel.
	write bool
}

func (k *Signature) loadProfiles() {
	for _, i := range AssetNames() {
		ymlData, err := Asset(i)
		if err != nil {
			log.Fatal(err)
		}
		p := ParseProfile(ymlData)
		k.Profiles = append(k.Profiles, p)
	}
}

// Run gets the configuration for the profile and updates the system
// settings with the new values.
func (k *Signature) Run() int {
	var (
		profile *Profile
	)

	k.loadProfiles()
	if k.ShowDeps {
		profile = k.Profiles.Get(k.Profile)
	} else {
		profile = k.Profiles.GetWithDeps(k.Profile)
	}

	switch k.Show {
	case "env":
		profile.PrintEnv()
	case "procfs":
		profile.PrintProcFS()
	case "sysfs":
		profile.PrintSysFS()
	default:
		e, err := json.MarshalIndent(&profile, "", "  ")
		if err != nil {
			return -1
		}
		os.Stdout.Write(e)
	}

	// k.Config = k.Profiles[k.signature]

	// if k.Config == nil {
	// 	return -1
	// }

	// k.updateProcFS()
	// k.updateSysFS()

	// TODO: This is not quite ready yet.
	// k.updateEnv()

	return 0
}

// // TODO: This operation still doesn't happen and needs to be updated
// // accordingly.
// func (k *Signature) updateEnv() {
// 	var (
// 		fileContent string
// 	)

// 	for k, v := range k.Config.GetEnv() {
// 		envVar := fmt.Sprintf("%s=%s\n", k, v)
// 		fileContent += envVar

// 		log.Println("INFO", envVar)
// 	}

// 	//	writeFile(EnvFileName, fileContent)
// }

// func (k *Signature) updateProcFS() {
// 	for kernelKey, kernelVal := range k.Config.GetProcFS() {
// 		log.Println("INFO", fmt.Sprintf("%s From: '%v' To: '%v'", kernelKey, procfsGet(kernelKey), kernelVal))
// 		if k.write {
// 			procfsSet(kernelKey, kernelVal)
// 		}
// 	}
// }

// func (k *Signature) updateSysFS() {
// 	for kernelKey, kernelVal := range k.Config.GetSysFS() {
// 		log.Println("INFO", fmt.Sprintf("%s From: '%v' To: '%v'", kernelKey, sysfsGet(kernelKey), kernelVal))
// 		if k.write {
// 			sysfsSet(kernelKey, kernelVal)
// 		}
// 	}
// }

// NewSignature returns a new Signature object that we will use to
// update the system settings.
func NewSignature(profile, showFlag string, showDepsFlag bool) *Signature {
	s := &Signature{
		Profile:  profile,
		Show:     showFlag,
		ShowDeps: showDepsFlag,
	}

	return s
}
