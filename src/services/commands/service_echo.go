package services

import (
	"fmt"
	"os"
	"strings"
)

func EchoCommand(args []string) {
	// fmt.Println(args)
	n := len(args[0])
	str := args[0]
	found := false
	for i := 0; i < len(args); i++ {
		for j := 0; j < n; j++ {
			result := []rune{}
			if str[0] == '\'' {
				// result := []rune{}
				for _, v := range str {
					if v != '\'' {
						result = append(result, v)
					}
				}
				fmt.Println(string(result))
				found = true
				break
			} else if str[0] == '"' {
				for p, v := range str {
					if v == '\\' {
						if p+1 < len(str) && (str[p+1] == '$' || str[p+1] == '\'' || str[p+1] == '"' || str[p+1] == '\\') {
							p++
						} else {
							result = append(result, v)
						}
					} else {
						result = append(result, v)
					}
				}
				fmt.Println(string(result[1 : len(result)-1]))
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
