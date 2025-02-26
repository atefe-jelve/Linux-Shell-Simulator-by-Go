package main

import (
	"fmt"
	// session_user "projectshell/src/services/users"

)

func main() {
	// session_user.AddUserCommand([]string{"anonymous"})

	exitStatus := ReceiveCommand()
	if exitStatus != 99999 {
		fmt.Printf("exit status %d\n", exitStatus)
	}
}
