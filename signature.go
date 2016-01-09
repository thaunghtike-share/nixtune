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
)

const (
	EnvFileName = "/etc/profile.d/99_anatma_autotune.sh"
)

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
	sigs := ""
	for _, i := range sig.NewSignatures().Types() {
		sigs += fmt.Sprintf(" - %s\n", i)
	}

	fmt.Fprintf(os.Stderr, `
Available signature profiles:

%s
`, sigs)

}

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
	for k, v := range k.Config.GetSysctl() {
		logMe("INFO", fmt.Sprintf("%s From: '%v' To: '%v'", k, sysctlGet(k), v))
		runCmd("sysctl", "-w", fmt.Sprintf("%s='%v'", k, v))
	}
}

func (k *Signature) Run() error {
	sigs := sig.NewSignatures()

	k.Config = sigs.Get(k.signature)
	if k.Config == nil {
		k.usage()
	}

	k.updateSysctl()

	// TODO: This is not quite ready yet.
	//	k.updateEnv()

	return nil
}

func NewSignature() *Signature {
	return &Signature{}
}
