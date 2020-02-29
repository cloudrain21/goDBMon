package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	var c1 int
	var c2 string
	rows, err := db.Query("SELECT c1, c2 FROM t1 where c1 >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&c1, &c2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c1, c2)
	}
	defer db.Close()
}
