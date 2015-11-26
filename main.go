package main

import (
	"os"
)

type ServerType int

const (
	GolangServer ServerType = iota
	NodejsServer
	ApacheServer
	NginxServer
	JavaServer
)

type KnightAgent struct{}

func NewKnightAgent() *KnightAgent {

	return &KnightAgent{}
}

func (k *KnightAgent) ParseArgs(args []string) {
	// profile-hints

}

func (k *KnightAgent) Run() {
	for {
		var (
			sc *SystemConfig
		)

		switch getServerType() {
		case GolangServer:
			// TODO: Need to save the server state. Need
			// to be able to update server state if config
			// is bad.
			sc = golangConfig(k)
		}

		sc.Update()
	}
}

func getServerType() ServerType {
	// guessServerType()

	return GolangServer
}

func main() {
	switch os.Args[1] {
	case "agent":
		agent := NewKnightAgent()
		agent.ParseArgs(os.Args[2:])
		agent.Run()
	}
}
