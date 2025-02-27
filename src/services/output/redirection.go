package services

import "strings"

type RedirectionInfo struct {
	Redirection      bool
	ErrorRedirection bool
	FilePath         string
	AppendMode       bool
	Args             []string
}

func CheckRedirection(input string) RedirectionInfo {
	var info RedirectionInfo
	if strings.Contains(input, ">>") {
		info.Redirection = true
		parts := strings.Split(input, ">>")
		info.Args = strings.Fields(parts[0])
		info.FilePath = strings.TrimSpace(parts[1])
		info.AppendMode = true
	} else if strings.Contains(input, ">") {
		info.Redirection = true
		parts := strings.Split(input, ">")
		info.Args = strings.Fields(parts[0])
		info.FilePath = strings.TrimSpace(parts[1])
	} else if strings.Contains(input, "2>>") {
		info.Redirection = true
		info.ErrorRedirection = true
		parts := strings.Split(input, "2>>")
		info.Args = strings.Fields(parts[0])
		info.FilePath = strings.TrimSpace(parts[1])
		info.AppendMode = true
	} else if strings.Contains(input, "2>") {
		info.Redirection = true
		info.ErrorRedirection = true
		parts := strings.Split(input, "2>")
		info.Args = strings.Fields(parts[0])
		info.FilePath = strings.TrimSpace(parts[1])
	}
	return info
}
