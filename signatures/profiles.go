package signatures

import (
	"os"
)

type Profiles []Profiler

func (p Profiles) Get(s string, withDeps bool) (profile *Profile) {
	if withDeps {
		profile = p.getWithDeps(s)
	} else {
		profile = p.getMergeDeps(s)
	}

	// Only return profile if they have the subscription.
	if profile.HasSubscription(currentSubscription) {
		return profile
	}

	return nil
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

	for _, dep := range profile.Deps {
		depProfile := dep.GetProfile()
		if depProfile == nil {
			continue
		}

		for k, v := range depProfile.ProcFS {
			_, ok := profile.ProcFS[k]
			if !ok {
				profile.ProcFS[k] = v
			}
		}

		for k, v := range depProfile.SysFS {
			_, ok := profile.SysFS[k]
			if !ok {
				profile.SysFS[k] = v
			}
		}

		for k, v := range depProfile.Env {
			_, ok := profile.Env[k]
			if !ok {
				profile.Env[k] = v
			}
		}
	}

	profile.Deps = []Profiler{}

	profile.parseValueTemplates()

	return profile
}

func (p Profiles) Add(profile Profiler) {
	p = append(p, profile)
}

func Load() Profiles {
	var p Profiles

	loadSubscription(os.Getenv("AUTOTUNE_API_KEY"))

	p.Add(&FastServer{})
	p.Add(&FS{})
	p.Add(&IO{})
	p.Add(&Memory{})
	p.Add(&Networking{})

	return p
}
