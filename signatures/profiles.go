package signatures

import (
	"errors"
	"os"
)

type Profiles []Profiler

func (p Profiles) Get(s string, withDeps bool) (profile *Profile, err error) {
	if withDeps {
		profile = p.getWithDeps(s)
	} else {
		profile = p.getMergeDeps(s)
	}

	if profile == nil {
		return nil, errors.New("profile doesn't exist")
	}

	// Only return profile if they have the subscription.
	if profile.HasSubscription(currentSubscription) {
		return nil, errors.New("need to have subscription for this signature")
	}

	return profile, nil
}

func (p Profiles) getWithDeps(s string) (profile *Profile) {
	for _, k := range p {
		if p2 := k.GetProfile(); p2.Name == s {
			profile = p2
			break
		}
	}

	profile.parseValueTemplates()

	return
}

func (p Profiles) getMergeDeps(s string) *Profile {
	var (
		profile *Profile
	)

	for _, k := range p {
		if p2 := k.GetProfile(); p2.Name == s {
			profile = p2
			break
		}
	}

	if profile == nil {
		return nil
	}

	if profile.ProcFS == nil {
		profile.ProcFS = make(map[string]ProfileKV)
	}

	if profile.SysFS == nil {
		profile.SysFS = make(map[string]ProfileKV)
	}

	if profile.Env == nil {
		profile.Env = make(map[string]ProfileKV)
	}

	if profile.Vars == nil {
		profile.Vars = make(map[string]interface{})
	}

	for _, dep := range profile.Deps {
		depProfile := dep.GetProfile()
		if depProfile == nil {
			continue
		}

		if depProfile.Vars != nil {
			for k, v := range depProfile.Vars {
				_, ok := profile.Vars[k]
				if !ok {
					profile.Vars[k] = v
				}
			}
		}

		if depProfile.ProcFS != nil {
			for k, v := range depProfile.ProcFS {
				_, ok := profile.ProcFS[k]
				if !ok {
					profile.ProcFS[k] = v
				}
			}
		}

		if depProfile.SysFS != nil {
			for k, v := range depProfile.SysFS {
				_, ok := profile.SysFS[k]
				if !ok {
					profile.SysFS[k] = v
				}
			}
		}

		if depProfile.Env != nil {
			for k, v := range depProfile.Env {
				_, ok := profile.Env[k]
				if !ok {
					profile.Env[k] = v
				}
			}
		}
	}

	profile.Deps = []Profiler{}

	profile.parseValueTemplates()

	return profile
}

func Load() (p Profiles) {
	loadSubscription(os.Getenv("ACKSIN_API_KEY"))

	// Open Profiles
	p = append(p, &FastServer{})
	p = append(p, &FS{})
	p = append(p, &IO{})
	p = append(p, &Memory{})
	p = append(p, &Networking{})

	// Startup Profiles
	//p = append(p, &Apache{})
	p = append(p, &Golang{})
	//p = append(p, &HAProxy{})
	p = append(p, &Nginx{})
	p = append(p, &NodeJS{})
	//p = append(p, &PostgreSQL{})

	return
}
