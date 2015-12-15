/*
 * Anatma Autotune - Kernel Autotuning
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func runCmd(cmdName string, cmdArgs ...string) (err error) {
	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() != 0 {
					return errors.New(fmt.Sprintf("Exit Status: %d\n", status.ExitStatus()))
				}
			}
		}
	}

	return
}

func writeFile(fileName, content string) {

}

func logMe(logType string, logString string) {
	log.Println(logType, logString)
}
