package services

import (
	"fmt"
	"os"
)

func CdCommand(path []string) {
	err := os.Chdir(path[0])
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}
	fmt.Printf("Changed directory to: %s\n", path)
}
