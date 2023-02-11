package main

import (
	"fmt"
	"strconv"
)

//string converter

//Define person struct
//type Person struct {
//	firstName string
//	lastName  string
//	city      string
//	gender    string
//	age       int
//}

type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

//////////////////////////////////////////////////////////////
//Struct methods//////////////////////////////////////////////
//////////////////////////////////////////////////////////////
// Greeting method (value reciever method)  Is a value receiver function for the structure, because we do not change anything here
//the return value is of type string
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I'm " + strconv.Itoa(person.age)
}

//hasBirthday medhod (Pointer receiver method)  Since it is a pointer receiver we need to put an asterisc on the Struct that we receive
// this method is not going to return anything, we are youst going to change values of the Struct
//thus we do not write a return value type
func (person *Person) hasBirthday() {
	person.age++
}

//getMarried (pointer receiver)
func (person *Person) getMarried(spouceLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouceLastName
	}
}

//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////
func main() {
	//Init Person using struct
	///////////////////////////////////////////
	//We declare a variable equal to our struct
	person1 := Person{firstName: "Samantha", lastName: "Rosas", city: "Tepoztlán", gender: "f", age: 12}

	///// Alternative
	person2 := Person{"Armando", "Olamendi", "Tepoztlán", "m", 12}

	//fmt.Println(person1)

	// Get a single field
	fmt.Println(person1.firstName)
	//We can change the values for that person
	//person1.age = person1.age + 1
	person1.age++

	fmt.Println(person1)

	//Two type of methods: Value receivers and pointers receivers
	//Value recievers are for methods that do calculations, they don't actually change
	//Pointer recievers are for when you actually are changing something.
	//The methods don't go inside of the struct as you might think, they go outside

	person1.hasBirthday() // Pointer receiver method. We change the age of the person. The has birthday method increases the age of the Person
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet()) // Value receiver method calling (We do not change, we just need to obtain info for printing)

	person1.getMarried("González") //Pointer receiver method. We are going to change the last name in case the person is female
	fmt.Println(person1.greet())

	person2.getMarried("Flores") //Armando has married Rocio Flores. But since this person is male, nothing is going to take effect.
	fmt.Println(person2.greet())
}
