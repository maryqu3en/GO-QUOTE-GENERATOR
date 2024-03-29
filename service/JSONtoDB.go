// MOVE DATA FROM JSON FILE TO MYSQL DB TABLE

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"log"
    _ "github.com/go-sql-driver/mysql"  // run: go get github.com/go-sql-driver/mysql
)

type Quote struct {
    ID     int    `json:"id"`
    Book   string `json:"book"`
    Author string `json:"author"`
    Quote  string `json:"quote"`
}

var quotes []Quote
var db *sql.DB

func loadQuotes() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &quotes)
	if err != nil {
		panic(err)
	}
}

func insertToDB() {
	for _, quote := range quotes {
		if len(quote.Quote) > 255 {
			quote.Quote = quote.Quote[:255]
		}
		_, err := db.Exec("INSERT IGNORE INTO Quote (book, author, quote) VALUES (?, ?, ?)",
			quote.Book, quote.Author, quote.Quote)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	var err error
    db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    loadQuotes()
    insertToDB()

    fmt.Println("Data inserted successfully.")
}
