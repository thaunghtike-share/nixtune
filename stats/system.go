/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package stats

import (
	"github.com/acksin/strum/memory"
	"github.com/acksin/strum/network"
)

// System contains information about the system
type System struct {
	// Memory stats of the system
	Memory *memory.Memory
	// Network stats of the system
	Network *network.Network
}
