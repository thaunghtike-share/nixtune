package signatures

type IO struct{}

func (f *IO) GetProfile() *Profile {
	p := &Profile{
		Name:         "io",
		Subscription: OpenSubscription,
		Description:  "Settings for IO optimizations",
		References: []string{
			"http://www.brendangregg.com/linuxperf.html",
		},
	}

	p.SysFS = make(map[string]ProfileKV)
	p.SysFS["/sys/block/*/queue/rq_afinity"] = ProfileKV{
		Value: "2",
	}

	p.SysFS["/sys/block/*/queue/scheduler"] = ProfileKV{
		Value: "noop",
	}

	p.SysFS["/sys/block/*/queue/read_ahead_kb"] = ProfileKV{
		Value: "256",
	}

	return p
}
