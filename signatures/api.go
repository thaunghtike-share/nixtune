package signatures

type ServerType int

const (
	GolangServer ServerType = iota
	NodejsServer
	ApacheServer
	NginxServer
	JavaServer
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
