package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	result, err := sqrt(-16)

	if err != nil {
		//fmt.Println(err)
		fmt.Printf("There was an error: %v\n", err)
		log.Printf("There was an error: %v\n", err)
	} else {
		fmt.Println(result)
	}
}

//recieves a float64, and it returns a float64 or an error
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
