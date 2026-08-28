package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc is analogous to Express's app.get('/', (req, res) => ...)
	// Notice we pass a function with two parameters:
	// 1. w http.ResponseWriter: The object we use to write data back to the client (like 'res').
	// 2. r *http.Request: A pointer to the incoming request object (like 'req').

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Welcome to GoShort Api")
	})

	fmt.Println("The server is starting on Port 8080")

	// http.ListenAndServe is analogous to app.listen(8080)
	// In Go, functions often return an 'error' type as their last return value.
	// We handle errors explicitly instead of relying on try/catch blocks.
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Server failed: %v",err)
	}
}