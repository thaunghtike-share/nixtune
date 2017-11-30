package signatures

type NodeJS struct{}

func (f *NodeJS) GetProfile() *Profile {
	p := &Profile{
		Name:         "nodejs",
		Subscription: StartupSubscription,
		Description:  "Configuration for NodeJS.",
		Deps:         []Profiler{&FastServer{}},
		References: []string{
			"https://engineering.gosquared.com/optimising-nginx-node-js-and-networking-for-heavy-workloads",
		},
	}

	return p
}
