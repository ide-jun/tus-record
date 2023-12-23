package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SignIn struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/sign-in", createUser)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	// w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
	// w.Header().Set("Content-Type", "application/json")

	var signIn SignIn
	if err := json.NewDecoder(r.Body).Decode(&signIn); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", signIn)
}
