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
package signatures

type ServerType int

const (
	// Async and High Network Throughput
	GolangServer ServerType = iota
	NodejsServer
	NginxServer

	// Forking Servers
	ApacheServer
	PostgreSQLServer

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
	var (
		networkLevel NetworkLevel
	)

	switch ServerSignature(signature) {
	case GolangServer:
		networkLevel = HighNetworkLevel
		configs = append(configs, NewGolangConfig())
		configs = append(configs, NewNetworkConfig(networkLevel))
	case NodejsServer:
		networkLevel = HighNetworkLevel
		configs = append(configs, NewNetworkConfig(networkLevel))
	//	configs = append(configs, NewGolangConfig())
	case NginxServer:
		networkLevel = HighNetworkLevel
		//configs = append(configs, NewGolangConfig())
		configs = append(configs, NewNetworkConfig(networkLevel))
	case ApacheServer:
		configs = append(configs, NewPostgresqlConfig())
		configs = append(configs, NewNetworkConfig(networkLevel))
	}

	return
}
