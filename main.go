package main

import (
	"log"

	"github.com/dracco-eu-logistics/api/user"
)

func main() {

	// Register a new user
	user := user.NewUser()
	log.Println(user)
}
