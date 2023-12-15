package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type fruit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var fruits = []fruit{
	{ID: 1, Name: "appleだよ", Icon: "🍎"},
	{ID: 2, Name: "bananaよ", Icon: "🍌"},
	{ID: 3, Name: "grapeだよ", Icon: "🍇"},
}

func main() {
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/fruit", getFruits)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func getFruits(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	json.NewEncoder(w).Encode(fruits)
}
