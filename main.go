package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"strings"
	"text/template"

	_ "embed"

	"github.com/spf13/viper"
)

//go:embed views/*
var content embed.FS

func config(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/config.html")
	if err != nil {
		http.Error(w, "lol no", http.StatusInternalServerError)
		return
	}
	settings := viper.AllSettings()
	t.Execute(w, settings)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html")
	if err != nil {
		http.Error(w, "that didn't work out so well for you did it", http.StatusInternalServerError)
		return
	}
	repoUrl := viper.GetString("gitrepo")
	u, _ := url.Parse(repoUrl)
	repoName := strings.TrimLeft(u.Path, "/")
	data := struct {
		AppName string
		GitRepo string
		GitURL  string
		Item1   string
		Item2   string
		Item3   string
		Item4   string
	}{
		AppName: viper.GetString("appName"),
		GitURL:  viper.GetString("gitrepo"),
		GitRepo: repoName,
		Item1:   viper.GetString("configthefirst"),
		Item2:   viper.GetString("configthesecond"),
		Item3:   viper.GetString("configthethird"),
		Item4:   viper.GetString("configthefourth"),
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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
	viper.SetDefault("port", "8081")
	viper.SetConfigName("gogogadget")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err == nil {
		log.Println(viper.ConfigFileUsed())
	}
}
