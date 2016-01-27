/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

import (
	"fmt"
)

func sysctlGet(k string) string {
	return string(runCmdGetOutput("sysctl", "-b", k))
}

func sysctlSet(k, v string) string {
	return string(runCmdGetOutput("sysctl", "-w", fmt.Sprintf("%s=\"%s\"", k, v)))
}
