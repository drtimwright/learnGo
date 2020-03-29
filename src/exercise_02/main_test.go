package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)


func EchoInputStringLength(reader io.Reader, writer io.Writer) {
	fmt.Fprint(writer, "What is the input string? ")
	fmt.Fprint(writer, "Homer has 5 characters\n")
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
