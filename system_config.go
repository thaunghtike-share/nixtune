/*
 * Anatma Knight - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// TODO: Need to be more intelligent. Update based on the system
// type. Defaulting to Ubuntu Linux for now.

package main

import (
	"fmt"
)

type SystemConfiger interface {
	GetEnv() map[string]string
	GetSysctl() map[string]string
	GetFiles() map[string]FileChange
}

type FileChange struct {
	Content string
	Append  bool
}

type SystemConfig struct{}

const (
	EnvFileName = "/etc/profile.d/99_anatma_knight.sh"
)

func (sc *SystemConfig) updateEnv(config SystemConfiger) {
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
func (sc *SystemConfig) updateSysctl(config SystemConfiger) {
	for k, v := range config.GetSysctl() {
		sysVal := fmt.Sprintf("%s='%v'", k, v)
		logMe("INFO", sysVal)

		//		runCmd("sysctl", "-a", sysVal)
	}
}

// TODO
func (sc *SystemConfig) updateFiles(config SystemConfiger) {
	// for k, v := range sc.config.GetFiles() {

	// }
}

func (sc *SystemConfig) Update(config SystemConfiger) {
	sc.updateEnv(config)
	sc.updateSysctl(config)
	sc.updateFiles(config)
}
