package main

import "fmt"

func foo() (string, int) {
	return "two", 3
}

func main() {
	fmt.Println(foo())

}
