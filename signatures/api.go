/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

// SystemConfiger is the interface which each signature should have to
// set the values in the environment and the kernel.
type SystemConfiger interface {
	GetEnv() map[string]string
	GetSysctl() map[string]string
}

// Signatures contains the name and configuration for each of the
// server profiles.
type Signatures map[string]SystemConfiger

// New returns a map of profile and system configurations that are
// currently supported.
func New() (serverConfigs Signatures) {
	serverConfigs = make(Signatures)

	// Async Server configurations
	serverConfigs["golang"] = NewGolangConfig()
	serverConfigs["nodejs"] = NewNodejsConfig()
	serverConfigs["nginx"] = NewNginxConfig()
	serverConfigs["haproxy"] = NewHaproxyConfig()

	// Forking server configurations
	serverConfigs["apache"] = NewApacheConfig()
	serverConfigs["postgresql"] = NewPostgresqlConfig()

	serverConfigs["java"] = NewJavaConfig()

	return
}

// Types returns the slice of the different configurations that we
// have available.
func (s Signatures) Types() (types []string) {
	for k := range s {
		types = append(types, k)
	}

	return
}

// Get returns the kernel configuration for a profile.
func (s Signatures) Get(sig string) SystemConfiger {
	return s[sig]
}
