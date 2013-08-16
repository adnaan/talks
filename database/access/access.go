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
		"adnaan:pass@tcp(127.0.0.1:3306)/gotalk")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

//ENDINIT OMIT
