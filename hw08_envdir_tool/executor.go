package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var exitError *exec.ExitError
	for envName, envValue := range env {
		if envValue.NeedRemove {
			os.Unsetenv(envName)
			continue
		}
		os.Setenv(envName, envValue.Value)
	}
	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		if errors.As(err, &exitError) {
			returnCode = exitError.ExitCode()
		}
	}
	return returnCode
}
