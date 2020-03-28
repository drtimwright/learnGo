package main

import (
	"./tip"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func DoTheTipStuff(reader *io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(*reader)

	fmt.Fprintln(writer, "Input:")
	fmt.Fprint(writer, "  bill amount: ")
	billAmountDollars, _ := getFloat(bufReader)

	fmt.Fprint(writer, "  tip rate: ")
	tipPercentageInt, _ := getInt(bufReader)

	totalAmount, tipAmount := tip.CalculateTip(convertDollarsToCents(billAmountDollars), tipPercentageInt)

	fmt.Fprintln(writer, "Expected result:")
	fmt.Fprintln(writer, "  Tip: "+formatCurrency(tipAmount))
	fmt.Fprintln(writer, "  Total: "+formatCurrency(totalAmount))
}

func convertDollarsToCents(billAmtFloat float64) int64 {
	return int64(billAmtFloat * 100.0)
}

func formatCurrency(tipAmount int64) string {
	return "$" + strconv.FormatFloat(float64(tipAmount)/100, 'f', 2, 64)
}

func getInt(bufReader *bufio.Reader) (int64, error) {
	tipPercentage, _ := bufReader.ReadString('\n')
	tipPercentageInt, error := strconv.ParseInt(strings.Trim(tipPercentage, " \n"), 10, 64)
	return tipPercentageInt, error
}

func getFloat(bufReader *bufio.Reader) (float64, error) {
	billAmt, _ := bufReader.ReadString('\n')
	billAmtFloat, err := strconv.ParseFloat(strings.Trim(billAmt, " \n"), 64)
	return billAmtFloat, err
}

func main() {
	reader := io.Reader(os.Stdin)
	DoTheTipStuff(&reader, os.Stdout)
}
