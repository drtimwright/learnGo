package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"../libs"
)

func EchoQuote(r io.Reader, writer io.Writer) {
	reader := bufio.NewReader(r)

	quote := libs.GetPromptedString(writer, reader, "What is the quote? ")
	author := libs.GetPromptedString(writer, reader, "Who said it? ")

	printQuote(writer, author, quote)
}

func printQuote(writer io.Writer, author string, quote string) (int, error) {
	return fmt.Fprintln(writer, author, "says, \""+quote+"\"")
}

func main() {
	EchoQuote(os.Stdin, os.Stdout)
}
