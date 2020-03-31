package libs

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func GetPromptedString(w io.Writer, reader *bufio.Reader, prompt string) string {
	fmt.Fprint(w, prompt)
	quote, _ := reader.ReadString('\n')
	return strings.Trim(quote, "\n")
}

func GetPromptedNumber(w io.Writer, reader *bufio.Reader, prompt string) (int64, error) {
	readString := GetPromptedString(w, reader, prompt)
	return strconv.ParseInt(readString, 10, 64)
}
