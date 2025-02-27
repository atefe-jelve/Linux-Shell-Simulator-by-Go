package services

import (
	"fmt"
	"io"
	"log"
	"os"
)

// var outputWriter io.Writer = os.Stdout

func LsCommand(args []string, writer io.Writer) {
	files_directory, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files_directory {
		fmt.Fprint(writer, file.Name(), " ")
		// fmt.Print(file.Name(), "  ")
	}
	// fmt.Println()
	fmt.Fprintln(writer)
}
