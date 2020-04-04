package main

import (
	"../libs"
	"bufio"
	"fmt"
	"io"
	"os"
)

func AreaOfARectangle(reader io.Reader, writer io.Writer) {
	bufreader := bufio.NewReader(reader)

	width, _ := libs.GetPromptedNumber(writer, bufreader, "What is the length of the room in feet? ")
	height, _ := libs.GetPromptedNumber(writer, bufreader, "What is the width of the room in feet? ")

	areaFeet := width * height;
	areaMeters := float64(areaFeet) * 0.09290304

	fmt.Fprintln(writer, "You entered dimensions of", width, "feet by", height, "feet.")

	fmt.Fprintln(writer, "The area is")
	fmt.Fprintln(writer, areaFeet, "square feet")
	fmt.Fprintf(writer, "%.3f square meters\n", areaMeters)
}


func main() {
	AreaOfARectangle(os.Stdin, os.Stdout)
}


