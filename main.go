package main

import (
	"database/sql"

	"github.com/alochym01/thecodecamp/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "foo.db")

	if err != nil {
		panic(err)
	}
	// ID       string
	// Email    string
	// Password string
	// Role     string
	// Status   string
	// sqltable := fmt.Sprintf("CREATE TABLE users (id CHARACTER (32) NOT NULL PRIMARY KEY, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, role VARCHAR(255) NOT NULL, status VARCHAR(255) NOT NULL)")
	// db.Exec(sqltable)
	router.Start(db)
}
