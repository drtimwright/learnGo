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

func TestOutputQuote1(t *testing.T) {
	reader := io.Reader(strings.NewReader("Quote\nTim\n"))
	writer := new(bytes.Buffer)

	EchoQuote(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Tim says, \"Quote\"")
}
