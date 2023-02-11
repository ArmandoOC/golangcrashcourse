package main

import "fmt"

//Parameter and parameter type   Return type
func greeting(name string) string {
	return "Hello " + name
}

//func getSum(num1 int, num2 int) int {
//The last line is the same as the following. In the following its is a way of doint it simple
func getSum(num1, num2 int) int {
	return num1 + num2

}

func main() {
	fmt.Println(greeting("Armando"))

	fmt.Println(getSum(1, 2))
}
