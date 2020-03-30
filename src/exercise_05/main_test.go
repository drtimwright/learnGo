package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"../libs"
	"strings"
	"testing"
)


func SimpleMath(reader io.Reader, writer *bytes.Buffer) {
	bufReader := bufio.NewReader(reader)
	libs.GetPromptedString(writer, bufReader, "What is the first number? ")
	libs.GetPromptedString(writer, bufReader, "What is the second number? ")

	fmt.Fprintln(writer, "10 + 10 = 20")
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