package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "agent":
		agent := NewKnightAgent()
		agent.ParseArgs(os.Args[2:])
		agent.Run()
	}
}
