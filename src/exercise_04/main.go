package main

import (
	"../libs"
	"bufio"
	"io"
	"os"
	"text/template"
)

type MadLibData struct {
	Noun      string
	Verb      string
	Adjective string
	Adverb    string
}

func MadLib(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	noun := libs.GetPromptedString(writer, bufReader, "Enter a noun: ")
	verb := libs.GetPromptedString(writer, bufReader, "Enter a verb: ")
	adjective := libs.GetPromptedString(writer, bufReader, "Enter an adjective: ")
	adverb := libs.GetPromptedString(writer, bufReader, "Enter an adverb: ")

	data := MadLibData{noun, verb, adjective, adverb}

	templ, _ := template.New("mad lib template").Parse("Do you {{ .Verb }} your {{ .Adjective }} {{ .Noun }} {{ .Adverb }}? That's hilarious")

	templ.Execute(writer, data)
}


func main() {
	MadLib(os.Stdin, os.Stdout)
}
