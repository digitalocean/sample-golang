package main

import (
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"

	"io/ioutil"
	"net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("-------------------------------------- this is submit request")
	// Convert the bytes to string and print it
	bodyString := string(bodyBytes)
	fmt.Println("this iis the body this %v", bodyString)
	fmt.Println("======================================")
	// You must close the original body
	defer r.Body.Close()
	// Unmarshal the JSON into your struct
	var initReq InitializeRequest
	if err := json.Unmarshal(bodyBytes, &initReq); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}
	fmt.Println("pretty's request %v", larkcore.Prettify(initReq))

	response := GetInitTicketCanvasBody()

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	fmt.Println("response %v", response)

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
