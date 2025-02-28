
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
		prompt := "$ "
		if currentUser != "" {
			prompt = fmt.Sprintf("%s :$ ", currentUser)
		}
		fmt.Print(prompt)

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		// Handle echo command separately
		if strings.HasPrefix(input, "echo") {
			handleEcho(input)
			continue
		}

		// Split input into command and arguments
		args = strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cmd := args[0]
		info := outputservices.CheckRedirection(input)

		if info.Redirection {
			handleRedirection(info)
		}

		if command, ok := commands[cmd]; ok {
			historyservices.LogHistory(args)

			command(args[1:], os.Stdout, os.Stderr)
			if cmd == "exit" {
				utils.CleanHistory(utils.GetUserId("anonymous"))
				break
			}
		} else {
			if !commandsservices.IsBuiltin(cmd) {
				commandsservices.ExecuteCommand(cmd, args[1:], outputservices.OutputWriter, &outputservices.ErrorBuffer)

			}
		}

		// Reset redirection state
		if info.Redirection {
			resetRedirection(info)
		}

	}
	return commandsservices.ExitStatus

}

// Handle the echo command
func handleEcho(input string) {
	if len(input) > 5 {
		commandsservices.EchoCommand([]string{input[5:]}, outputservices.OutputWriter, &outputservices.ErrorBuffer)
		historyservices.LogHistory([]string{input})
	}else{
		fmt.Println()
	}
}

// Handle redirection
func handleRedirection(info outputservices.RedirectionInfo) {
	if info.ErrorRedirection {
		outputservices.SetErrorWriter()
	} else {
		if err := outputservices.SetOutputFile(info.FilePath, info.AppendMode); err != nil {
			fmt.Println(err)
		}
	}
}

// Reset redirection (restore
func resetRedirection(info outputservices.RedirectionInfo) {
	if info.ErrorRedirection {
		if err := outputservices.WriteErrorToFile(info.FilePath); err != nil {
			fmt.Println(err)
		}
		outputservices.ResetErrorWriter()
	} else {
		outputservices.ResetOutput()
	}
}
