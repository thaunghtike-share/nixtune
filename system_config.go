/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package main

// TODO: Need to be more intelligent. Update based on the system
// type. Defaulting to Ubuntu Linux for now.

import (
	"fmt"

	sig "github.com/anatma/autotune/signatures"
)

type SystemConfig struct{}

const (
	EnvFileName = "/etc/profile.d/99_anatma_autotune.sh"
)

func (sc *SystemConfig) updateEnv(config sig.SystemConfiger) {
	var (
		fileContent string
	)

	for k, v := range config.GetEnv() {
		envVar := fmt.Sprintf("%s=%s\n", k, v)
		fileContent += envVar

		logMe("INFO", envVar)
	}

	//	writeFile(EnvFileName, fileContent)
}

// TODO: Need to save the server state. Need
// to be able to update server state if config
// is bad.
func (sc *SystemConfig) updateSysctl(config sig.SystemConfiger) {
	for k, v := range config.GetSysctl() {
		sysVal := fmt.Sprintf("%s='%v'", k, v)
		logMe("INFO", sysVal)

		//		runCmd("sysctl", "-a", sysVal)
	}
}

// TODO
func (sc *SystemConfig) updateFiles(config sig.SystemConfiger) {
	// for k, v := range sc.config.GetFiles() {

	// }
}

func (sc *SystemConfig) Update(configs []sig.SystemConfiger) {
	for _, config := range configs {
		sc.updateEnv(config)
		sc.updateSysctl(config)
		sc.updateFiles(config)
	}
}
