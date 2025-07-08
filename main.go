package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	repoUrl := viper.GetString("gitrepo")
	u, _ := url.Parse(repoUrl)
	repoName := strings.TrimLeft(u.Path, "/")
	data := struct {
		AppName string
		GitRepo string
		GitURL  string
	}{
		AppName: viper.GetString("appName"),
		GitURL:  viper.GetString("gitrepo"),
		GitRepo: repoName,
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func config(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/config.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	settings := viper.AllSettings()
	t.Execute(w, settings)
}

func main() {
	http.HandleFunc("/", home)
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
