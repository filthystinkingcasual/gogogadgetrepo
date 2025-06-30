package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("something is wrong: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Name          string
			FavoriteThing string
		}{
			Name:          "Mara",
			FavoriteThing: "makin bacon pancakes",
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "you oopsed: "+err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Listening on :8080, like a real server...")
	http.ListenAndServe(":8080", nil)
}
