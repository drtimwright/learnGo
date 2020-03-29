package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("Tim\n"))
	writer := new(bytes.Buffer)

	EchoName(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "What is your name? ")
}

func TestEchoName(t *testing.T) {
	reader := io.Reader(strings.NewReader("Tim\n"))
	writer := new(bytes.Buffer)

	EchoName(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Hello, Tim, nice to meet you!")
}

func TestEchoDifferentName(t *testing.T) {
	reader := io.Reader(strings.NewReader("Brian\n"))
	writer := new(bytes.Buffer)

	EchoName(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Hello, Brian, nice to meet you!")
}
