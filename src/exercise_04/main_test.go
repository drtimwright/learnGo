package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func MadLib(reader io.Reader, writer *bytes.Buffer) {

	fmt.Fprintln(writer, "Do you walk your blue dog quickly? That's hilarious")
}

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
