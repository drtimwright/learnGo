package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func EchoQuote(reader io.Reader, writer io.Writer) {

	fmt.Fprint(writer, "What is the quote? ")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("Quote\nTim\n"))
	writer := new(bytes.Buffer)

	EchoQuote(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "What is the quote? ")
}

