package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	commandsservices "projectshell/src/services/commands"
	historyservices "projectshell/src/services/history"
	manageoutout "projectshell/src/services/output_manager"
	userservices "projectshell/src/services/users"
	"projectshell/src/utils"
	"strings"
)

var commands = map[string]func([]string, io.Writer){
	// "exit": commandsservices.ExitCommand,
	// "echo": commandsservices.EchoCommand,
	// "cat":  commandsservices.CatCommand,
	// "type": commandsservices.TypeCommand,
	// "pwd":  commandsservices.PwdCommand,
	// "cd":   commandsservices.CdCommand,
	"ls": commandsservices.LsCommand,

	// 	"login":   userservices.LoginCommand,
	// 	"logout":  userservices.LogoutCommand,
	// 	"adduser": userservices.AddUserCommand,

	// "history": historyservices.HistoryCommand,
}

var outputWriter io.Writer = os.Stdout
var errorWriter io.Writer = os.Stderr

// func setOutputFile(filePath string, appendMode bool) error {
//     var file *os.File
//     var err error
//     if appendMode {
//         file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     } else {
//         file, err = os.Create(filePath)
//     }
//     if err != nil {
//         return err
//     }
//     outputWriter = file
//     return nil
// }

// func setErrorFile(filePath string, appendMode bool) error {
//     var file *os.File
//     var err error
//     if appendMode {
//         file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     } else {
//         file, err = os.Create(filePath)
//     }
//     if err != nil {
//         return err
//     }
//     errorWriter = file
//     return nil
// }

func resetOutput() {
	outputWriter = os.Stdout
	errorWriter = os.Stderr
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
		//
		redirection := false
		input := scanner.Text()
		filePath := ""
		appendMode := false
		errorRedirection := false

		if !strings.HasPrefix(input, "echo") {
			args = strings.Fields(input)

		} else {
			commandsservices.EchoCommand([]string{input[5:]})
			historyservices.LogHistory([]string{input})
		}

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		manageoutout.ManageOutout(filePath, appendMode, errorRedirection)
		// // Check for output redirection
		// if strings.Contains(input, ">>") {
		//     redirection = true
		//     parts := strings.Split(input, ">>")
		//     args = strings.Fields(parts[0])
		//     filePath = strings.TrimSpace(parts[1])
		//     appendMode = true
		// } else if strings.Contains(input, ">") {
		//     redirection = true
		//     parts := strings.Split(input, ">")
		//     args = strings.Fields(parts[0])
		//     filePath = strings.TrimSpace(parts[1])
		// } else if strings.Contains(input, "2>>") {
		//     redirection = true
		//     errorRedirection = true
		//     parts := strings.Split(input, "2>>")
		//     args = strings.Fields(parts[0])
		//     filePath = strings.TrimSpace(parts[1])
		//     appendMode = true
		// } else if strings.Contains(input, "2>") {
		//     redirection = true
		//     errorRedirection = true
		//     parts := strings.Split(input, "2>")
		//     args = strings.Fields(parts[0])
		//     filePath = strings.TrimSpace(parts[1])
		// }

		// // Set output or error redirection if needed
		// if redirection {
		//     if errorRedirection {
		//         if err := setErrorFile(filePath, appendMode); err != nil {
		//             fmt.Println(err)
		//             continue
		//         }
		//     } else {
		//         if err := setOutputFile(filePath, appendMode); err != nil {
		//             fmt.Println(err)
		//             continue
		//         }
		//     }
		// }

		if command, ok := commands[cmd]; ok {
			historyservices.LogHistory(args)
			//
			command(args[1:], outputWriter)
			if cmd == "exit" {
				utils.CleanHistory(utils.GetUserId("anonymous"))
				break
			}
		} else {
			if !commandsservices.IsBuiltin(cmd) {
				commandsservices.ExecuteCommand(cmd, args[1:])
			}
		}
		//
		if redirection {
			resetOutput()
		}

	}
	return commandsservices.ExitStatus

}
