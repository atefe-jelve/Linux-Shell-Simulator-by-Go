package services

import (
	"fmt"
	"io"
	"os"
)

func LsCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Fprintln(errorWriter, err)
		return
	}
	for _, file := range files {
		fmt.Fprint(outputWriter, file.Name(), " ")
	}
	fmt.Fprintln(outputWriter)
}
