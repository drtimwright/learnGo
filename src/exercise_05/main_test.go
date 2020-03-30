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
