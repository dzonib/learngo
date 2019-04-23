package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	// takes value and how much bytes we want (32 or 64)
	// returns float and error
	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feet  is %g meters.\n", feet, meters)
}
