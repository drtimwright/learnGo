package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoQuote(r io.Reader, writer io.Writer) {
	reader := bufio.NewReader(r)

	quote := getPromptedString(writer, reader, "What is the quote? ")
	author := getPromptedString(writer, reader, "Who said it? ")

	printQuote(writer, author, quote)
}

func printQuote(writer io.Writer, author string, quote string) (int, error) {
	return fmt.Fprintln(writer, strings.Trim(author, "\n")+" says, \""+strings.Trim(quote, "\n")+"\"")
}

func getPromptedString(w io.Writer, reader *bufio.Reader, prompt string) string {
	fmt.Fprint(w, prompt)
	quote, _ := reader.ReadString('\n')
	return quote
}

func main() {
	EchoQuote(os.Stdin, os.Stdout)
}
