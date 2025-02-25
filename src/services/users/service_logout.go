package services

func LogoutCommand([]string) {
	
	username := ""
	if !IsSessionValid(username) {
		AddUserSession(username)
	}
	SetCurrentUser(username)
}
