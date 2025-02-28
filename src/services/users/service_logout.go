package services

import "io"

func LogoutCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	currentUser := GetCurrentUser()

	if currentUser == "" {
		return
	}

	cacheMutex.Lock()
	delete(sessionCache, currentUser)
	delete(sessionCache, "user_login")
	cacheMutex.Unlock()

	// fmt.Fprintln(outputWriter, "User logged out successfully.")
}
