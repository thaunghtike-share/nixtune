package main

type ServerType int

const (
	GolangServer ServerType = iota
	NodejsServer
	ApacheServer
	NginxServer
	JavaServer
)

func getServerType() ServerType {
	// guessServerType()

	return GolangServer
}
