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
		log.Println("FAIL:", err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() != 0 {
					err = errors.New(fmt.Sprintf("Exit Status: %d\n", status.ExitStatus()))
					fmt.Println(err)
					return err
				}
			}
		}
	}

	return
}

func writeFile(fileName, content string) {

}
