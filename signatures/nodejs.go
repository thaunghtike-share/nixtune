/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

type NodejsConfig GolangConfig

// NewNodejsConfig returns the configuration for Node.js servers.
func NewNodejsConfig() *NodejsConfig {
	return &NodejsConfig{}
}
