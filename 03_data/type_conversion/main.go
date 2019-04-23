package main

import "fmt"

func main() {
	speed := 100 //int
	force := 2.5 //float64

	//speed = speed * int(force) int with 2 as value so we get wrong result

	// at the beginning speed was declared as int so we must convert whole expresion
	// speed = float64(speed) * force

	// now it all works
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
