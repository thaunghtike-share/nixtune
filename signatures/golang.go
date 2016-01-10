/* Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package signatures

// ##  Linux Optimizations for High Throughput Golang Apps
//
// Go applications have unique characteristics which require certain
// Linux kernel tuning to achieve high throughput.
//
// ### Go\'s utilization Profile
//
// CPU will not be a bottleneck with Golang applications. Our research
// shows that applications, even those that utilize CGO, do no see CPU be
// a bottleneck. The places where performance become bottlenecks are the
// following:
//
//  - Garbage Collection
//  - Default ulimits
//  - Networking
//
// ## Assumptions
//
// We will be under the assumption that there will be one primary Go
// application running on the machine and can have access to all of the
// resources. We also assume that we want high network throughput as the
// goal is to have high response rate. We want to be able to handle
// millions of requests.
//
// ## Prescription
//
// ### GC Optimizations
//
// For all intents and purposes we should be able to increase the GOGC to
// a number based on the size of the machine. If I am using a m4.large
// instance on Amazon I use GOGC=10000. The higher the GOGC value the
// less frequent the Garbage Collection will run. Further, since we are
// optimizing the server to be heavily utilized for a primary Golang
// service we want to use up all the RAM available to us.
//
// ### Ulimits
//
// Ulimits are a security mechanism in POSIX based systems which gives
// each user a certain amount of allocation of various
// resources. However, the resource we are concerned with is file
// descriptors. (ulimit -n) Since a file descriptor can be a file or a
// socket we can quickly saturate how many connections an app not running
// as root can use. Further, the default open files ulimit on an Ubuntu
// Server 14.04 are ridiculously low at 1024.
//
// The server will reach network saturation quickly if this is not dealt
// with. Further, since we want to optimize for the single Golang
// application we will give every user on the Linux machine unlimited
// open files.
//
// ### Networking
//
// https://engineering.gosquared.com/optimising-nginx-node-js-and-networking-for-heavy-workloads
type GolangConfig struct{}

func NewGolangConfig() *GolangConfig {
	return &GolangConfig{}
}

func (c *GolangConfig) GetEnv() map[string]string {
	env := make(map[string]string)

	// Set the value of GOGC to be really high.

	// TODO: Consider how this is being used as part of a bigger
	// setting. Based on RAM etc.

	env["GOGC"] = "2000"

	return env
}

func (c *GolangConfig) GetSysctl() map[string]string {
	nc := &NetworkConfig{}
	return nc.GetSysctl()
}
