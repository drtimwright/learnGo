package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"../libs"
)

func RetirementCalculator(reader io.Reader, writer *bytes.Buffer, currentYear int64) {

	bufReader := bufio.NewReader(reader)

	currentAge, _ := libs.GetPromptedNumber(writer, bufReader, "What is your current age? ")
	retirementAge, _ := libs.GetPromptedNumber(writer, bufReader, "At what age would you like to retire? ")

	numYears := retirementAge - currentAge
	retirementDate := currentYear + numYears

	fmt.Fprintln(writer, "You have", numYears, "years left until you can retire.")
	fmt.Fprintln(writer, "It's", strconv.FormatInt(currentYear, 10)+",", "so you can retire in", strconv.FormatInt(retirementDate, 10)+".")
}

