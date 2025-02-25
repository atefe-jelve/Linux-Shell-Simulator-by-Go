package services

import (
	"sync"
)

var sessionCache = make(map[string]string)
var cacheMutex = sync.Mutex{}

// AddUserSession adds a user session to the cache.
func AddUserSession(username string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	sessionCache[username] = "active"
}

// IsSessionValid checks if a user session is valid.
func IsSessionValid(username string) bool {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	_, exists := sessionCache[username]
	return exists
}

// GetCurrentUser returns the current logged-in user.
func GetCurrentUser() string {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	return sessionCache["user_login"]
}

// SetCurrentUser sets the current logged-in user.
func SetCurrentUser(username string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	sessionCache["user_login"] = username
}
