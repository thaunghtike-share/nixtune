package signatures

type HAProxy struct{}

func (f *HAProxy) GetProfile() *Profile {
	p := &Profile{
		Name:         "haproxy",
		Subscription: StartupSubscription,
		Description:  "Configuration for HAProxy",
		Deps:         []Profiler{&FastServer{}},
	}

	return p
}
