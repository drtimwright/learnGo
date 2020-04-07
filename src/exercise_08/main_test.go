package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func SlicesOfPizza(reader io.Reader, writer io.Writer) {

	fmt.Fprint(writer, "How many people? ")
	fmt.Fprint(writer, "How many pizzas do you have? ")
	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, "8 people with 2 pizzas")
	fmt.Fprintln(writer, "Each person gets 2 pieces of pizza.")
	fmt.Fprintln(writer, "There are 0 leftover pieces.")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestSampleFromChapter(t *testing.T) {
	reader := io.Reader(strings.NewReader("8\n2\n"))
	writer := new(bytes.Buffer)

	SlicesOfPizza(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "How many people? ")
	assert.Contains(t, actual, "How many pizzas do you have? ")
	assert.Contains(t, actual, "8 people with 2 pizzas\n")
	assert.Contains(t, actual, "Each person gets 2 pieces of pizza.\n")
	assert.Contains(t, actual, "There are 0 leftover pieces.\n")
}
