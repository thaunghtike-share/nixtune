/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

type NginxConfig GolangConfig

// NewNginxConfig returns the configuration for Nginx servers.
func NewNginxConfig() *NginxConfig {
	return &NginxConfig{}
}
