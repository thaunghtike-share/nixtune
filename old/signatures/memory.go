package signatures

type Memory struct{}

func (f *Memory) GetProfile() *Profile {
	p := &Profile{
		Name:         "memory",
		Subscription: StartupSubscription,
		Description:  "Settings for memory optimizations",
	}

	p.ProcFS = f.procfs()
	p.SysFS = f.sysfs()

	return p
}

func (f *Memory) procfs() map[string]*ProfileKV {
	p := make(map[string]*ProfileKV)
	p["vm.swappiness"] = &ProfileKV{
		Value:       "0",
		Description: " Disable swapping and clear the file system page cache to free memory first.",
	}

	p["proc.min_free_kbytes"] = &ProfileKV{
		Value:       "65536",
		Description: "Amount of memory to keep free. Don't want to make this too high as Linux will spend more time trying to reclaim memory.",
	}

	return p
}

func (f *Memory) sysfs() map[string]*ProfileKV {
	p := make(map[string]*ProfileKV)

	p["/sys/kernel/mm/transparent_hugepage/enabled"] = &ProfileKV{
		Value:       "always",
		Description: "Explit huge page usage making the page size of 2 or 4 MB instead of 4kb. Should reduce CPU overhead and improve MMU page translation.",
	}

	return p
}
