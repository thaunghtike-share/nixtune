# Acksin

<a href="https://travis-ci.org/acksin/acksin"><img src="https://travis-ci.org/acksin/acksin.svg?branch=master" /></a>
<a href="https://godoc.org/github.com/acksin/acksin"><img src="https://godoc.org/github.com/acksin/acksin?status.svg" alt="GoDoc"></a>

## Introduction

Acksin is Automatic Transmission for Linux. It attempts to performance
enhance your Linux server based on the type of services that you run
on it making sure that Linux is pulling its weight.

## [Quick Start and Download](https://www.acksin.com/#quickstart)

Acksin outputs its data in JSON to the command line. Run the
following command:

    sudo acksin output

Acksin primarily runs as a daemon which regularily pushes diagnostics
to a central server. Acksin runs a service called
[Acksin Console](https://www.acksin.com/console/login?redirectTo=https://www.acksin.com/console/)
providing this capability. You can get the configuration on the Acksin
Console or you can check out
[config.json.template](config.json.template) for agent
configuration. We will open source the server side in the near future.

Run the following:

    sudo acksin agent config.json

## Getting Started & Documentation

All documentation is on the [Acksin website](https://www.acksin.com/).

## Developing Acksin

Acksin's command line portion is primarily written in Go whereas the
Machine Learning is written in Python. The code is split into a
couple different sections:

 - [Command Line Tool](stats): Collects stats from the System and
   Containers.
 - [Mental Models](ai/mental_models): Take System stats and creates
   models for AI. Currently this is a program that runs on AWS Lambda.

### Primary Dependencies

The primary dependency of Acksin is the
[ProcFS Library](https://github.com/acksin/procfs) we use.  Any code
that needs to read from ProcFS should go there. Most of the Command
Line App is a wrapper for that library. In the future we will have
similar dependencies for SysFS. In addition to that we use the Go
libraries provided by the Cloud providers.

### Deploying Acksin

Acksin has several components that need to be deployed to make a
complete system. This includes the Command Line Tool, Mental Models,
AI, Console, and Server.

#### Command Line Tool

To build the command line tool run the following:

```
make deps
make build
```

#### Mental Models

Currently to run the Mental Models you need to install the
[Serverless Framework](https://www.serverless.com) as they run on AWS
Lambda. We will split this out so it doesn't rely on Lambda in the
near future.

Make sure to create a `ai/mental_models/serverless.env.yaml` file.

Example:
```
vars: null
stages:
    dev:
        vars: null
        regions:
            us-west-2:
    prod:
        vars: null
        regions:
            us-west-2:
```

To deploy this function run.

```
cd ai
make deps
make dev
```


### [Contributing](CONTRIBUTING.md)

We love contributors to the project. Please check out the
[CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

Copyright (C) 2016 Acksin <hey@acksin.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at <http://mozilla.org/MPL/2.0/>.
