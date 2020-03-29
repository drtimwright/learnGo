package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func EchoInputStringLength(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	fmt.Fprint(writer, "What is the input string? ")
	name, _ := bufReader.ReadString('\n')
	name = strings.Trim(name, " \n")

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
