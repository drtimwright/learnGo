package main


import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func EchoName(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)
	name, _ := bufReader.ReadString('\n')
	name = strings.Trim(name, " \n")
	fmt.Fprintln(writer, "Hello, " + name + ", nice to meet you!")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestEchoName(t *testing.T) {
	reader := io.Reader(strings.NewReader("Tim\n"))
	writer := new(bytes.Buffer)

	EchoName(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Hello, Tim, nice to meet you!")
}

func TestEchoDifferentName(t *testing.T) {
	reader := io.Reader(strings.NewReader("Brian\n"))
	writer := new(bytes.Buffer)

	EchoName(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Hello, Brian, nice to meet you!")
}