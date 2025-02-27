
package services

import (
    "fmt"
    "io"
)

// ClearCommand clears the terminal screen by printing the ANSI escape sequence.
func ClearCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

    if len(args) > 0 {
		fmt.Fprintln(errorWriter, "ClearCommand does not accept any arguments")
		return
	}

	const clearScreen = "\033[H\033[2J"
	if _, err := fmt.Fprint(outputWriter, clearScreen); err != nil {

        fmt.Fprintf(errorWriter, "Error clearing screen: %v\n", err)
	}
}
