/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	// EnvFileName is the file where Environment variables will be
	// stored for settings defined by the various signatures.
	EnvFileName = "/etc/profile.d/99_acksin_autotune.sh"
)

// Signature is the command used to update the system settings based
// on the profile specified by the user.
type Signature struct {
	CmdName string `json:"-"`

	Profiles []Profile

	Signature string

	// write dictates if we want to actually write the settings to
	// kernel.
	write bool
}

func (k *Signature) Synopsis() string {
	return "Tune the kernel for server profile."

}

func (k *Signature) Help() string {
	var (
		sigs []string
	)

	for k, _ := range k.Profiles {
		sigs = append(sigs, fmt.Sprintf(" - %s", k))
	}
	sort.Strings(sigs)

	return fmt.Sprintf(`
Available signature profiles:

%s
`, strings.Join(sigs, "\n"))

}

// ParseArgs parses the commandline arguments passed for the Signature
// command.
// Run gets the configuration for the profile and updates the system
// settings with the new values.
func (k *Signature) Run(args []string) int {
	flags := flag.NewFlagSet(k.CmdName, flag.ContinueOnError)
	flags.BoolVar(&k.write, "write", true, "Write the settings.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}

	leftovers := flags.Args()
	if len(leftovers) == 0 {
		return -1
	}

	k.Signature = leftovers[0]

	ymlData, err := Asset("signatures/" + k.Signature + ".yml")
	if err != nil {
		log.Fatal(err)
	}
	ParseProfile(ymlData)

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
func NewSignature(cmdName string) *Signature {
	// New returns a map of profile and system configurations that are
	// currently supported.
	s := &Signature{
		CmdName: cmdName,
	}

	return s
}
