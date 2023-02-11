package main

import (
	"errors"
	"fmt"
)

func main() {
	total, err := sumOf(1, 4)

	if err != nil {
		fmt.Println("There was an error ", err)
	} else {
		fmt.Println("Total: ", total)
	}
}

//Recieves start and end in int format. It outputs an int and an error
func sumOf(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("Start is greter than end")
	}

	total := 0
	for i := start; i < end; i++ {
		total = total + i
		//total += i
	}
	return total, nil
}
