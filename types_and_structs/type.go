package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Vikas",
		LastName:    "Tembhre",
		PhoneNumber: "1 555-5555-1212",
	}

	log.Println(user.FirstName, user.LastName, "Phone number", user.PhoneNumber)
}
