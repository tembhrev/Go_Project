package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Vikas",
		LastName:  "Tembhre",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	var myNewVar float32
	myNewVar = 11.1
	log.Println("New float value:", myNewVar)

}
