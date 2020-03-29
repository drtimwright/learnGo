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
	fmt.Fprint(writer, "Who said it? ")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestQuotePrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("Quote\nTim\n"))
	writer := new(bytes.Buffer)

	EchoQuote(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "What is the quote? ")
}

func TestAuthorPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("Quote\nTim\n"))
	writer := new(bytes.Buffer)

	EchoQuote(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Who said it? ")
}

