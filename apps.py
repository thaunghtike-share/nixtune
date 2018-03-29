# type Apache struct{}

# func (f *Apache) GetProfile() *Profile {
# 	p := &Profile{
# 		Name:          "apache",
# 		Subscription:  StartupSubscription,
# 		Description:   "Configuration for the Apache HTTP Server",
# 		Documentation: "Apache",
# 		Deps:          []Profiler{&FastServer{}},
# 		ProcFS:        f.procfs(),
# 	}

# 	return p
# }

# func (f *Apache) procfs() map[string]*ProfileKV {
# 	p := make(map[string]*ProfileKV)

# 	p["kernel.sched_migration_cost_ns"] = &ProfileKV{
# 		Value:       "5000000",
# 		Description: "Since 2.6.32 Linux kernel is using the Completely Fair Scheduler but in a machine running forked processes you may run into problems. This parameter determines how long a migrated process has to be running before the kernel will consider migrating it again to another core. Apache scales much higher if this number is higher.",
# 	}

# 	p["kernel.sched_autogroup_enabled"] = &ProfileKV{
# 		Value: "0",
# 	}

# 	return p
# }

# package autotune

# type Golang struct{}

# func (f *Golang) GetProfile() *Profile {
# 	p := &Profile{
# 		Name:         "golang",
# 		Subscription: StartupSubscription,
# 		Description:  "Configuration for high throughput Golang apps",
# 		Documentation: `# Linux Optimizations for High Throughput Golang Apps

# Go applications have unique characteristics which require certain
# Linux kernel tuning to achieve high throughput.

# ## Go's Utilization Profile

# CPU will not be a bottleneck with Golang applications. Our research
# shows that applications, even those that utilize CGO, do no see CPU be
# a bottleneck. The places where performance become bottlenecks are the
# following:

#  - Garbage Collection
#  - Default ulimits
#  - Networking

# ## Assumptions

# We will be under the assumption that there will be one primary Go
# application running on the machine and can have access to all of the
# resources. We also assume that we want high network throughput as the
# goal is to have high response rate. We want to be able to handle
# millions of requests.

# ## GC Optimizations

# For all intents and purposes we should be able to increase the GOGC to
# a number based on the size of the machine. If I am using a m4.large
# instance on Amazon I use GOGC=10000. The higher the GOGC value the
# less frequent the Garbage Collection will run. Further, since we are
# optimizing the server to be heavily utilized for a primary Golang
# service we want to use up all the RAM available to us.

# ## Ulimits

# Ulimits are a security mechanism in POSIX based systems which gives
# each user a certain amount of allocation of various
# resources. However, the resource we are concerned with is file
# descriptors. (ulimit -n) Since a file descriptor can be a file or a
# socket we can quickly saturate how many connections an app not running
# as root can use. Further, the default open files ulimit on an Ubuntu
# Server 14.04 are ridiculously low at 1024.

# The server will reach network saturation quickly if this is not dealt
# with. Further, since we want to optimize for the single Golang
# application we will give every user on the Linux machine unlimited
# open files.`,
# 		Deps: []Profiler{&FastServer{}},
# 		References: []string{
# 			"http://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables",
# 		},
# 	}

# 	p.Vars = make(map[string]interface{})
# 	p.Vars["nfConntrackMax"] = 200000

# 	p.Env = make(map[string]*ProfileKV)

# 	// This value changes with the process that is running. It
# 	// needs to modified based on available ram and process since
# 	// we will be looking at the amount of memory that the process
# 	// itself is taking.
# 	p.Env["GOGC"] = &ProfileKV{
# 		Value:       "500",
# 		Description: `Set the value of GOGC to be really high. Step up by 100 for each additional GB of ram. Can likely be higher but it is a value that needs to be modified based on the heap used so it is better done within the Go application itself.`,
# 	}

# 	return p
# }
# # Copyright (C) 2016 opszero <hey@opszero.com>
# #
# # This Source Code Form is subject to the terms of the Mozilla Public
# # License, v. 2.0. If a copy of the MPL was not distributed with this
# # file, You can obtain one at http://mozilla.org/MPL/2.0/.

# from decorators import *
# import mental_model

# package autotune

# import (
# 	"fmt"
# 	"io/ioutil"
# 	"regexp"
# )

# type Nginx struct {
# 	p *Profile
# }

