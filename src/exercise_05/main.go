package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"../libs"
	"text/template"
)

type SimpleMathData struct {
	Num1 int64
	Num2 int64
}

func (sm SimpleMathData) Added() int64 {
	return sm.Num1 + sm.Num2
}

func (sm SimpleMathData) Subtracted() int64 {
	return sm.Num1 - sm.Num2
}

func (sm SimpleMathData) Multiplied() int64 {
	return sm.Num1 * sm.Num2
}

func (sm SimpleMathData) Divided() int64 {
	return sm.Num1 / sm.Num2
}

func SimpleMath(reader io.Reader, writer *bytes.Buffer) {
	bufReader := bufio.NewReader(reader)
	firstNumber, _ := strconv.ParseInt(libs.GetPromptedString(writer, bufReader, "What is the first number? "), 10, 64)
	secondNumber, _ := strconv.ParseInt(libs.GetPromptedString(writer, bufReader, "What is the second number? "), 10, 64)

	data := SimpleMathData{firstNumber, secondNumber}

	templ, _ := template.New("simple math template").Parse("{{.Num1}} + {{.Num2}} = {{.Added}}\n{{.Num1}} - {{.Num2}} = {{.Subtracted}}\n{{.Num1}} * {{.Num2}} = {{.Multiplied}}\n{{.Num1}} / {{.Num2}} = {{.Divided}}\n")
	templ.Execute(writer, data)
}

