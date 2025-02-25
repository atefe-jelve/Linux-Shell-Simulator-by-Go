package services

import (
	"fmt"
	"log"
	"os"
)

func LsCommand([]string) {
	files_directory, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files_directory {
		fmt.Print(file.Name(), "  ")
	}
	fmt.Println()
}
