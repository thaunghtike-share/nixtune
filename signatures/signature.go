/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

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

// SystemConfiger is the interface which each signature should have to
// set the values in the environment and the kernel.
type SystemConfiger interface {
	GetEnv() map[string]string
	GetProcFS() map[string]string
	GetSysFS() map[string]string
}

// Signature is the command used to update the system settings based
// on the profile specified by the user.
type Signature struct {
	CmdName string

	// Config is configuration for the signature that we want to set.
	Config SystemConfiger

	// Signatures contains the name and configuration for each of the
	// server profiles.
	Profiles map[string]SystemConfiger

	// signature is the string representation of the signature we
	// want to get
	signature string
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

	k.signature = leftovers[0]

	k.Config = k.Profiles[k.signature]

	if k.Config == nil {
		return -1
	}

	k.updateProcFS()
	k.updateSysFS()

	// TODO: This is not quite ready yet.
	// k.updateEnv()

	return 0
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

		log.Println("INFO", envVar)
	}

	//	writeFile(EnvFileName, fileContent)
}

func (k *Signature) updateProcFS() {
	for kernelKey, kernelVal := range k.Config.GetProcFS() {
		log.Println("INFO", fmt.Sprintf("%s From: '%v' To: '%v'", kernelKey, procfsGet(kernelKey), kernelVal))
		if k.write {
			procfsSet(kernelKey, kernelVal)
		}
	}
}

func (k *Signature) updateSysFS() {
	for kernelKey, kernelVal := range k.Config.GetSysFS() {
		log.Println("INFO", fmt.Sprintf("%s From: '%v' To: '%v'", kernelKey, sysfsGet(kernelKey), kernelVal))
		if k.write {
			sysfsSet(kernelKey, kernelVal)
		}
	}
}

// NewSignature returns a new Signature object that we will use to
// update the system settings.
func New(cmdName string) *Signature {
	// New returns a map of profile and system configurations that are
	// currently supported.
	s := &Signature{
		CmdName: cmdName,
	}

	s.Profiles = make(map[string]SystemConfiger)

	// Async Server configurations
	s.Profiles["golang"] = NewGolangConfig()
	s.Profiles["nodejs"] = NewNodejsConfig()
	s.Profiles["nginx"] = NewNginxConfig()
	s.Profiles["haproxy"] = NewHaproxyConfig()

	// Forking server configurations
	s.Profiles["apache"] = NewApacheConfig()
	s.Profiles["postgresql"] = NewPostgresqlConfig()

	s.Profiles["java"] = NewJavaConfig()

	return s
}
