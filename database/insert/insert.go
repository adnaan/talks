package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql",
		"adnaan:pass@tcp(127.0.0.1:3306)/gotalk")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//INIT OMIT
	stmt, err := db.Prepare("INSERT INTO person(name,age) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec("Gopher", 4)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	//ENDINIT OMIT
}
