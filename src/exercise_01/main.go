package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoName(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)
	fmt.Fprint(writer, "What is your name? ")
	name, _ := bufReader.ReadString('\n')
	name = strings.Trim(name, " \n")
	fmt.Fprintln(writer, "Hello, "+name+", nice to meet you!")
}

func main() {
	EchoName(os.Stdin, os.Stdout)
}
