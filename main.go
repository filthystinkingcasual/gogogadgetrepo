package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func getClientIP(r *http.Request) (string, error) {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		ip := strings.TrimSpace(ips[0])
		return ip, nil

	}
	realip := r.Header.Get("X-Real-IP")
	if realip != "" {
		return realip, nil
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", fmt.Errorf("failed to parse remote address: %v", err)
	}
	return ip, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip, err := getClientIP(r)
	if err != nil {
		log.Printf("Error getting IP: %v\n", err)
		http.Error(w, "Could not determine client IP", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Hello, your IP is %s", ip)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
