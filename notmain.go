// main.go
package main

import (
	"html/template"
	"log"
	"net/http"
)

type FormData struct {
	Name      string
	Message   string
	UserAgent string
	Language  string
}

var formTemplate = template.Must(template.ParseFiles("form.html"))
var resultTemplate = template.Must(template.ParseFiles("result.html"))

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := formTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	data := FormData{
		Name:      r.FormValue("name"),
		Message:   r.FormValue("message"),
		UserAgent: r.FormValue("userAgent"),
		Language:  r.FormValue("langData"),
	}

	if err := resultTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", resultHandler)

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
