package services

import "io"

func LogoutCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	username := ""
	if !IsSessionValid(username) {
		AddUserSession(username)
	}
	SetCurrentUser(username)
}
