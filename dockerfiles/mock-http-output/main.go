package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----------------------------")

	var decodedBody []byte

	// Check if the request's Content-Encoding header indicates Gzip encoding
	if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
		// Create a Gzip reader to decode the request body
		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, "Failed to create Gzip reader", http.StatusBadRequest)
			return
		}
		defer reader.Close()

		// Read the decoded request body
		var body = new(bytes.Buffer)
		_, err = io.Copy(body, reader)
		if err != nil {
			http.Error(w, "Failed to read Gzip body", http.StatusInternalServerError)
			return
		}

		// Convert the body to a string and print it
		decodedBody = body.Bytes()
	} else {
		// Read the request body
		var err error
		decodedBody, err = ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
	}

	// Print the request body
	fmt.Println(string(decodedBody))
}

func main() {
	// Define the endpoint and request handler function
	http.HandleFunc("/ingest", requestHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
