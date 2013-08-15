package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//*sql.DB
	db, err := sql.Open("mysql",
		"adnaan:pass@tcp(127.0.0.1:3306)/gotalk")
	//no connection has been established yet. validation has not taken place
	if err != nil {
		log.Fatal(err)
	}
	//sql.DB is a long living object. Idiomatic to close it after use
	defer db.Close()

	//INIT OMIT
	// validate
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	//ENDINIT OMIT
}
