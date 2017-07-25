package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	INSERT INTO users(name, email)
	VALUES($1, $2)`,
		"Jon Calhoun", "jon@calhoun.io")

	if err != nil {
		panic(err)
	}

	var id int
	var name, email string
	rows, err := db.Query(`
	SELECT id, name, email
	FROM users
	WHERE email = $1
	OR ID > $2`,
	"jon@calhoun.io", 3)

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&id, &name, &email)
		fmt.Println("ID:", id, "Name:", name, "Email:", email)
	}
	//fmt.Println("Successfully connected!")
	db.Close()
}
