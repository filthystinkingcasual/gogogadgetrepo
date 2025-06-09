package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "things happened. yay.")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8080 for reasons")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Something exploded:", err)
	}
}
