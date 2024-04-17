package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 32})
	users = append(users, User{"Slvia", "jones", "slvia@jones.com", 37})
	users = append(users, User{"Bill", "Will", "bill@will.com", 39})
	users = append(users, User{"kathrine", "Smith", "kathrine@smith.com", 28})
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
