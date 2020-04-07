package main

import (
	"bufio"
	"fmt"
	"io"
	"../libs"
	"os"
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

func main() {
	SlicesOfPizza(os.Stdin, os.Stdout)
}
