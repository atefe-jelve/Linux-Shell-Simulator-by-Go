package services

import (
	"fmt"
	"io"
	"os"
)

func PwdCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}
