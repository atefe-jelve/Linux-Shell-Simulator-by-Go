package services

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

func HistoryCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

	db := databases.GetDB()
	var id uint

	currentUser := session_user.GetCurrentUser()
	if currentUser == "" {
		id = utils.GetUserId("anonymous")
	} else {
		id = utils.GetUserId(currentUser)
	}

	if len(args) > 0 && args[0] == "clean" {
		utils.CleanHistory(id)
		return
	}

	var commandObjs []command_model.Command
	if err := db.Where("created_by = ?", id).Find(&commandObjs).Error; err != nil {
		fmt.Printf("Error retrieving commands: %v\n", err)
		return
	}

	if len(commandObjs) == 0 {
		fmt.Println("empty command history")
		return
	}

	commandsMap := make(map[string]CommandInfo)
	for _, commandObj := range commandObjs {
		if val, exists := commandsMap[commandObj.Text]; exists {
			val.Count++
			val.Time = commandObj.CreatedAt
			commandsMap[commandObj.Text] = val
		} else {
			commandsMap[commandObj.Text] = CommandInfo{
				Count: 1,
				Time:  commandObj.CreatedAt,
			}
		}
	}

	sortedCommands := sortedHistory(commandsMap)

	fmt.Printf("|%-20s|%-10s|\n", "Command", "Count")
	fmt.Println(strings.Repeat("-", 32))

	for _, kv := range sortedCommands {
		fmt.Printf("|%-20s|%-10d|\n", kv.Key, kv.Value.Count)
	}
}

func sortedHistory(commandsMap map[string]CommandInfo) []kv {

	var sortedCommands []kv
	for k, v := range commandsMap {
		sortedCommands = append(sortedCommands, kv{k, v})
	}

	sort.Slice(sortedCommands, func(i, j int) bool {
		if sortedCommands[i].Value.Count == sortedCommands[j].Value.Count {
			return sortedCommands[i].Value.Time.After(sortedCommands[j].Value.Time)
		}
		return sortedCommands[i].Value.Count > sortedCommands[j].Value.Count
	})

	return sortedCommands

}
