/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package signatures

type SystemConfiger interface {
	GetEnv() map[string]string
	GetSysctl() map[string]string
	GetFiles() map[string]FileChange
}

type FileChange struct {
	Content string
	Append  bool
}

type Signatures struct {
	ServerConfigs map[string]SystemConfiger
}

func NewSignatures() *Signatures {
	s := &Signatures{}
	s.ServerConfigs = make(map[string]SystemConfiger)

	// Async Server configurations
	s.ServerConfigs["golang"] = NewGolangConfig()
	s.ServerConfigs["nodejs"] = NewNodejsConfig()
	s.ServerConfigs["nginx"] = NewNginxConfig()
	s.ServerConfigs["haproxy"] = NewHaproxyConfig()

	// Forking server configurations
	s.ServerConfigs["apache"] = &PostgresqlConfig{}
	s.ServerConfigs["postgresql"] = &PostgresqlConfig{}

	s.ServerConfigs["java"] = NewJavaConfig()

	return s
}

// SignatureTypes returns the slice of the different configurations
// that we have available.
func (s *Signatures) Types() (types []string) {
	for k, _ := range s.ServerConfigs {
		types = append(types, k)
	}

	return
}

func (c *Signatures) Get(sig string) SystemConfiger {
	return c.ServerConfigs[sig]
}
