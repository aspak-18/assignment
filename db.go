package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create the "employee" table
	employeeTable := `CREATE TABLE employee (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        department VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`
	_, err = db.Exec(employeeTable)
	if err != nil {
		panic(err)
	}

	// Create the "department" table
	departmentTable := `CREATE TABLE department (
        id INT NOT NULL AUTO_INCREMENT,
        frontend VARCHAR(255) NOT NULL,
        backend VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`
	_, err = db.Exec(departmentTable)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tables created successfully!")
}
