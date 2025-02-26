package main

import (
	"bufio"
	"fmt"
	"os"
	commandsservices "projectshell/src/services/commands"
	historyservices "projectshell/src/services/history"
	userservices "projectshell/src/services/users"
	"strings"
)

var commands = map[string]func([]string){
	"exit": commandsservices.ExitCommand,
	"echo": commandsservices.EchoCommand,
	"cat":  commandsservices.CatCommand,
	"type": commandsservices.TypeCommand,
	"pwd":  commandsservices.PwdCommand,
	"cd":   commandsservices.CdCommand,
	"ls":   commandsservices.LsCommand,

	"login":   userservices.LoginCommand,
	"logout":  userservices.LogoutCommand,
	"adduser": userservices.AddUserCommand,

	"history": historyservices.HistoryCommand,
}

func ReceiveCommand() int {
	var args []string
	scanner := bufio.NewScanner(os.Stdin)

	for {

		currentUser := userservices.GetCurrentUser()
		if currentUser == "" {
			fmt.Print("$ ")
		} else {
			fmt.Printf("%s :$ ", currentUser)
		}

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if !strings.HasPrefix(input, "echo") {
			args = strings.Fields(input)

		} else {
			commandsservices.EchoCommand([]string{input[5:]})
			historyservices.LogHistory(args)

		}

		if len(args) == 0 {
			continue
		}

		cmd := args[0]
		if command, ok := commands[cmd]; ok {
			historyservices.LogHistory(args)

			command(args[1:])
			if cmd == "exit" {
				break
			}
		} else {
			if !commandsservices.IsBuiltin(cmd) {
				commandsservices.ExecuteCommand(cmd, args[1:])
			}
		}
	}
	return commandsservices.ExitStatus

}
