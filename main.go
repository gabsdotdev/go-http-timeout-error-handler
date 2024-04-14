package main

import (
	"encoding/base64"
	"fmt"
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
	// Set status code
	w.WriteHeader(statusCode)

	// Set body if provided
	if len(body) > 0 {
		w.Write(body)
	}

	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Método HTTP:", r.Method)
	fmt.Println("Status Code:", statusCode)
	fmt.Println("------------------")
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8000", nil)
}
