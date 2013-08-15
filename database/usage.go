package main

//INIT OMIT

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Importing the driver anonymously under the alies "_"
//None of it's exported methods are visible
//Driver "registers" itself as being available to database/sql

//ENDINIT OMIT
