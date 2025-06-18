package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "It's broken. Just like my spirit.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8080 for reasons")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Something exploded:", err)
	}
}
