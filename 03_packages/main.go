package main

import (
	"fmt"
	"math"

	"github.com/ArmandoOC/golangcrashcourse/03_packages/strutil"
	//We have to write from the root folder
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutil.Reverse("Armando"))
}
