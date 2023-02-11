package main

import "fmt"

//           an anonymous function with no name that recieves an int and return an int
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//We set variable sum to function adder
	sum := adder()
	for i := 0; i < 10; i++ {
		//Since we set sum to function adder, we can pass an int
		fmt.Println(sum(i))
	}

}
