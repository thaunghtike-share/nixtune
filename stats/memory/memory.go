/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package memory

import (
	"time"
)

type Memory struct {
	Duration *time.Duration

	Total int64
	Free  int64
}

func New(duration *time.Duration) *Memory {
	return &Memory{
		Duration: duration,
	}
}

/*
https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html

Look to see how memory changes over time.

 - How many malloc allocations are done.
 - How quickly does the

*/
