package main


import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func EchoName(reader io.Reader, writer io.Writer) {
	fmt.Fprintln(writer, "Hello, Tim, nice to meet you!")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestEchoName(t *testing.T) {
	reader := io.Reader(strings.NewReader("100\n10\n"))
	writer := new(bytes.Buffer)

	EchoName(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Hello, Tim, nice to meet you!")
}