package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	//a, b := 10, 0
	var numero_a, numero_b int
	fmt.Print("Ingresa primer numero: ")
	fmt.Scanf("%d", &numero_a)
	fmt.Print("Ingresa segundo numero: ")
	fmt.Scanf("%d", &numero_b)

	result, err := divide(numero_a, numero_b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", numero_a, numero_b, result)
}