//Interfaces are data types that represents a set of methods signatures for structs

package main

import (
	"fmt"
	"math"
)

//Define interface
type Shape interface {
	//All the methods that you want this interface to implement
	area() float64 //I declare a method called area which returns a float64
}

//And we can use this interfaces with different structs
type Circle struct { //A circle struct
	x, y, radius float64 //I declare three properties of type float64 for this struct
}

type Rectangle struct {
	width, height float64
}

//These are methods that are going to be used by both of the structs.

//This is a value receiver method because we are just doing a calculation, we are not changing anything. So we don't need the * Because it is not a pointer receiver method that changes things. We are just calculating
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//an area method for the rectangle
func (c Rectangle) area() float64 {
	return c.width * c.height
}

//This is going to receive our interface
func getAreaOfShape(s Shape) float64 { //Interface is a method that can be implemented on different structs.
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{height: 5, width: 10}

	circleArea := getAreaOfShape(circle)       //Interface is a method that can be implemented on different structs.
	rectangleArea := getAreaOfShape(rectangle) //Interface is a method that can be implemented on different structs.

	fmt.Printf("Circle Area: %f\n", circleArea)
	fmt.Printf("Rectangle Area: %f\n", rectangleArea)

}

//Methods that can be implemented on different structs.
