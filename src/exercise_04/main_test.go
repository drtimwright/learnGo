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

func TestOutput1(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you walk your blue dog quickly? That's hilarious")
}

func TestNounPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("cat\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Enter a noun: ")
}

func TestOutputNoun(t *testing.T) {
	reader := io.Reader(strings.NewReader("cat\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you walk your blue cat quickly? That's hilarious")
}

func TestVerbPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nrun\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Enter a verb: ")
}

func TestOutputVerb(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nrun\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you run your blue dog quickly? That's hilarious")
}

func TestAdjectivePrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nwalk\nred\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Enter an adjective: ")
}

func TestOutputAdjective(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nwalk\nred\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you walk your red dog quickly? That's hilarious")
}

func TestAdverbPrompt(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nwalk\nblue\nquickly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Enter an adverb: ")
}

func TestOutputAdverb(t *testing.T) {
	reader := io.Reader(strings.NewReader("dog\nwalk\nblue\nslowly\n"))
	writer := new(bytes.Buffer)

	MadLib(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Do you walk your blue dog slowly? That's hilarious")
}
