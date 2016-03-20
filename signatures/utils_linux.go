/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

import (
	"fmt"
)

func procfsGet(k string) string {
	return string(runCmdGetOutput("sysctl", "-b", k))
}

func procfsSet(k, v string) string {
	return string(runCmdGetOutput("sysctl", "-w", fmt.Sprintf("%s=\"%s\"", k, v)))
}

func sysfsGet(k string) string {
	return string(runCmdGetOutput("cat", k))
}

func sysfsSet(k, v string) string {
	return string(runCmdGetOutput("echo", v, ">", k))
}
