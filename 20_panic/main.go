package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	random()
	fmt.Println("This won't be reached")
}

func random() {
	panic("We have panic right now")
}
