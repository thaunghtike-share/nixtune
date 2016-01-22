/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"flag"
	"fmt"
	sig "github.com/anatma/autotune/signatures"
	"os"
	"sort"
	"strings"
)

const (
	// EnvFileName is the file where Environment variables will be
	// stored for settings defined by the various signatures.
	EnvFileName = "/etc/profile.d/99_anatma_autotune.sh"
)

// Signature is the command used to update the system settings based
// on the profile specified by the user.
type Signature struct {
	// Config is configuration for the signature that we want to set.
	Config sig.SystemConfiger
	// signature is the string representation of the signature we
	// want to get
	signature string
	// write dictates if we want to actually write the settings to
	// kernel.
	write bool
}

func (k *Signature) usage() {
	var (
		sigs []string
	)

	for _, i := range sig.New().Types() {
		sigs = append(sigs, fmt.Sprintf(" - %s", i))
	}
	sort.Strings(sigs)

	fmt.Fprintf(os.Stderr, `
Available signature profiles:

%s
`, strings.Join(sigs, "\n"))

}

// ParseArgs parses the commandline arguments passed for the Signature
// command.
func (k *Signature) ParseArgs(args []string) {
	flags := flag.NewFlagSet(subCmd("signature"), flag.ContinueOnError)
	flags.BoolVar(&k.write, "write", true, "Write the settings.")

	if err := flags.Parse(args); err != nil {
		os.Exit(-1)
	}

	leftovers := flags.Args()
	if len(leftovers) == 0 {
		fmt.Fprintf(os.Stderr, "Usage of %s signature [profile]:\n", CmdName)
		flags.PrintDefaults()
		k.usage()
		os.Exit(-1)
	}

	k.signature = leftovers[0]
}

// TODO: This operation still doesn't happen and needs to be updated
// accordingly.
func (k *Signature) updateEnv() {
	var (
		fileContent string
	)

	for k, v := range k.Config.GetEnv() {
		envVar := fmt.Sprintf("%s=%s\n", k, v)
		fileContent += envVar

		logMe("INFO", envVar)
	}

	//	writeFile(EnvFileName, fileContent)
}

// Goes through the sysctl changes and updates them. They are not
// written to disk so if things fail we can just restart the
// machine.
func (k *Signature) updateSysctl() {
	for kernelKey, kernelVal := range k.Config.GetSysctl() {
		logMe("INFO", fmt.Sprintf("%s From: '%v' To: '%v'", kernelKey, sysctlGet(kernelKey), kernelVal))
		if k.write {
			sysctlSet(kernelKey, kernelVal)
		}
	}
}

// Run gets the configuration for the profile and updates the system
// settings with the new values.
func (k *Signature) Run() error {
	sigs := sig.New()

	k.Config = sigs.Get(k.signature)
	if k.Config == nil {
		k.usage()
	}

	k.updateSysctl()

	// TODO: This is not quite ready yet.
	// k.updateEnv()

	return nil
}

// NewSignature returns a new Signature object that we will use to
// update the system settings.
func NewSignature() *Signature {
	return &Signature{}
}
