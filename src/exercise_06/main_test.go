package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strconv"
	"strings"
	"testing"
	"../libs"
)

func RetirementCalculator(reader io.Reader, writer *bytes.Buffer, currentYear int64) {

	bufReader := bufio.NewReader(reader)

	currentAge, _ := libs.GetPromptedNumber(writer, bufReader, "What is your current age? ")
	retirementAge, _ := libs.GetPromptedNumber(writer, bufReader, "At what age would you like to retire? ")

	numYears := retirementAge - currentAge
	retirementDate := currentYear + numYears

	fmt.Fprintln(writer, "You have", numYears, "years left until you can retire.")
	fmt.Fprintln(writer, "It's 2015, so you can retire in", strconv.FormatInt(retirementDate, 10)+".")
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

func TestRetirementDate(t *testing.T) {
	reader := io.Reader(strings.NewReader("25\n65\n"))
	writer := new(bytes.Buffer)

	RetirementCalculator(reader, writer, 2015)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "It's 2015, so you can retire in 2055.\n")
}

func TestRetirementDate2045(t *testing.T) {
	reader := io.Reader(strings.NewReader("25\n55\n"))
	writer := new(bytes.Buffer)

	RetirementCalculator(reader, writer, 2015)

	actual := string(writer.Bytes())
	assert.Contains(t, actual, "It's 2015, so you can retire in 2045.\n")
}
