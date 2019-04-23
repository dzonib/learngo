package main

import "fmt"

func main() {
	// var speed int
	// var heat float64
	// var off bool
	// var brand string

	// to print the type of the value we use %T
	// fmt.Printf("%T\n", speed)
	// fmt.Printf("%T\n", heat)
	// fmt.Printf("%T\n", off)
	// fmt.Printf("%T\n", brand)

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// %v can print any value
	// fmt.Printf("Planet: %v\n", planet)
	// fmt.Printf("Distance: %v milions kms\n",
	// 	distance)
	// fmt.Printf("Orbital Period: %v days\n",
	// 	orbital)
	// fmt.Printf("Does %v has life: %v\n",
	// 	planet, hasLife)

	// s - string, d - decimal(for int), f - float, t - for booleans
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d milions kms\n",
		distance)

	fmt.Printf("Orbital Period: %f days\n",
		orbital)

	// change precision (how much decimals you want)
	fmt.Printf("Orbital Period: %.0f days\n",
		orbital)

	// change precision agein
	fmt.Printf("Orbital Period: %.1f days\n",
		orbital)

	fmt.Printf("Does %s has life: %t\n",
		planet, hasLife)

	// you can do it by argument index as well
	fmt.Printf("%v is %v away., %[2]v kms! %[1]v omg!\n",
		planet, distance)
}
