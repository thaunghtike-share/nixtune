/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package signatures

type ServerType int

const (
	// Async and High Network Throughput
	GolangServer ServerType = iota
	NodejsServer
	NginxServer

	// Forking Servers
	ApacheServer
	PostgresqlServer

	JavaServer

	Unknown
)

type SystemConfiger interface {
	GetEnv() map[string]string
	GetSysctl() map[string]string
	GetFiles() map[string]FileChange
}

type FileChange struct {
	Content string
	Append  bool
}

func ServerSignature(signature string) ServerType {
	switch signature {
	case "golang":
		return GolangServer
	case "nodejs":
		return NodejsServer
	case "nginx":
		return NginxServer
	case "apache":
		return ApacheServer
	}

	return Unknown
}

func Configs(signature string) (configs []SystemConfiger) {
	switch ServerSignature(signature) {
	case GolangServer:
		configs = append(configs, NewNetworkConfig(HighNetworkLevel))
		configs = append(configs, NewGolangConfig())
	case NodejsServer:
		configs = append(configs, NewNetworkConfig(HighNetworkLevel))
		configs = append(configs, NewGolangConfig())
	case NginxServer:
		configs = append(configs, NewNetworkConfig(HighNetworkLevel))
		configs = append(configs, NewNginxConfig())
	case PostgresqlServer:
		configs = append(configs, NewNetworkConfig(LowNetworkLevel))
		configs = append(configs, NewPostgresqlConfig())
	case ApacheServer:
		configs = append(configs, NewNetworkConfig(HighNetworkLevel))
		configs = append(configs, NewApacheConfig())
	}

	return
}
