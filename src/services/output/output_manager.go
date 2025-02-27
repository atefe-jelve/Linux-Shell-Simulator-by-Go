package services

import (
	"bytes"
	"io"
	"os"
)

var OutputWriter io.Writer = os.Stdout

var ErrorBuffer bytes.Buffer

func SetOutputFile(filePath string, appendMode bool) error {
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
	OutputWriter = file
	return nil
}

func SetErrorWriter() {
	ErrorBuffer.Reset()
}

func ResetOutput() {
	OutputWriter = os.Stdout
}

func ResetErrorWriter() {
	ErrorBuffer.Reset()
}

func WriteErrorToFile(filePath string) error {
	if ErrorBuffer.Len() > 0 {
		var file *os.File
		var err error
		file, err = os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.Write(ErrorBuffer.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}
