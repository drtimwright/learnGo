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

func TestPrintPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "Input:\n")
}

func TestAllOutput(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "  bill amount:\n")
}

func TestPromptTipPercent(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "  tip rate:\n")
}

func TestPromptExpectedResult(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "Expected result:\n")
}

func TestPromptTipAmount(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "  Tip: $10.00\n")
}

func TestPromptDifferentTipAmount(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n11\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "  Tip: $11.00\n")
}

func TestPromptTotalAmount(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "  Total: $110.00\n")
}

func TestPromptTotalDifferentAmount(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n11\n"))
	writer := new(bytes.Buffer)

	DoTheTipStuff(&reader, writer)

	assert.Contains(t, string(writer.Bytes()), "  Total: $111.00\n")
}

