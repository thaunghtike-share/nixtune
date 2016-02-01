/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

type ApacheConfig struct{}

// NewApacheConfig returns the configuration for the Apache HTTP
// Server.
//
// TODO: Eventually it should be split into apache2-mpm and
// apache2-fork.
func NewApacheConfig() *ApacheConfig {
	return &ApacheConfig{}
}

// GetEnv returns configurations Environment configurations.
func (c *ApacheConfig) GetEnv() map[string]string {
	return nil
}

// GetProcFS returns configurations for the kernel.
func (c *ApacheConfig) GetProcFS() map[string]string {
	nc := newNetworkConfig()
	proc := nc.GetProcFS()

	proc["kernel.sched_migration_cost_ns"] = "5000000"
	proc["kernel.sched_autogroup_enabled"] = "0"

	return proc
}

// GetSysFS returns configurations Environment configurations.
func (c *ApacheConfig) GetSysFS() map[string]string {
	return nil
}
