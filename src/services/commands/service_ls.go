package services

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func LsCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	if len(args) > 0 {
		fmt.Fprintln(errorWriter, "Ls does not accept arguments")
		return
	}

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Fprintf(errorWriter, "Error reading directory: %v\n", err)
		return
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	fmt.Fprintln(outputWriter, strings.Join(filenames, " "))
}
