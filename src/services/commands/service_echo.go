package services

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
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
				fmt.Fprintln(outputWriter, string(result))
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
				fmt.Fprintln(outputWriter, string(result[1:len(result)-1]))
				found = true
				break

			} else if str[j] == '$' {
				command := str[j+1:]
				fmt.Fprintln(outputWriter, str[:j]+os.Getenv(command))
				found = true
				break
			}
		}

		if !found {
			if len(args) == 1 {
				fmt.Fprintln(outputWriter, str)
			} else {
				result := strings.Join(args, " ")
				fmt.Fprintln(outputWriter, result)

			}
		}
	}
}

// package services

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func EchoCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {

// 	if len(args) == 0 {
// 		fmt.Fprintln(errorWriter, "No arguments provided")
// 		return
// 	}

// 	if len(args) > 1 {
// 		result := strings.Join(args, " ")
// 		fmt.Fprintln(outputWriter, result)
// 		return
// 	}

// 	str := args[0]
// 	var result string

// 	if str[0] == '\'' {
// 		result = removeQuotes(str, '\'')
// 	} else if str[0] == '"' {
// 		result = handleEscapes(str[1:len(str)-1])
// 	} else if strings.HasPrefix(str, "$") {
// 		result = handleEnvVariable(str[1:])
// 	} else {
// 		result = str
// 	}

// 	fmt.Fprintln(outputWriter, result)
// }

// func removeQuotes(s string, quote rune) string {
// 	var result []rune
// 	for _, v := range s {
// 		if v != quote {
// 			result = append(result, v)
// 		}
// 	}
// 	return string(result)
// }

// func handleEscapes(s string) string {
// 	var result []rune
// 	i := 0
// 	for i < len(s) {
// 		if s[i] == '\\' && i+1 < len(s) {
// 			switch s[i+1] {
// 			case '$', '\'', '"', '\\':
// 				i++
// 			}
// 		}
// 		result = append(result, rune(s[i]))
// 		i++
// 	}
// 	return string(result)
// }

// func handleEnvVariable(s string) string {
// 	return os.Getenv(s)
// }
