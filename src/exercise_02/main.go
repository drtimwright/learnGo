package main

import (
	"../libs"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func EchoInputStringLength(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	name := libs.GetPromptedString(writer, bufReader, "What is the input string? ")

	nameLength := int64(len(name))
	if nameLength > 0 {
		fmt.Fprintln(writer, name+" has "+strconv.FormatInt(nameLength, 10)+" characters\n")
	} else {
		fmt.Fprintln(writer, "You must enter a non-empty string")
	}
}

func main() {
	EchoInputStringLength(os.Stdin, os.Stdout)
}
