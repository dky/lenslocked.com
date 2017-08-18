package main

import (
	"fmt"
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

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.Close()
}
