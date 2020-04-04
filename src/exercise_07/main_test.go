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

func TestComputeArea(t *testing.T) {
	reader := io.Reader(strings.NewReader("15\n20\n"))
	writer := new(bytes.Buffer)

	AreaOfARectangle(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "The area is\n300 square feet\n27.871 square meters\n")
}

func TestComputeAreaFeet(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n10\n"))
	writer := new(bytes.Buffer)

	AreaOfARectangle(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "The area is\n100 square feet\n")
}

func TestComputeAreaMeters(t *testing.T) {
	reader := io.Reader(strings.NewReader("10\n10\n"))
	writer := new(bytes.Buffer)

	AreaOfARectangle(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "9.290 square meters\n")
}
