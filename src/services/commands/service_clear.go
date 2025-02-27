
package services

import (
    "fmt"
    "io"
)

func ClearCommand(args []string, outputWriter io.Writer, errorWriter io.Writer) {
    const clearScreen = "\033[H\033[2J"
    fmt.Fprint(outputWriter, clearScreen)
}
