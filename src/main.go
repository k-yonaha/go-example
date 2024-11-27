package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker and Compose!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}