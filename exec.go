package gosugar

import (
	"os/exec"
	"github.com/weakish/goaround"
	"syscall"
	"os"
)

// Exec completely replaces the Go current process with the specified command.
func Exec(cmdSlice []string) {
	var command string = cmdSlice[0]
	var commandPath string
	commandPath, err := exec.LookPath(command)
	goaround.FatalIf(err)

	err = syscall.Exec(commandPath, cmdSlice, os.Environ())
	goaround.FatalIf(err)
}
