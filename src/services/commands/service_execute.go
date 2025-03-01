package services

import (
	"fmt"
	"io"
	"os/exec"
)

// this func define unlike others func and evaluate deferent for execute commands
func ExecuteCommand(command string, args []string, outputWriter io.Writer, errorWriter io.Writer) {

	cmd := exec.Command(command, args...)
	cmd.Stdout = outputWriter
	cmd.Stderr = errorWriter

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(errorWriter, "%s: command not found\n", cmd)
	}
}
