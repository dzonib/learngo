package main

import (
	"fmt"
	// import "fmt" as f
)

// if you declare variables outside you will not get error if unused
// in function, if you are not using them you will get error

// you can declare multiple variables like this
var (
	speed string
	hello string
	age   int
)

// or like this
var city, street string

func main() {

	// declare multiple variables with different type (short)
	test, hungry := "lalal", true

	name, lastname := "Nikola", "Tesla"

	fmt.Println(test, hungry, name, lastname)
}
