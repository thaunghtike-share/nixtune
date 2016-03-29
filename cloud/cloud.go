/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package cloud

type Cloud struct {
	AWS *AWSStats
}

func New() (c *Cloud) {
	c = &Cloud{
		AWS: NewAWS(),
	}

	return c
}
