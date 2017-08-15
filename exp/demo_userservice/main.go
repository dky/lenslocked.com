package main

import (
	"fmt"
	"lenslocked.com/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "lenslocked_dev"
	password = "abc123"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	user := models.User {
		Name: "Hillary Duff",
		Email: "Hillary@aol.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("foundUser")

	/*
	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	*/
}
