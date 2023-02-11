package main

import "fmt"

//second way of declaring a variable
//global variable
var name3 = "Peter"

//Short hand method. But this type of assignment we cannot do it outside of a function
//name4 := "Peter"

func main() {
	var name string = "Brad"
	var name2 = "Peter"
	var age int = 37
	var age2 = 37
	var isCool = true
	
	const isLate = true

	//I cannot modify a constant.
	//isLate = false

	//cast
	var age3 int32 = 37

	fmt.Println(name, age)
	fmt.Println(name2, age2)

	//Print type  See golang.org/pkg/fmt/
	//fmt.Printf("%T\n", name)
	//fmt.Printf("%T\n", age)

	isCool = false
	fmt.Println(isCool)
	fmt.Printf("%T\n", age3)

	fmt.Printf("%T\n", isCool)

	///Third way of declaring a varible: the short method.
	//This method you can only usint within a function.
	name5 := "Armando"
	//This assumes that the type is string
	fmt.Printf("%T\n", name5)

	//float64. It is infered that it is a float64
	size := 1.3
	fmt.Printf("%T\n", size)

	//If you want to force a float32, this is the way
	var size2 float32 = 2.3
	fmt.Println(size2)
	fmt.Printf("%T\n", size2)

	//Short way of declaring assigning variables at the same time
	//Instead of doing this:
	//name6 := "Brad"
	//email := "brad@gmail.com"
	//we do this:
	name6, email := "Brad", "brad@gmail.com"
	fmt.Println(name6, email)

}
