package signatures

type FastServer struct{}

func (fs *FastServer) GetProfile() *Profile {
	return &Profile{
		Name:         "fast-server",
		Description:  "Settings for a general fast-server machine.",
		Subscription: StartupSubscription,
		Deps: []Profiler{
			&FS{},
			&IO{},
			&Memory{},
			&Networking{},
		},
	}

}
