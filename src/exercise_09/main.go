package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"../libs"
)

func GallonsOfPaint(reader io.Reader, writer io.Writer) {
	bufReader := bufio.NewReader(reader)

	width, _ := libs.GetPromptedNumber(writer, bufReader, "What is the room width? ")
	length, _ := libs.GetPromptedNumber(writer, bufReader, "What is the room length? ")

	area := width * length
	paint := math.Ceil(float64(area) / float64(350))

	fmt.Fprintln(writer, "You will need to purchase", paint, "gallons of paint to cover", area, "square feet.")
}


func main() {
	GallonsOfPaint(os.Stdin, os.Stdout)
}
