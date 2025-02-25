package services

import (
	"fmt"
	"io"
	"os"
)



func CatCommand(args []string) {
	n := len(args)

	for i := 0; i < n; i++ {
		file, err := os.Open(string(args[i]))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(string(content))
	}
}
