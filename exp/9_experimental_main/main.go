package main

import (
	"fmt"

	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "lenslocked_dev"
	password = "abc123"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	ug, err := models.NewUserGorm(psqlInfo)
	if err != nil {
		panic(err)
	}

	ug.DestructiveReset()

	user := &models.User{
		Name: "Jon Calhoun",
		Email: "jon@calhoun.io",
	}

	fmt.Println("Creating a user...")
	if err := ug.Create(user); err != nil {
		panic(err)
	}
	fmt.Println("Created the user:", user)

	fmt.Println("Retrieving the user w/ID:", user.ID)
	userByID := ug.ByID(user.ID)
	if userByID == nil {
		panic("No user found by ID!")
	}
	fmt.Println("Found the user:", userByID)

	fmt.Println("Updating the user...")
	user.Email = "jon@lenslocked.com"
	if err := ug.Update(user); err != nil {
		panic(err)
	}
	fmt.Println("Updated the user:", user)

	fmt.Println("Retrieving the user w/email:", user.Email)
	userByEmail := ug.ByEmail(user.Email)
	if userByEmail == nil {
		panic("No user found by email!")
	}
	fmt.Println("Found the user:", userByEmail)

	fmt.Println("Deleting the user...")
	if err := ug.Delete(user.ID); err != nil {
		panic(err)
	}
	//verify that the user was deleted
	deletedUser := ug.ByID(user.ID)
	if deletedUser != nil {
		panic("User was not found but should be deleted!")
	}
	fmt.Println("Deleted the user w/ ID:", user.ID)
}
