package services

import (
	"fmt"
	"os"
	"strings"
)

func EchoCommand(args []string) {
	str := args[0]
	found := false

	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[0]); j++ {

			result := []rune{}
			if str[0] == '\'' {
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
				fmt.Println(str[:j] + os.Getenv(command))
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
