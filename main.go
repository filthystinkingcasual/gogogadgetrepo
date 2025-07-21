package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// structs
type GeoResponse struct {
	Status  string  `json:"status"`
	Country string  `json:"country"`
	Region  string  `json:"region"`
	City    string  `json:"city"`
	Zip     string  `json:"zip"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Query   string  `json:"query"`
}

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

func getGeoData(ip string) (*GeoResponse, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get geo data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var geoData GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return nil, fmt.Errorf("failed to decode geo data: %v", err)
	}

	return &geoData, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip, err := getClientIP(r)
	if err != nil {
		log.Printf("Error getting IP: %v\n", err)
		http.Error(w, "Could not determine client IP", http.StatusInternalServerError)
		return
	}
	//	fmt.Fprintf(w, "Hello, your IP is %s", ip)

	geo, err := getGeoData(ip)
	if err != nil {
		log.Printf("Geo lookup failed: %v", err)
		http.Error(w, "Geo lookup failed", http.StatusInternalServerError)
		return
	}

	log.Printf("Geo data for IP %s, %s %s \n", geo.City, geo.Region, geo.Zip)

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
