package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Convert the bytes to string and print it
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// You must close the original body
	defer r.Body.Close()

}

func InitializeCanvasHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body of the POST request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Log the body, can remove this in production
	fmt.Println("Received initialize request with body:", string(body))

	// Construct the response object
	response := GetInitTicketCanvasBody()

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
