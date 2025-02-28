package services

// import (
// 	"fmt"
// 	"io"
// 	"projectshell/src/databases"
// 	command_model "projectshell/src/services/commands"
// 	session_user "projectshell/src/services/users"
// 	"projectshell/src/utils"
// 	"sort"
// 	"strings"
// 	"time"
// )

// type CommandInfo struct {
// 	Count int
// 	Time  time.Time
// }

// type kv struct {
// 	Key   string
// 	Value CommandInfo
// }

// const CleanCommand = "clean"

// func HistoryCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

// 	db := databases.GetDB()
// 	var id uint

// 	currentUser := session_user.GetCurrentUser()
// 	if currentUser == "" {
// 		id = utils.GetUserId("anonymous")
// 	} else {
// 		id = utils.GetUserId(currentUser)
// 	}

// 	if len(args) > 0 && args[0] == CleanCommand {
// 		utils.CleanHistory(id)
// 		return
// 	}

// 	var commandObjs []command_model.Command
// 	if err := db.Where("created_by = ?", id).Find(&commandObjs).Error; err != nil {
// 		fmt.Printf("Error retrieving commands: %v\n", err)
// 		return
// 	}

// 	if len(commandObjs) == 0 {
// 		fmt.Println("empty command history")
// 		return
// 	}

// 	commandsMap := make(map[string]CommandInfo)
// 	for _, commandObj := range commandObjs {
// 		if val, exists := commandsMap[commandObj.Text]; exists {
// 			val.Count++
// 			val.Time = commandObj.CreatedAt
// 			commandsMap[commandObj.Text] = val
// 		} else {
// 			commandsMap[commandObj.Text] = CommandInfo{
// 				Count: 1,
// 				Time:  commandObj.CreatedAt,
// 			}
// 		}
// 	}

// 	sortedCommands := sortedHistory(commandsMap)

// 	fmt.Printf("|%-20s|%-10s|\n", "Command", "Count")
// 	fmt.Println(strings.Repeat("-", 32))

// 	for _, kv := range sortedCommands {
// 		fmt.Printf("|%-20s|%-10d|\n", kv.Key, kv.Value.Count)
// 	}
// }

// func sortedHistory(commandsMap map[string]CommandInfo) []kv {

// 	var sortedCommands []kv
// 	for k, v := range commandsMap {
// 		sortedCommands = append(sortedCommands, kv{k, v})
// 	}

// 	sort.Slice(sortedCommands, func(i, j int) bool {
// 		if sortedCommands[i].Value.Count == sortedCommands[j].Value.Count {
// 			return sortedCommands[i].Value.Time.After(sortedCommands[j].Value.Time)
// 		}
// 		return sortedCommands[i].Value.Count > sortedCommands[j].Value.Count
// 	})

// 	return sortedCommands

// }

import (
	"fmt"
	"io"
	"projectshell/src/databases"
	command_model "projectshell/src/services/commands"
	session_user "projectshell/src/services/users"
	"projectshell/src/utils"
	"sort"
	"strings"
	"time"
)

type CommandInfo struct {
	Count int
	Time  time.Time
}

type kv struct {
	Key   string
	Value CommandInfo
}

const CleanCommand = "clean"

func HistoryCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	id := resolveUserID()

	if len(args) > 0 && args[0] == CleanCommand {
		utils.CleanHistory(id)
		return
	}

	history, err := fetchUserHistory(id)
	if err != nil {
		fmt.Fprintf(errorWriter, "Error retrieving commands: %v\n", err)
		return
	}

	if len(history) == 0 {
		fmt.Fprintln(outputWriter, "empty command history")
		return
	}

	commandCounts := aggregateCommandHistory(history)
	printCommandHistory(commandCounts, outputWriter)
}

func resolveUserID() uint {
	currentUser := session_user.GetCurrentUser()
	if currentUser == "" {
		return utils.GetUserId("anonymous")
	}
	return utils.GetUserId(currentUser)
}

func fetchUserHistory(userID uint) ([]command_model.Command, error) {
	db := databases.GetDB()
	var commands []command_model.Command
	err := db.Where("created_by = ?", userID).Find(&commands).Error
	return commands, err
}

func aggregateCommandHistory(commands []command_model.Command) map[string]CommandInfo {
	historyMap := make(map[string]CommandInfo)
	for _, command := range commands {
		if val, exists := historyMap[command.Text]; exists {
			val.Count++
			val.Time = command.CreatedAt
			historyMap[command.Text] = val
		} else {
			historyMap[command.Text] = CommandInfo{Count: 1, Time: command.CreatedAt}
		}
	}
	return historyMap
}

func printCommandHistory(historyMap map[string]CommandInfo, outputWriter io.Writer) {
	sorted := sortHistory(historyMap)

	fmt.Fprintf(outputWriter, "|%-20s|%-10s|\n", "Command", "Count")
	fmt.Fprintln(outputWriter, strings.Repeat("-", 32))

	for _, entry := range sorted {
		fmt.Fprintf(outputWriter, "|%-20s|%-10d|\n", entry.Key, entry.Value.Count)
	}
}

func sortHistory(historyMap map[string]CommandInfo) []kv {
	var sorted []kv
	for k, v := range historyMap {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value.Count == sorted[j].Value.Count {
			return sorted[i].Value.Time.After(sorted[j].Value.Time)
		}
		return sorted[i].Value.Count > sorted[j].Value.Count
	})
	return sorted
}
