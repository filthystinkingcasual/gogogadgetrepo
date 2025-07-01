package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/spf13/viper"
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

func config(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/config.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	hostname := viper.GetString("hostname")
	t.Execute(w, hostname)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/config", config)
	port := viper.GetString("port")
	log.Printf("Starting sever on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Something exploded:", err)
	}
}

func init() {
	viper.SetConfigName("gogogadget")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		log.Println(viper.ConfigFileUsed())
	}
}
