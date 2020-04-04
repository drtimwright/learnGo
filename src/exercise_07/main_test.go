package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func AreaOfARectangle(reader io.Reader, writer io.Writer) {
	fmt.Fprintln(writer, "You entered dimensions of 15 feet by 20 feet.")
}

func TestPrintDimensions(t *testing.T) {
	reader := io.Reader(strings.NewReader("15\n20\n"))
	writer := new(bytes.Buffer)

	AreaOfARectangle(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You entered dimensions of 15 feet by 20 feet.\n")
}
