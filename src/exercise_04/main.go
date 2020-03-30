package main

import (
	"bufio"
	"bytes"
	"io"
	"../libs"
	"text/template"
)

type MadLibData struct {
	Noun      string
	Verb      string
	Adjective string
	Adverb    string
}

func MadLib(reader io.Reader, writer *bytes.Buffer) {
	bufReader := bufio.NewReader(reader)

	noun := libs.GetPromptedString(writer, bufReader, "Enter a noun: ")
	verb := libs.GetPromptedString(writer, bufReader, "Enter a verb: ")
	adjective := libs.GetPromptedString(writer, bufReader, "Enter an adjective: ")
	adverb := libs.GetPromptedString(writer, bufReader, "Enter an adverb: ")

	data := MadLibData{noun, verb, adjective, adverb}

	templ, _ := template.New("mad lib template").Parse("Do you {{ .Verb }} your {{ .Adjective }} {{ .Noun }} {{ .Adverb }}? That's hilarious")

	templ.Execute(writer, data)
}
