package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//INIT OMIT
func main() {
	//*sql.DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	//no connection has been established yet. validation has not taken place
	if err != nil {
		log.Fatal(err)
	}
	//sql.DB is a long living object. Idiomatic to close it after use
	defer db.Close()
}

//ENDINIT OMIT
