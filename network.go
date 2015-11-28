/*
 * Anatma Knight - Kernel Autotuning
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

/*

Network Access ALgorithm for Anatma Knight

To be able to do a profile on network throughput we need to first take
a metric of teh connection.

 - Need to look at network connections over a period of time.
 - See if the connections are to the same pid of different ones.
 - See how many timeouts there are.
 - Look at how many connections there are that are open and how tey are operating.
 - See if there is an increase see the throughput as it grows over time.
 - Need to profile the machine over time. Over a long period of time.
*/

// getNetworkSettings()
