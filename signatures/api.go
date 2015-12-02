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
