package main

import (
	"fmt"
	"runtime"

	// import "fmt" as f
	f "fmt"
)

func main() {
	// you dont have to import fuction if it is in same package scope
	fmt.Println(bye())
	f.Println("yo")

	check(4)
	f.Println(runtime.NumCPU())
}
