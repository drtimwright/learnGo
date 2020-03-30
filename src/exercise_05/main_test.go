package main

import (
	"../libs"
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strconv"
	"strings"
	"testing"
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

func printComputation(writer *bytes.Buffer, firstNumber int64, operator string, secondNumber int64, total int64) (int, error) {
	return fmt.Fprintln(writer, strconv.FormatInt(firstNumber, 10), operator, strconv.FormatInt(secondNumber, 10), "=", strconv.FormatInt(total, 10))
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestFirstPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n10\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "What is the first number? ")
}

func TestSecondPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n10\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "What is the second number? ")
}

func TestAddition(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n10\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "10 + 10 = 20\n")
}

func TestAddition2(t *testing.T) {
	reader := io.Reader(strings.NewReader("20\n10\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "20 + 10 = 30\n")
}

func TestAddition3(t *testing.T) {
	reader := io.Reader(strings.NewReader("20\n30\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "20 + 30 = 50\n")
}

func TestSubtraction(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n5\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "10 - 5 = 5\n")
}

func TestMultiply(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n5\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "10 * 5 = 50\n")
}

func TestDivide(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n5\n"))
	writer := new(bytes.Buffer)

	SimpleMath(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "10 / 5 = 2\n")
}
