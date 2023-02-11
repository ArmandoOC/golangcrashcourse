package main

import "fmt"

func main() {
	//Range is used to loop through arrays, maps, slices

	////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	// Range with slices
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	ids := []int{32, 12, 44, 54, 66}

	//Loop through ids (slice)
	for index, value := range ids {
		//fmt.Printf("Index %d - Value: %d \n", index, value)
		fmt.Println("Index: ", index, " Value: ", value)
	}

	//Notes: When ranging over a slice, two values are returned for each iteration. The first is the index,
	//and the second is a copy of the element at that index
	//In this example I declare that the index is going to be called index and the copy of the element at that
	//index is going to be called value, buit it could be called v or wathever

	//////////////////////////////
	////Not using index  (you replace i with an _ underscore)
	/////////////////////////////
	//Whenever you are not going to use the index
	for _, values := range ids {
		fmt.Printf("Value: %d \n", values)
	}

	// Add ids together
	sum := 0
	for _, values := range ids {
		//sum = sum + id
		sum += values
	}
	fmt.Println("Sum ", sum)

	////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	// Range with maps
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	emails := map[string]string{"Armando": "armando@gmail.com", "David": "david@gmail.com"}
	//If you want to obtain, both Keys and Values
	for key, value := range emails {
		fmt.Printf("%s: %s \n", key, value)
		//fmt.Println("Key: ", key, " Value: ", value)
	}

	//If you want to print only the keys
	for k := range emails {
		fmt.Println("Key: " + k)
	}

	//If you want to print only the values
	for _, v := range emails {
		fmt.Println("Value: " + v)
	}

}
