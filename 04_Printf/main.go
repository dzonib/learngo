package main

import "fmt"

func main() {
	// you need to to pass \n after to go to other line.
	fmt.Println("Hi\nhi\nhi")

	// if you want to print backslash you do it like this
	fmt.Println("Hi\\hi\\hi\\")

	// if you want doublequote inside of a string literal
	fmt.Println("\"Hi\"\"Hi\"")

	// var brand = "Google"

	// %q is verb, brand replaces it when printed
	// Println do it automaticly, Printf does not
	// fmt.Printf("%q\n", brand)

	ops, ok, fail := 2500, 543, 433

	//%d is for integer value
	fmt.Printf("Total %d success: %d / %d \n", ops, ok, fail)

}
