package output

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// var outputWriter io.Writer = os.Stdout
// var errorWriter io.Writer = os.Stderr

func setOutputFile(filePath string, appendMode bool) error {
	var file *os.File
	var err error
	if appendMode {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		file, err = os.Create(filePath)
	}
	if err != nil {
		return err
	}
	outputWriter = file
	return nil
}

func setErrorFile(filePath string, appendMode bool) error {
	var file *os.File
	var err error
	if appendMode {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		file, err = os.Create(filePath)
	}
	if err != nil {
		return err
	}
	errorWriter = file
	return nil
}

// func resetOutput() {
// 	outputWriter = os.Stdout
// 	errorWriter = os.Stderr
// }

func ManageOutout() {

	// Check for output redirection
	if strings.Contains(input, ">>") {
		redirection = true
		parts := strings.Split(input, ">>")
		args = strings.Fields(parts[0])
		filePath = strings.TrimSpace(parts[1])
		appendMode = true
	} else if strings.Contains(input, ">") {
		redirection = true
		parts := strings.Split(input, ">")
		args = strings.Fields(parts[0])
		filePath = strings.TrimSpace(parts[1])
	} else if strings.Contains(input, "2>>") {
		redirection = true
		errorRedirection = true
		parts := strings.Split(input, "2>>")
		args = strings.Fields(parts[0])
		filePath = strings.TrimSpace(parts[1])
		appendMode = true
	} else if strings.Contains(input, "2>") {
		redirection = true
		errorRedirection = true
		parts := strings.Split(input, "2>")
		args = strings.Fields(parts[0])
		filePath = strings.TrimSpace(parts[1])
	}

	// Set output or error redirection if needed
	if redirection {
		if errorRedirection {
			if err := setErrorFile(filePath, appendMode); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			if err := setOutputFile(filePath, appendMode); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

}
