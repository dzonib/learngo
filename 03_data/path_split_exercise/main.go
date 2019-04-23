package main

import (
	"fmt"
	"path"
)

func main() {
	// var dir, file string

	// dir, file = path.Split("css/main.css")
	// fmt.Println("Dir", dir)
	// fmt.Println("File", file)

	// or

	// skip directory and without string declaration
	_, file := path.Split("css/main.js")

	fmt.Println("File", file)
}
