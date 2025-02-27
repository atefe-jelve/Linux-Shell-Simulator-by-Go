package services

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func ExecuteCommand(command string, args []string, outputWriter io.Writer, errorWriter io.Writer) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(command + ": command not found")

	}
}
