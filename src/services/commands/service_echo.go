// package services

// import (
// 	"fmt"
// )

// func EchoCommand(args []string) {
// 	fmt.Println(args)
// 	inQuotes := false
// 	var words []string
// 	currentWord := ""
// 	// space := " "

// 	for _, arg := range args {
// 		for _, char := range arg {
// 			if char == '\'' {
// 				if inQuotes {
// 					words = append(words, currentWord)
// 					currentWord = ""
// 					inQuotes = false
// 				} else {
// 					inQuotes = true
// 				}
// 			} else if inQuotes {
// 				currentWord += string(char)
// 			}
// 		}
// 	}

// 	if inQuotes {
// 		words = append(words, currentWord)
// 	}

// 	fmt.Println(words)
// }

// func compareTypeWithString(v interface{}, typeName string) bool {
// 	return reflect.TypeOf(v).String() == typeName
// }

package services

import (
	"fmt"
	"os"
	"strings"
)

func EchoCommand(args []string) {
	fmt.Println(args)

	n := len(args[0])
	str := args[0]
	found := false
	for i := 0; i < len(args); i++ {
		for j := 0; j < n; j++ {
			if str[0] == '\'' {
				fmt.Println(str[1 : len(str)-1])
				found = true
				break
			} else if str[j] == '$' {
				command := str[j+1:]
				path := str[:j]
				fmt.Println(command)
				path_ := os.Getenv(command)
				fmt.Println(path + path_)
				found = true
				break
			}
		}

		if !found {
			if len(args) == 1 {
				fmt.Println(str)
			} else {
				result := strings.Join(args, " ")
				fmt.Println(result)

			}
		}
	}
}

// func compareTypeWithString(v interface{}, typeName string) bool {
// 	return reflect.TypeOf(v).String() == typeName
// }
