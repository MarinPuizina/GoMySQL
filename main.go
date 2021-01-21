package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	Name string
}

func main() {
	fmt.Println("-----Trying to connect to the MySQL databse")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bank")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("-----Successfully connected to the MySQL databse")

	getRows(db)
}

func getRows(db *sql.DB) {
	results, err := db.Query("SELECT name FROM bank.clients")
	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var client Client

		err = results.Scan(&client.Name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Client name=" + client.Name)
	}
}
