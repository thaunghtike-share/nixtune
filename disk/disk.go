/*
 * Copyright (C) 2016 opszero <hey@opszero.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

// Package disk shows information about block drives in the machine.
package disk

import (
	"encoding/json"
	"log"

	"os/exec"

	"github.com/acksin/acksin/stats/mvp"
	"github.com/opszero/go-fstab"
)

// Disk returns information about the block drives on the machine.

type Disk struct {
	BlockDevices []mvp.BlockDevice
	FStab        fstab.Mounts
}

// New returns a Disk object representing system disk information.
func New() *Disk {

	m, err := fstab.ParseSystem()
	if err != nil {
		log.Println(err)
	}

	return &Disk{
		BlockDevices: mvp.GetLsBlk(),
		FStab:        m,
	}
}

type BlockDevice struct {
	Name       string        `json:"name"`
	MajMin     string        `json:"maj:min"`
	RM         string        `json:"rm"`
	Size       string        `json:"size"`
	RO         string        `json:"ro"`
	Type       string        `json:"type"`
	MountPoint string        `json:"mountpoint"`
	Children   []BlockDevice `json:"children,omitempty"`
}

// GetLsBlk gets the output of lsblk and returns that as an object.
func GetLsBlk() []BlockDevice {
	out, err := exec.Command("lsblk", "-J", "-a").Output()
	if err != nil {
		log.Println(err)
	}

	var lsblk struct {
		BlockDevices []BlockDevice `json:"blockdevices"`
	}

	json.Unmarshal(out, &lsblk)

	return lsblk.BlockDevices
}
