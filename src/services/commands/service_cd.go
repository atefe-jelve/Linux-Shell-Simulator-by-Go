package services

import (
	"fmt"
	"io"
	"os"
)

func CdCommand(path []string, outputWriter io.Writer, errorWriter io.Writer) {
	if len(path) < 1 {
		fmt.Println()
		return
	}
	err := os.Chdir(path[0])
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}
	fmt.Printf("Changed directory to: %s\n", path[0])
}
