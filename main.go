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

	insertData(db)
	getRows(db)
	preparedStatement(db)
}

func getRows(db *sql.DB) {
	fmt.Println("-----Get rows from database")
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

func preparedStatement(db *sql.DB) {
	fmt.Println("-----Select prepared statement")
	stmt, err := db.Prepare("SELECT name FROM bank.clients WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	results, err := stmt.Query(1)
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

func insertData(db *sql.DB) {
	fmt.Println("-----Insert prepared statement")
	stmt, err := db.Prepare("INSERT INTO bank.clients(department,name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec("IT", "john")
	if err != nil {
		log.Fatal(err)
	}
}
