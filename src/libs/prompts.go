package libs

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func GetPromptedString(w io.Writer, reader *bufio.Reader, prompt string) string {
	fmt.Fprint(w, prompt)
	quote, _ := reader.ReadString('\n')
	return strings.Trim(quote, "\n")
}
