package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	commandsservices "projectshell/src/services/commands"
	historyservices "projectshell/src/services/history"
	outputservices "projectshell/src/services/output"
	userservices "projectshell/src/services/users"
	"projectshell/src/utils"
	"strings"
)

var commands = map[string]func([]string, io.Writer, io.Writer){
	"exit":  commandsservices.ExitCommand,
	"echo":  commandsservices.EchoCommand,
	"cat":   commandsservices.CatCommand,
	"type":  commandsservices.TypeCommand,
	"pwd":   commandsservices.PwdCommand,
	"cd":    commandsservices.CdCommand,
	"ls":    commandsservices.LsCommand,
	"clear": commandsservices.ClearCommand,

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
			commandsservices.EchoCommand([]string{input[5:]}, outputservices.OutputWriter, &outputservices.ErrorBuffer)
			historyservices.LogHistory([]string{input})
		}

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		//
		info := outputservices.CheckRedirection(input)
		if info.Redirection {
			if info.ErrorRedirection {
				outputservices.SetErrorWriter()
			} else {
				if err := outputservices.SetOutputFile(info.FilePath, info.AppendMode); err != nil {
					fmt.Println(err)
					continue
				}
			}
		}

		if command, ok := commands[cmd]; ok {
			historyservices.LogHistory(args)
			//
			command(args[1:], outputservices.OutputWriter, &outputservices.ErrorBuffer)
			if cmd == "exit" {
				utils.CleanHistory(utils.GetUserId("anonymous"))
				break
			}
		} else {
			if !commandsservices.IsBuiltin(cmd) {
				commandsservices.ExecuteCommand(cmd, args[1:], outputservices.OutputWriter, &outputservices.ErrorBuffer)
			}
		}

		//
		if info.Redirection {
			if info.ErrorRedirection {
				if err := outputservices.WriteErrorToFile(info.FilePath); err != nil {
					fmt.Println(err)
				}
				outputservices.ResetErrorWriter()
			} else {
				outputservices.ResetOutput()

			}

		}

	}
	return commandsservices.ExitStatus

}
