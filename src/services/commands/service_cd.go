package services

import (
	"fmt"
	"io"
	"os"
)

func CdCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	if len(args) == 0 {
		fmt.Fprintln(errorWriter, "No directory provided")
		return
	}

	if len(args) > 1 {
		fmt.Fprintln(errorWriter, "Only one directory can be provided")
		return
	}

	err := os.Chdir(args[0])
	if err != nil {
		fmt.Fprintf(errorWriter, "Error changing directory to %s: %v\n", args[0], err)
		return
	}

	fmt.Fprintf(outputWriter, "Changed directory to: %s\n", args[0])

}
