package main

type KnightAgent struct{}

func NewKnightAgent() *KnightAgent {

	return &KnightAgent{}
}

func (k *KnightAgent) ParseArgs(args []string) {
	// profile-hints
	// service-name

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
