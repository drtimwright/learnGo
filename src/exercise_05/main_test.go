package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"../libs"
	"strconv"
	"strings"
	"testing"
)

func SimpleMath(reader io.Reader, writer *bytes.Buffer) {
	bufReader := bufio.NewReader(reader)
	firstNumber, _ := strconv.ParseInt(libs.GetPromptedString(writer, bufReader, "What is the first number? "), 10, 64)
	libs.GetPromptedString(writer, bufReader, "What is the second number? ")

	total := firstNumber + 10

	fmt.Fprintln(writer, strconv.FormatInt(firstNumber, 10), "+ 10 =", strconv.FormatInt(total, 10))
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
