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

func SlicesOfPizza(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	numPeople, _ := libs.GetPromptedNumber(writer, bufReader, "How many people? ")
	numPizzas, _ := libs.GetPromptedNumber(writer, bufReader, "How many pizzas do you have? ")

	numSlices := numPizzas * 8 / numPeople
	remainder := numPizzas * 8 % numPeople

	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, numPeople, "people with", numPizzas, "pizzas")
	fmt.Fprintln(writer, "Each person gets", numSlices, "pieces of pizza.")
	fmt.Fprintln(writer, "There are", remainder, "leftover pieces.")
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

func TestDiffNumPizzas(t *testing.T) {
	reader := io.Reader(strings.NewReader("4\n1\n"))
	writer := new(bytes.Buffer)

	SlicesOfPizza(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "4 people with 1 pizzas\n")
}

func TestDiffNumSlicesStillNoRemainder(t *testing.T) {
	reader := io.Reader(strings.NewReader("4\n2\n"))
	writer := new(bytes.Buffer)

	SlicesOfPizza(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Each person gets 4 pieces of pizza.\n")
	assert.Contains(t, actual, "There are 0 leftover pieces.\n")
}

func TestDiffNumSlicesStillWithRemainder(t *testing.T) {
	reader := io.Reader(strings.NewReader("3\n1\n"))
	writer := new(bytes.Buffer)

	SlicesOfPizza(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "Each person gets 2 pieces of pizza.\n")
	assert.Contains(t, actual, "There are 2 leftover pieces.\n")
}
