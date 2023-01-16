package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:#Ashfaque18@tcp(127.0.0.1:3306)/kloudone")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO department (frontend, backend) VALUES (?, ?)", "Ashfaque", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted successfully!")
}
