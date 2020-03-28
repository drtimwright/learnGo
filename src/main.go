package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"./tip"
	"strconv"
	"strings"
)

func DoTheTipStuff(reader *io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(*reader)

	fmt.Fprintln(writer, "Input:")
	fmt.Fprint(writer, "  bill amount: ")
	billAmt, _ := bufReader.ReadString('\n')
	billAmtFloat, _ := strconv.ParseFloat(strings.Trim(billAmt, " \n"), 64)
	billAmtInt := int64(billAmtFloat * 100.0)

	fmt.Fprint(writer, "  tip rate: ")
	tipPercentage, _ := bufReader.ReadString('\n')
	tipPercentageInt, _ := strconv.ParseInt(strings.Trim(tipPercentage, " \n"), 10, 64)

	totalAmount, tipAmount := tip.CalculateTip(billAmtInt, tipPercentageInt)

	formattedTipAmount := strconv.FormatFloat(float64(tipAmount) / 100, 'f', 2, 64)
	formattedTotalAmount := strconv.FormatFloat(float64(totalAmount) / 100, 'f', 2, 64)

	fmt.Fprintln(writer, "Expected result:")
	fmt.Fprintln(writer, "  Tip: $"+formattedTipAmount)
	fmt.Fprintln(writer, "  Total: $" + formattedTotalAmount)
}

func main() {
	reader := io.Reader(os.Stdin)
	DoTheTipStuff(&reader, os.Stdout)
}
