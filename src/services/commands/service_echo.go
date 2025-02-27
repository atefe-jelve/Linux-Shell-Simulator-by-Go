// package services

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func EchoCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
// 	str := args[0]
// 	found := false

// 	for i := 0; i < len(args); i++ {
// 		for j := 0; j < len(args[0]); j++ {

// 			result := []rune{}
// 			if str[0] == '\'' {
// 				for _, v := range str {
// 					if v != '\'' {
// 						result = append(result, v)
// 					}
// 				}
// 				fmt.Println(string(result))
// 				found = true
// 				break

// 			} else if str[0] == '"' {
// 				for p, v := range str {
// 					if v == '\\' {
// 						if p+1 < len(str) && (str[p+1] == '$' || str[p+1] == '\'' || str[p+1] == '"' || str[p+1] == '\\') {
// 							p++
// 						} else {
// 							result = append(result, v)
// 						}
// 					} else {
// 						result = append(result, v)
// 					}
// 				}
// 				fmt.Println(string(result[1 : len(result)-1]))
// 				found = true
// 				break

// 			} else if str[j] == '$' {
// 				command := str[j+1:]
// 				fmt.Println(str[:j] + os.Getenv(command))
// 				found = true
// 				break
// 			}
// 		}

// 		if !found {
// 			if len(args) == 1 {
// 				fmt.Println(str)
// 			} else {
// 				result := strings.Join(args, " ")
// 				fmt.Println(result)

// 			}
// 		}
// 	}
// }


package services

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
	if len(args) == 0 {
		fmt.Fprintln(errorWriter, "No arguments provided")
		return
	}

	// Handle case where there are multiple arguments
	if len(args) > 1 {
		// Join args with spaces for the default case
		result := strings.Join(args, " ")
		fmt.Fprintln(outputWriter, result)
		return
	}

	// Process the first argument
	str := args[0]
	var result string

	// Check for single quotes (remove them)
	if str[0] == '\'' {
		result = removeQuotes(str, '\'')
	} else if str[0] == '"' {
		// Check for double quotes (handle escape sequences)
		result = handleEscapes(str[1:len(str)-1]) // Remove outer quotes and handle escapes
	} else if strings.HasPrefix(str, "$") {
		// Process environment variables
		result = handleEnvVariable(str[1:])
	} else {
		// No special handling, just print the string as-is
		result = str
	}

	// Output the result
	fmt.Fprintln(outputWriter, result)
}

// Helper function to remove quotes from a string
func removeQuotes(s string, quote rune) string {
	var result []rune
	for _, v := range s {
		if v != quote {
			result = append(result, v)
		}
	}
	return string(result)
}

// Helper function to handle escape sequences in double-quoted strings
func handleEscapes(s string) string {
	var result []rune
	i := 0
	for i < len(s) {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case '$', '\'', '"', '\\':
				// Skip the escape character
				i++
			}
		}
		result = append(result, rune(s[i]))
		i++
	}
	return string(result)
}

// Helper function to handle environment variables prefixed with $
func handleEnvVariable(s string) string {
	return os.Getenv(s)
}
