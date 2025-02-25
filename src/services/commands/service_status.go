package services

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(command + ": command not found")

	}
}
