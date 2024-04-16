package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	otherMap := make(map[string]int)

	otherMap["First"] = 1
	otherMap["Second"] = 2

	log.Println(myMap["dog"])

	log.Println(otherMap["First"], otherMap["Second"])

}
