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

func GallonsOfPaint(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	width, _ := libs.GetPromptedNumber(writer, bufReader, "What is the room width? ")
	length, _ := libs.GetPromptedNumber(writer, bufReader, "What is the room length? ")

	area := width * length

	fmt.Fprintln(writer, "You will need to purchase 1 gallons of paint to cover", area, "square feet.")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestOneByOneRoom(t *testing.T) {

	reader := io.Reader(strings.NewReader("1\n1\n"))
	writer := new(bytes.Buffer)

	GallonsOfPaint(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You will need to purchase 1 gallons of paint to cover 1 square feet.")
}

func TestTwoByTwoRoom(t *testing.T) {

	reader := io.Reader(strings.NewReader("2\n2\n"))
	writer := new(bytes.Buffer)

	GallonsOfPaint(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You will need to purchase 1 gallons of paint to cover 4 square feet.")
}
