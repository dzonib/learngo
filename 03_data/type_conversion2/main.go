package main

import "fmt"

func main() {
	var apple int
	var orange int32

	// types are different so you can not do this
	// apple = orange

	// now it works
	apple = int(orange)

	// 65 is character for A
	apple = 65

	fmt.Println(apple)

	testStringConversion := string(apple)

	fmt.Println(testStringConversion)
}
