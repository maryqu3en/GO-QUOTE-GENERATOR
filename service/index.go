package main

import (
	"encoding/json" //convert struct to json format
	"fmt"
	"math/rand"
	"net/http"
	"os"      //read file
	"strconv" //convert string to int
)

type Quote struct {
	ID     int    `json:"id"`
	Book   string `json:"book"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

var quotes []Quote

func loadQuotes() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &quotes) //similar to JSON.parse in JS
	if err != nil {
		panic(err)
	}
}

func enableCors(w *http.ResponseWriter) { //enable CORS for API access from different domains
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(quotes)
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	randomQuote := quotes[rand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(randomQuote)
}
func getQuoteByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	id, _ := strconv.Atoi(r.URL.Path[len("/api/quotes/"):])
	for _, quote := range quotes {
		if quote.ID == id {
			json.NewEncoder(w).Encode(quote)
			return
		}
	}
	http.Error(w, "Quote not found", http.StatusNotFound)
}

func main() {
	loadQuotes()

	http.HandleFunc("/api/quotes", getQuotes)
	http.HandleFunc("/api/quotes/random", getRandomQuote)
	http.HandleFunc("/api/quotes/", getQuoteByID)

	fmt.Println("Server running on http://localhost:3030")
	http.ListenAndServe(":3030", nil)  // nil = null value for pointers
}
