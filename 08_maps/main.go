package main

import "fmt"

func main() {

	////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////
	//Define a map first and then assign values to it

	//Define map
	//A map is a key value pair

	// Se definen con    make(map)[DataTypeOfKeys]DataTypeOfValues)
	emails := make(map[string]string)
	fmt.Println("Hello World")

	//Assign key values to map
	emails["Armando"] = "armando@gmail.com"
	emails["David"] = "david@gmail.com"
	emails["Leticia"] = "lety@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Armando"])

	fmt.Println(len(emails))

	//Delete from map
	delete(emails, "Armando") //Se le pasa el mapa y la clave del record a borrar
	fmt.Println(emails)

	////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////
	//Or define a map and assign values at the same time

	//Assign key value pairs when you declare the map
	//Declare map and add key values

	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	//I can also continue assigning key values to map
	emails2["Juan"] = "juan@gmail.com"

	fmt.Println(emails2)

	//Map
	map1 := make(map[string]string) //map compuesto de claves tipo string y valores tipo string
	map1["A"] = "Alpha"
	map1["B"] = "Beta"
	map1["C"] = "Crunch"

	//Loop through a map

	for key, value := range map1 {
		fmt.Println("Key: ", key, " Value: ", value)
	}

}
