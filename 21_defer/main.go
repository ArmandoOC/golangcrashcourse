package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	random()

}

func random() {
	defer fmt.Println("0")
	fmt.Println("1")
	panic("We have panic")
	fmt.Println("2")
}

//Defer is for defining that something must run after the statements next to it have run.

//defer function will run even if the function panics at the end. So If you open a file, after it you might want to say
//defer file close. So no mather what happens the file will be closed at the end.
