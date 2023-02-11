package main

import "fmt"

func main() {
	//POINTERs
	//A pointer allows yo to point a memory address or the location of a value that is in a variable

	a := 5
	b := &a //Here it is assgning b to a pointer of a

	fmt.Println(a, b) //B is printing the memory address of the variable a

	fmt.Printf("%T \n", b) //The type of b is a pointer to int

	//If you want to read the value from the memory address you can use the *

	fmt.Println(b) //b is printing the memory address
	//Use * to read val from address
	fmt.Println(*b)
	//b was set to &a, so the next line is equal to the previous one
	fmt.Println(*&a)

	//////////////////////////////////////
	// Change the value with the pointer
	*b = 10
	//Remember b is set to the memory address of a
	fmt.Println(a)
}
