package main

import (
	"fmt"
	"io"
	"os"
)

func DoTheTipStuff(reader *io.Reader, writer io.Writer) {
	fmt.Fprintln(writer, "Input:")
}

func main() {
	reader := io.Reader(os.Stdin)
	DoTheTipStuff(&reader, os.Stdout)
}
