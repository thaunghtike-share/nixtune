/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package fd

import (
	"time"
)

type FD struct {
	Duration *time.Duration
}

func New(duration *time.Duration) *FD {
	return &FD{
		Duration: duration,
	}
}
