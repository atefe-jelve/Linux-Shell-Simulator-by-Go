package services

import (
	"fmt"
	"os"
)

func PwdCommand([]string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}
