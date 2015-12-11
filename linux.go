/*
 * Anatma Autotune - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"github.com/abhiyerra/procfs"
)

type Memory struct {
	Total uint
}

func Compute() (m *Memory) {
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return
	}

	return
}

// getProcessList()
// guessServerProfile()
// numberOfLogins()
// mainProcess()

/*

 The best way to figure ot what processes are there is a way to look
 for certain process names.

 - Also can look at the file itself to see how it is constructed.
 - Can usually tell based on interpreter
 - Have to learn to guess executables.

*/

// http://techblog.netflix.com/2015/11/linux-performance-analysis-in-60s.html
