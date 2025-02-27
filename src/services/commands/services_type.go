package services

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var built_in_commands = []string{"exit", "echo", "cat", "type", "ls", "pwd", "cd", "clear"}

func TypeCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	if len(args) == 0 {
		fmt.Fprintln(outputWriter, "Please insert a command to check")
		return
	}

	cmd := args[0]
	if IsBuiltin(cmd) {
		fmt.Fprintf(outputWriter, "%s is a shell builtin\n", cmd)
	} else if fullPath, found := findExecutable(cmd); found {
		fmt.Fprintf(outputWriter, "%s is %s\n", cmd, fullPath)
	} else {
		fmt.Fprintf(outputWriter, "%s: command not found\n", cmd)
	}
}

func IsBuiltin(cmd string) bool {
	for _, builtin := range built_in_commands {
		if cmd == builtin {
			return true
		}
	}
	return false
}

func findExecutable(cmd string) (string, bool) {
	path := os.Getenv("PATH")
	dirs := strings.Split(path, string(os.PathListSeparator))
	for _, dir := range dirs {
		fullPath := filepath.Join(dir, cmd)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
}
