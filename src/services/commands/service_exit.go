package services

import (
	"fmt"
	"io"
	"strconv"
)

const (
	StatusOK    = 0
	StatusError = 99999
)

var ExitStatus int

func ExitCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	switch len(args) {
	case 0:
		ExitStatus = StatusOK
	case 1:
		status, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(errorWriter, "Invalid status code")
			ExitStatus = StatusError
		} else {
			ExitStatus = status
		}
	default:
		fmt.Fprintln(errorWriter, "Too many arguments")
		ExitStatus = StatusError
	}
}
