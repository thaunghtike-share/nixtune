/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package signatures

type GolangConfig struct{}

func NewGolangConfig() *GolangConfig {
	return &GolangConfig{}
}

func (c *GolangConfig) GetEnv() map[string]string {
	env := make(map[string]string)

	// Set the value of GOGC to be really high.

	// TODO: Consider how this is being used as part of a bigger
	// setting. Based on RAM etc.

	env["GOGC"] = "2000"

	return env
}

func (c *GolangConfig) GetSysctl() map[string]string {
	nc := &NetworkConfig{}
	return nc.GetSysctl()
}
