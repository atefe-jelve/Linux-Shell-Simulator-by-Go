package services

import (
	"fmt"
	"io"
	"os"
)

func CatCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	if len(args) == 0 {
		fmt.Fprintln(errorWriter, "No files provided")
		return
	}

	for _, fileName := range args {
		file, err := os.Open(string(fileName))
		if err != nil {
			fmt.Fprintf(errorWriter, "Error opening file %s: %v\n", fileName, err)
			continue
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(errorWriter, "Error opening file %s: %v\n", fileName, err)
			continue
		}

		fmt.Fprintln(outputWriter, string(content))

		defer file.Close()

	}
}
