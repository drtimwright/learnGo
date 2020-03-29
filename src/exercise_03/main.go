package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoQuote(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)

	fmt.Fprint(w, "What is the quote? ")
	quote, _ := reader.ReadString('\n')

	fmt.Fprint(w, "Who said it? ")
	author, _ := reader.ReadString('\n')

	fmt.Fprintln(w, strings.Trim(author, "\n")+" says, \""+strings.Trim(quote, "\n")+"\"")
}

func main() {
	EchoQuote(os.Stdin, os.Stdout)
}
