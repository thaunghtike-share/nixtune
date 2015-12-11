/*
 * Anatma Autotune - Kernel Autotuning
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
