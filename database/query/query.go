package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	Id   int
	Name string
	Age  string
}

func main() {

	db, err := sql.Open("mysql",
		"adnaan:pass@tcp(127.0.0.1:3306)/gotalk")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//INIT OMIT
	rows, err := db.Query("select id, name, age from person where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	person := Person{}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(person.Id, person.Name, person.Age)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//ENDINIT OMIT
}
