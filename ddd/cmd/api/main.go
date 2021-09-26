package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/alochym01/thecodecamp_1/domain/users"
	"github.com/alochym01/thecodecamp_1/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Command-line Flags
	addr := flag.String("addr", ":8080", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line
	// This reads in the command-line flag value and assigns it to the addr
	flag.Parse()

	db := SetupSqliteConn("foo.db")
	defer db.Close()

	app := router.Router(db)
	app.Run(*addr)
}

func SetupSqliteConn(f string) *sql.DB {
	dbConn, err := sql.Open("sqlite3", f)

	if err != nil {
		panic(err)
	}

	return dbConn
}

func mockDB(db *sql.DB) {
	// ID       string `json:"id"`
	// Email    string `json:"email"`
	// Password string `json:"password"`
	// Role     string `json:"role"`
	// Status   string `json:"status"`
	sqltable := `CREATE TABLE users (id CHARACTER (32) NOT NULL PRIMARY KEY, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, role VARCHAR(255) NOT NULL, status VARCHAR(255) NOT NULL)`

	// Execute SQL Statements
	_, err := db.Exec(sqltable)
	if err != nil {
		fmt.Println(err.Error())
	}
	var temp = []users.User{
		{ID: "1000", Email: "hadn4@fpt.com.vn", Password: "Alochym@123", Role: "Husband", Status: "1"},
		{ID: "1001", Email: "1nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
		{ID: "1002", Email: "2nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "0"},
		{ID: "1003", Email: "3nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
		{ID: "1004", Email: "4nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "0"},
		{ID: "1005", Email: "5nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
		{ID: "1006", Email: "6nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
		{ID: "1007", Email: "7nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "10"},
		{ID: "1008", Email: "8nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
		{ID: "1009", Email: "9nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	}
	for i := range temp {
		sqlstmt := `INSERT INTO users(id, email, password, role, status) VALUES(?, ?, ?, ?, ?)`
		_, err := db.Exec(sqlstmt, temp[i].ID, temp[i].Email, temp[i].Password, temp[i].Role, temp[i].Status)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
