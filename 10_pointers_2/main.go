package main

import "fmt"

func main() {
	x := 5
	pointer := &x

	fmt.Println(pointer)  //It will print a memory address
	fmt.Println(*pointer) //It will print the value of that memory address

	*pointer = 8 //This updates the value of x.
	fmt.Println(x)

	setTo(&x, 1000) //we pass the pointer to the variable x
	//setTo(pointer, 1000)   //This is the same, but it uses the variable that I created

	fmt.Println(x)
}

//How to pass pointers to a function

//The data type of a pointer is *int
func setTo(addr *int, newValue int) {

	*addr = newValue
}

//That's how we use pointers and receivers in Go and that's how you pass pointers in the function.
