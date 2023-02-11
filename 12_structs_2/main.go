package main

import (
	"fmt"
)

type user struct {
	username string
	age      int
	nickname string
}

func main() {
	//Initialize the struct

	jack := user{
		username: "Jack",
		age:      28,
		nickname: "J",
	}

	/*
		jack := user{
			username: "Jack",
			age:      28,
		}
	*/

	fmt.Println(jack)
	fmt.Println(jack.age)

	fmt.Println(jack.getEmail())
	jack.updateUsernameAndNickName("Armando", "A")

	fmt.Println(jack.username)
}

//Struct Methods
//There are value recievers methods and there are pointer receiver methods.

//Value receiver method
func (u user) getEmail() string {
	return u.username + "@gmail.com"
}

//Pointer receiver method.
//A pointer receiver method can update the fields of the struct. All you need to do, to define a pointer receiver method, is just type in an asterisc before the struct
func (u *user) updateUsernameAndNickName(newName string, newNickName string) {
	u.username = newName
	u.nickname = newNickName
}
