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
	defer db.Close()

	// ID       string
	// Email    string
	// Password string
	// Role     string
	// Status   string
	// sqltable := fmt.Sprintf("CREATE TABLE users (id CHARACTER (32) NOT NULL PRIMARY KEY, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, role VARCHAR(255) NOT NULL, status VARCHAR(255) NOT NULL)")
	// db.Exec(sqltable)
	// temp := []domain.Customer{
	// 	{ID: "1000", Email: "hadn4@fpt.com.vn", Password: "Alochym@123", Role: "Husband", Status: "1"},
	// 	{ID: "1001", Email: "1_nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1002", Email: "2_nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1003", Email: "3_nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1004", Email: "4_nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1005", Email: "5_nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1006", Email: "6nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1007", Email: "7nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1008", Email: "n8huntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1009", Email: "nh9untt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1010", Email: "nhu10ntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1011", Email: "nhunt11t@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1012", Email: "nhuntt12@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1013", Email: "13nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1014", Email: "nh14untt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1015", Email: "nhun15tt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1016", Email: "nhuntt16@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1017", Email: "17nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1018", Email: "nh18untt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1019", Email: "nhun19tt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1020", Email: "nhuntt20@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1021", Email: "21nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1022", Email: "nh22untt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1023", Email: "nhun23tt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1024", Email: "nhuntt24@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// 	{ID: "1025", Email: "25nhuntt@fpt.com.vn", Password: "Alochym@123", Role: "Wife", Status: "1"},
	// }

	// for i := range temp {
	// 	sqlstmt := "INSERT INTO users (id, email, password, role, status) VALUES(?, ?, ?, ?, ?)"
	// 	result, err := db.Exec(sqlstmt, temp[i].ID, temp[i].Email, temp[i].Password, temp[i].Role, temp[i].Status)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	id, err := result.LastInsertId()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(id)
	// }

	router.Start(db)
}
