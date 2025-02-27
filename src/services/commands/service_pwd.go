package services

import (
	"fmt"
	"io"
	"os"
)

// PwdCommand prints the current directory to outputWriter.
func PwdCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	if len(args) > 0 {
		fmt.Fprintln(errorWriter, "pwd does not accept arguments")
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(errorWriter, "error getting current directory:", err)
		return
	}

	fmt.Fprintln(outputWriter, dir)
}
