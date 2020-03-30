package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"../libs"
)

func EchoName(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)
	name := libs.GetPromptedString(writer, bufReader, "What is your name? ")
	fmt.Fprintln(writer, "Hello, "+name+", nice to meet you!")
}

func main() {
	EchoName(os.Stdin, os.Stdout)
}
