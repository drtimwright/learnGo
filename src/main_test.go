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

	assert.Contains(t, string(writer.Bytes()), "Input:")
}
