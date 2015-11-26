// TODO: Need to be more intelligent. Update based on the system
// type. Defaulting to Ubuntu Linux for now.

package main

import (
	"fmt"
)

type FileChange struct {
	Content string
	Append  bool
}

type SystemConfig struct {
	Sysctl map[string]string
	Env    map[string]string
	Files  map[string]FileChange
}

const (
	EnvFileName = "/etc/profile.d/99_anatma_knight.sh"
)

func (sc *SystemConfig) updateEnv() {
	var (
		fileContent string
	)

	for k, v := range sc.Env {
		fileContent += fmt.Sprintf("%s=%s\n", k, v)

	}

	writeFile(EnvFileName, fileContent)
}

func (sc *SystemConfig) updateSysctl() {
	for k, v := range sc.Sysctl {
		runCmd("sysctl", "-a", fmt.Sprintf("%s='%v'", k, v))
	}
}

// TODO
func (sc *SystemConfig) updateFiles() {
	// for k, v := range sc.Files {

	// }
}

func (sc *SystemConfig) Update() {
	sc.updateEnv()
	sc.updateSysctl()
	sc.updateFiles()
}
