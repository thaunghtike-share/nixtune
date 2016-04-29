package signatures

type FS struct{}

func (f *FS) files() (p map[string]map[string]ProfileKV) {
	p = make(map[string]map[string]ProfileKV)
	// files:
	//   "/etc/fstab":
	//     "noattime":
	//       value: '{{ eachLine | contains "noatime" }}'
	//     "discard":
	//       value: '{{ eachLine | notcontains "discard" }}'
	//       description: |
	//         Avoid having a discard mount attribute as every time a file is
	//         deleted the SSD will also do a TRIM for future writing. This
	//         will increase time it takes to delete a file.

	//         Better option is to run a daily/weekly cron.
	//       if:
	//         - ssd
	//   "/etc/security/limits.d/*":
	//     "ulimits":
	//       value: |
	//         List the ulimit values.

	return
}

func (f *FS) cron() (p map[string]ProfileKV) {
	p = make(map[string]ProfileKV)
	// cron:
	//   "fs-trim":
	//     value: |
	//       #!/bin/sh
	//       #
	//       # To find which FS support trim, we check that DISC-MAX (discard max bytes)
	//       # is great than zero. Check discard_max_bytes documentation at
	//       # https://www.kernel.org/doc/Documentation/block/queue-sysfs.txt
	//       #
	//       for fs in $(lsblk -o MOUNTPOINT,DISC-MAX,FSTYPE | grep -E '^/.* [1-9]+.* ' | awk '{print $1}'); do
	//         fstrim "$fs"
	//       done
	//     schedule: weekly

	return
}

func (f *FS) procfs() (p map[string]ProfileKV) {
	p = make(map[string]ProfileKV)

	p["vm.dirty_ratio"] = ProfileKV{
		Value:       "80",
		Description: "",
	}

	p["vm.dirty_background_ratio"] = ProfileKV{
		Value:       "5",
		Description: "",
	}
	p["vm.dirty_expire_centisecs"] = ProfileKV{
		Value:       "1200",
		Description: "",
	}
	p["proc.file-max"] = ProfileKV{
		Value:       "2097152",
		Description: "The max amount of file handlers that the Linux kernel will allocate. This is one part the other part is setting the ulimits.",
	}

	p["fs.file-max"] = ProfileKV{
		Value:       "{{ multiply .Stats.System.Memory.Physical.Total 11 }}",
		Description: "",
	}

	return p
}

func (f *FS) GetProfile() *Profile {
	p := &Profile{
		Name:         "fs",
		Subscription: OpenSubscription,
		Description:  "Settings for fs optimizations",
		References: []string{
			"https://tweaked.io/guide/kernel/",
			"http://blog.neutrino.es/2013/howto-properly-activate-trim-for-your-ssd-on-linux-fstrim-lvm-and-dmcrypt/",
		},
		ProcFS: f.procfs(),
		Files:  f.files(),
		Cron:   f.cron(),
	}

	return p
}
