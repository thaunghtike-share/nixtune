/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

// PostgresqlConfig defines the interface for PostgreSQL configration.
//
// http://www.postgresql.org/message-id/50E4AAB1.9040902@optionshouse.com
// http://www.postgresql.org/docs/9.1/static/kernel-resources.html
type PostgresqlConfig struct{}

// NewPostgresqlConfig returns the config for PostgreSQL.
func NewPostgresqlConfig() *PostgresqlConfig {
	return &PostgresqlConfig{}
}

// GetEnv returns configurations Environment configurations.
func (c *PostgresqlConfig) GetEnv() map[string]string {
	return nil
}

// GetSysctl returns configurations for the kernel.
func (c *PostgresqlConfig) GetSysctl() map[string]string {
	sysctl := make(map[string]string)

	sysctl["kernel.sched_migration_cost_ns"] = "5000000"
	sysctl["kernel.sched_autogroup_enabled"] = "0"

	sysctl["kernel.shmmax"] = "17179869184"
	sysctl["kernel.shmall"] = "4194304"

	return sysctl
}
