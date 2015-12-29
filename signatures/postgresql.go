/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package signatures

type PostgresqlConfig struct{}

func NewPostgresqlConfig() *PostgresqlConfig {
	return &PostgresqlConfig{}
}

func (c *PostgresqlConfig) GetEnv() map[string]string {
	return nil
}

func (c *PostgresqlConfig) GetSysctl() map[string]string {
	sysctl := make(map[string]string)

	sysctl["kernel.sched_migration_cost_ns"] = "5000000"
	sysctl["kernel.sched_autogroup_enabled"] = "0"

	return sysctl
}

func (c *PostgresqlConfig) GetFiles() map[string]FileChange {
	return nil
}
