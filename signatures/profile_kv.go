/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package signatures

type ProfileKV struct {
	// Value for the Key associated with the ProfileKV. Note that
	// this is also used as a Golang text/template which gets
	// updated.
	Value string
	// Description of the Value and this property.
	Description string `json:",omitempty"`
	// Schedule is used for Cron tasks
	Schedule string `json:",omitempty"`
	// Default Value if it isn't specified.
	Default string
}
