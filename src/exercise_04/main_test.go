package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
	"../libs"
)

func MadLib(reader io.Reader, writer *bytes.Buffer) {
	bufReader := bufio.NewReader(reader)

	noun := libs.GetPromptedString(writer, bufReader, "Enter a noun: ")
	verb := libs.GetPromptedString(writer, bufReader, "Enter a verb: ")

	fmt.Fprintln(writer, "Do you", verb, "your blue", noun, "quickly? That's hilarious")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestOutput1(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you walk your blue dog quickly? That's hilarious")
}

func TestNounPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("cat\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Enter a noun: ")
}

func TestOutputNoun(t *testing.T) {
	reader := io.Reader(strings.NewReader("cat\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you walk your blue cat quickly? That's hilarious")
}

func TestVerbPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("cat\nrun\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Enter a verb: ")
}

func TestOutputVerb(t *testing.T) {
	reader := io.Reader(strings.NewReader("cat\nrun\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you run your blue cat quickly? That's hilarious")
}