# func (f *Nginx) GetProfile() *Profile {
# 	f.p = &Profile{
# 		Name:         "nginx",
# 		Subscription: StartupSubscription,
# 		Description:  "Configuration for Nginx.",
# 		Deps:         []Profiler{&FastServer{}},
# 		References: []string{
# 			"https://www.digitalocean.com/community/tutorials/how-to-optimize-nginx-configuration",
# 		},
# 	}

# 	f.p.Flags = make(map[string]*ProfileKV)
# 	f.p.Flags["nginx-conf"] = &ProfileKV{
# 		Description: "Location of the nginx.conf file.",
# 		Default:     "/etc/nginx/nginx.conf",
# 	}

# 	return f.p
# }

# func (f *Nginx) workerProcessesRegex() *regexp.Regexp {
# 	return regexp.MustCompile(`^worker_processes ([0-9]*|auto);$`)
# }

# func (f *Nginx) confWorkerProcesses() *ProfileKV {
# 	configFilename := f.p.GetFlag("nginx-conf")

# 	b, err := ioutil.ReadFile(configFilename)
# 	if err != nil {
# 		return nil
# 	}

# 	fmt.Println(f.workerProcessesRegex().FindAllStringSubmatch(string(b), -1))

# 	return nil
# }

# // func (f *Nginx) workerConnectionsRegex() *regexp.Regexp {
# // 	return regexp.MustCompile(`^worker_processes ([0-9]*|auto);$`)
# // }

# // func (f *Nginx) confWorkerConnections() *ProfileKV {
# // 	configFilename := f.p.GetFlag("nginx-conf")

# // 	b, err := ioutil.ReadFile(configFilename)
# // 	if err != nil {
# // 		return nil
# // 	}

# // 	regexp.MustCompile(`^worker_processes (\d*|auto);$`)

# // 	matched, err := regexp.MatchString("foo.*", "seafood")

# // 	return nil
# // }

# func (f *Nginx) conf() map[string]*ProfileKV {
# 	p := make(map[string]*ProfileKV)

# 	p["worker_processes"] = f.confWorkerProcesses()
# 	// p["worker_connections"] = f.confWorkerConnections()

# 	return p
# }
# package autotune

# import (
# 	"testing"
# )

# func TestNginx_workerProcessesRegex(t *testing.T) {
# 	n := &Nginx{}
# 	re := n.workerProcessesRegex()

# 	a := re.FindAllStringSubmatch("worker_processes 1;", -1)
# 	if len(a) != 1 && len(a[0]) != 2 && a[0][1] != "1" {
# 		t.Errorf("Invalid number of worker_processes")
# 	}

# 	b := re.FindAllStringSubmatch("worker_processes auto;", -1)
# 	if len(b) != 1 && len(b[0]) != 2 && b[0][1] != "auto" {
# 		t.Errorf("Invalid number of worker_processes")
# 	}

# 	if len(re.FindAllStringSubmatch("worker_processes 1", -1)) > 0 {
# 		t.Errorf("Invalid number of worker_processes")
# 	}

# 	if len(re.FindAllStringSubmatch("none", -1)) > 0 {

# 		t.Errorf("Invalid number of worker_processes")
# 	}
# }
# package autotune

# type PostgreSQL struct{}

# func (f *PostgreSQL) GetProfile() *Profile {
# 	p := &Profile{
# 		Name:         "postgresql",
# 		Subscription: StartupSubscription,
# 		Description:  "PostgreSQL optimizations",
# 		References: []string{
# 			"http://www.postgresql.org/message-id/50E4AAB1.9040902@optionshouse.com",
# 			"http://www.postgresql.org/docs/9.1/static/kernel-resources.html",
# 		},
# 		ProcFS: f.procfs(),
# 	}

# 	return p
# }

# func (f *PostgreSQL) procfs() (p map[string]*ProfileKV) {
# 	p = make(map[string]*ProfileKV)

# 	p["kernel.sched_migration_cost_ns"] = &ProfileKV{
# 		Value: "5000000",
# 	}

# 	p["kernel.sched_autogroup_enabled"] = &ProfileKV{
# 		Value: "0",
# 	}

# 	p["kernel.shmmax"] = &ProfileKV{
# 		Value: "17179869184",
# 	}

# 	p["kernel.shmall"] = &ProfileKV{
# 		Value: "4194304",
# 	}

# 	return p
# }

