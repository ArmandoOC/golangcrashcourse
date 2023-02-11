package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Starting the application...")
	output := flag.Bool("output", false, "Should there be output?")
	input := flag.String("input", "file.csv", "The path to the input file")
	flag.Parse()
	fmt.Println(*output)
	fmt.Println(*input)
}

//go run main.go --help
//https://www.youtube.com/watch?v=v_8584-jm7I
