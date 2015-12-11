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

type Memory struct {
	Total uint

	/*

		MemTotal:       15666184 kB
		MemFree:         1373032 kB
		Buffers:          958356 kB
		Cached:         11171148 kB
		SwapCached:            0 kB
		Active:          6450528 kB
		Inactive:        5992220 kB
		Active(anon):     313820 kB
		Inactive(anon):      280 kB
		Active(file):    6136708 kB
		Inactive(file):  5991940 kB
		Unevictable:           0 kB
		Mlocked:               0 kB
		SwapTotal:             0 kB
		SwapFree:              0 kB
		Dirty:               148 kB
		Writeback:             0 kB
		AnonPages:        313244 kB
		Mapped:            44500 kB
		Shmem:               856 kB
		Slab:            1723520 kB
		SReclaimable:    1663564 kB
		SUnreclaim:        59956 kB
		KernelStack:        1616 kB
		PageTables:         5184 kB
		NFS_Unstable:          0 kB
		Bounce:                0 kB
		WritebackTmp:          0 kB
		CommitLimit:     7833092 kB
		Committed_AS:     587008 kB
		VmallocTotal:   34359738367 kB
		VmallocUsed:       36660 kB
		VmallocChunk:   34359636840 kB
		HardwareCorrupted:     0 kB
		AnonHugePages:     40960 kB
		HugePages_Total:       0
		HugePages_Free:        0
		HugePages_Rsvd:        0
		HugePages_Surp:        0
		Hugepagesize:       2048 kB
		DirectMap4k:       91136 kB
		DirectMap2M:    16039936 kB
	*/
}

func Compute() (m *Memory) {
	m = &Memory{}

	"cat /proc/meminfo"

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
