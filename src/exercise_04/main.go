package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"../libs"
)

func MadLib(reader io.Reader, writer *bytes.Buffer) {
	bufReader := bufio.NewReader(reader)

	noun := libs.GetPromptedString(writer, bufReader, "Enter a noun: ")
	verb := libs.GetPromptedString(writer, bufReader, "Enter a verb: ")
	adjective := libs.GetPromptedString(writer, bufReader, "Enter an adjective: ")
	adverb := libs.GetPromptedString(writer, bufReader, "Enter an adverb: ")

	fmt.Fprintln(writer, "Do you", verb, "your", adjective, noun, adverb+"? That's hilarious")
}

