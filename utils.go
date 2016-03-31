/* Acksin Autotune - Kernel Autotuning
 * Copyright (C) 2015 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func runCmd(cmdName string, cmdArgs ...string) (err error) {
	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println(cmdArgs)

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() != 0 {
					return fmt.Errorf("Exit Status: %d\n", status.ExitStatus())
				}
			}
		}
	}

	return
}

func runCmdGetOutput(cmdName string, cmdArgs ...string) []byte {
	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return nil
	}

	return cmdOut
}

func procfsGet(k string) string {
	return string(runCmdGetOutput("sysctl", "-b", k))
}

func procfsSet(k, v string) string {
	return string(runCmdGetOutput("sysctl", "-w", fmt.Sprintf("%s=\"%s\"", k, v)))
}

func sysfsGet(k string) string {
	return string(runCmdGetOutput("cat", k))
}

func sysfsSet(k, v string) string {
	return string(runCmdGetOutput("echo", v, ">", k))
}
