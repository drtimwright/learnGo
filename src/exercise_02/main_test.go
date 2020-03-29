package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strconv"
	"strings"
	"testing"
)

func EchoInputStringLength(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	fmt.Fprint(writer, "What is the input string? ")
	name, _ := bufReader.ReadString('\n')
	name = strings.Trim(name, " \n")

	nameLength := int64(len(name))

	fmt.Fprint(writer, name+" has "+strconv.FormatInt(nameLength, 10)+" characters\n")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("Tim\n"))
	writer := new(bytes.Buffer)

	EchoInputStringLength(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "What is the input string? ")
}

func TestStringLength5(t *testing.T) {
	reader := io.Reader(strings.NewReader("Homer\n"))
	writer := new(bytes.Buffer)

	EchoInputStringLength(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Homer has 5 characters\n")
}

func TestStringLength5NotHomer(t *testing.T) {
	reader := io.Reader(strings.NewReader("James\n"))
	writer := new(bytes.Buffer)

	EchoInputStringLength(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "James has 5 characters\n")
}

func TestStringLength3(t *testing.T) {
	reader := io.Reader(strings.NewReader("Tim\n"))
	writer := new(bytes.Buffer)

	EchoInputStringLength(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Tim has 3 characters\n")
}

func TestStringLength0(t *testing.T) {
	reader := io.Reader(strings.NewReader("\n"))
	writer := new(bytes.Buffer)

	EchoInputStringLength(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, " has 0 characters\n")
}
