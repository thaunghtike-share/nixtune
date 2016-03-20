/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
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

// GetProcFS returns configurations for the kernel.
func (c *PostgresqlConfig) GetProcFS() map[string]string {
	proc := make(map[string]string)

	proc["kernel.sched_migration_cost_ns"] = "5000000"
	proc["kernel.sched_autogroup_enabled"] = "0"

	proc["kernel.shmmax"] = "17179869184"
	proc["kernel.shmall"] = "4194304"

	return proc
}

// GetSysFS returns configurations Environment configurations.
func (c *PostgresqlConfig) GetSysFS() map[string]string {
	return nil
}
