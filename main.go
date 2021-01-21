package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bank")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to the MySQL databse")
}
