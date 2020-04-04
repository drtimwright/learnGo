package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
	"../libs"
)

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func AreaOfARectangle(reader io.Reader, writer io.Writer) {
	bufreader := bufio.NewReader(reader)

	width, _ := libs.GetPromptedNumber(writer, bufreader, "What is the length of the room in feet? ")
	height, _ := libs.GetPromptedNumber(writer, bufreader, "What is the width of the room in feet? ")

	fmt.Fprintln(writer, "You entered dimensions of", width, "feet by", height, "feet.")
}

func TestPrintDimensions(t *testing.T) {
	reader := io.Reader(strings.NewReader("15\n20\n"))
	writer := new(bytes.Buffer)

	AreaOfARectangle(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You entered dimensions of 15 feet by 20 feet.\n")
}

func TestPrintDimensionsDifferentSize(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n30\n"))
	writer := new(bytes.Buffer)

	AreaOfARectangle(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You entered dimensions of 10 feet by 30 feet.\n")
}
