package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
	wg  sync.WaitGroup
)

func main() {
	db, err = sql.Open("mysql", "root:#Ashfaque18@tcp(127.0.0.1:3306)/kloudone")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		rows, err := db.Query("SELECT * FROM employee")
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		var id int
		var name string
		var department string
		for rows.Next() {
			err := rows.Scan(&id, &name, &department)
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprintf(w, "ID: %d, Name: %s, Depaetment: %s\n", id, name, department)
		}
	}()
	wg.Wait()
}
