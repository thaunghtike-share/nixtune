/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

// NewApacheConfig returns the configuration for the Apache HTTP
// Server.
//
// TODO: Eventually it should be split into apache2-mpm and
// apache2-fork.
func NewApacheConfig() *PostgresqlConfig {
	return &PostgresqlConfig{}
}
