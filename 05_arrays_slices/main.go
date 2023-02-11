package main

import "fmt"

func main() {

	////////////////////////////////////////////////////
	//Arrays
	///////////////////////////////////////////////////

	//With Go Arrays have to have a fixed length and have to have a defined type
	//A slice is like an array that does not have a type and size defined

	//Arrays
	var fruitArr [2]string //An string array with size of 2

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//Declare the array and assign values at the same time
	fruitArray2 := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArray2)

	////////////////////////////////////////////////////
	//Slices
	///////////////////////////////////////////////////
	fruitSlice := []string{"Apple", "Orange", "Grapes", "Bananas"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))

	//It starts at index 1, and stps BEFORE INDEX 3
	fmt.Println(fruitSlice[1:3])
}
