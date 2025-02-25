package services

import (
	"fmt"
	"strconv"
)

var ExitStatus int

func ExitCommand(args []string) {
	if len(args) == 0 {
		ExitStatus = 0
	} else if len(args) == 1 {
		status, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid status code")
			ExitStatus = 99999
		} else {
			ExitStatus = status
		}
	} else {
		fmt.Println("too many arguments")
		ExitStatus = 99999
	}
}
