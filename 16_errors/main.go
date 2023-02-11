package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	resp, err := getSomething(1)
	if err != nil {
		//fmt.Println("There was an error: ", err)
		log.Printf("There was an error: %v\n", err)
	} else {
		fmt.Println("Response: ", resp)
	}
}

//function that recieves an int parameter
//and retrieves an string an an error
func getSomething(a int) (string, error) {

	if a == 1 {
		return "", errors.New("Error calculating something")
		//In this case we are returning empty string an an error
	} else {
		return "My message", nil
		//I return a message and no error (nil)
	}

}
