package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()
	statusCode, err := strconv.Atoi(queryParams.Get("status"))
	if err != nil {
		statusCode = http.StatusOK // Default status code is 200 OK
	}

	responseTime, err := strconv.Atoi(queryParams.Get("response_time"))
	if err != nil {
		responseTime = 0 // Default response time is 0 seconds
	}

	bodyBase64 := queryParams.Get("body")
	var body []byte
	if bodyBase64 != "" {
		decodedBody, err := base64.StdEncoding.DecodeString(bodyBase64)
		if err != nil {
			http.Error(w, "Failed to decode body", http.StatusBadRequest)
			return
		}
		body = decodedBody
	}

	// Wait for the specified response time
	time.Sleep(time.Duration(responseTime) * time.Second)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Set status code
	w.WriteHeader(statusCode)

	// Set body if provided
	if len(body) > 0 {
		w.Write(body)
	}

	logMessage := fmt.Sprintf("Path: %s, Método HTTP: %s, Status Code: %d", r.URL.Path, r.Method, statusCode)
	log.Println(logMessage)
}

var (
	project = "http-timeout-error-handler"
	version = "1.0.0-SNAPSHOT"
)

func main() {
	fmt.Println("Project:", project)
	fmt.Println("Version:", version)
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8000", nil)
}
