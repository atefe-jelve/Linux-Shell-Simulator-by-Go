package main

import (
	"fmt"
	outputservices "projectshell/src/services/output"
	session_user "projectshell/src/services/users"
)

func main() {
	session_user.AddUserCommand([]string{"anonymous"}, outputservices.OutputWriter, &outputservices.ErrorBuffer)

	exitStatus := ReceiveCommand()
	if exitStatus != 99999 {
		fmt.Printf("exit status %d\n", exitStatus)
	}
}
