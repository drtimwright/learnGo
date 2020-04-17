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

func TestTwoGallonRoomRoom(t *testing.T) {
	reader := io.Reader(strings.NewReader("200\n2\n"))
	writer := new(bytes.Buffer)

	GallonsOfPaint(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You will need to purchase 2 gallons of paint to cover 400 square feet.")
}

func TestExactlyThreeGallonRoomRoom(t *testing.T) {
	reader := io.Reader(strings.NewReader("350\n3\n"))
	writer := new(bytes.Buffer)

	GallonsOfPaint(reader, writer)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You will need to purchase 3 gallons of paint to cover 1050 square feet.")
}
