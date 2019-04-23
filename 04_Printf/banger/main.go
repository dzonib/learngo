package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputString := os.Args[1]

	for i := len(inputString); i > 0; i-- {
		inputString += "!"
	}

	fmt.Printf("%s", strings.ToUpper(inputString))

	// Or

	msg := inputString + strings.Repeat("!", len(inputString))

	fmt.Printf("%s", strings.ToUpper(msg))

}
