package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "James",
		"last_name": "Bond",
		"hair_color": "black",
		"has_dog": false
	},
	{
		"first_name": "Steve",
		"last_name": "Williams",
		"hair_color": "black",
		"has_dog": true
	}
]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json:", err)
	}

	log.Printf("Unmarshalled json: %v", unmarshalled)

	//write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "Bond"
	m1.HairColor = "black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Lucky"
	m2.LastName = "Williams"
	m2.HairColor = "black"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "		")
	if err != nil {
		log.Println("Error marshalling json:", err)
	}

	fmt.Println(string(newJson))

}
