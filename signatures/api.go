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
