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



func RetirementCalculator(reader io.Reader, writer *bytes.Buffer, currentYear int64) {

	bufReader := bufio.NewReader(reader);

	currentAge, _ := libs.GetPromptedNumber(writer, bufReader, "")
	retirementAge, _ := libs.GetPromptedNumber(writer, bufReader, "")

	numYears := retirementAge - currentAge

	fmt.Fprintln(writer, "You have", numYears, "years left until you can retire.")
}

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func Test40YearsLeft(t *testing.T) {
	reader := io.Reader(strings.NewReader("25\n65\n"))
	writer := new(bytes.Buffer)

	RetirementCalculator(reader, writer, 2018)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You have 40 years left until you can retire.\n")
}


func TestUseCurrentAgeInComputation(t *testing.T) {
	reader := io.Reader(strings.NewReader("35\n65\n"))
	writer := new(bytes.Buffer)

	RetirementCalculator(reader, writer, 2018)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You have 30 years left until you can retire.\n")
}

func TestUseRetirementAgeInComputation(t *testing.T) {
	reader := io.Reader(strings.NewReader("25\n55\n"))
	writer := new(bytes.Buffer)

	RetirementCalculator(reader, writer, 2018)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "You have 30 years left until you can retire.\n")
}