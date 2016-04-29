package signatures

type Nginx struct{}

func (f *Nginx) GetProfile() *Profile {
	p := &Profile{
		Name:         "nginx",
		Subscription: StartupSubscription,
		Description:  "Configuration for Nginx.",
		Deps:         []Profiler{&FastServer{}},
	}

	p.Flags = make(map[string]*ProfileKV)
	p.Flags["nginx-conf"] = &ProfileKV{
		Description: "Location of the nginx.conf file.",
		Default:     "/etc/nginx/nginx.conf",
	}

	return p
}
