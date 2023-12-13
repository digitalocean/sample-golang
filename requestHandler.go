package main

import (
	"encoding/json"
	"fmt"
	"log"

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
	defer r.Body.Close()
	// Convert the bytes to string and print it
	bodyString := string(bodyBytes)
	fmt.Println("this iis the body this %v", bodyString)
	fmt.Println("======================================")
	// You must close the original body

	// Unmarshal the JSON into your struct
	response, err := HandlePreoncallCanvasSubmitAction(bodyString)
	if err != nil {
		//log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

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
	response := HandlePreoncallInitializationAction()
	fmt.Println("response %v", response)

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	// Convert the byte slice to a string and print it
	jsonString := string(jsonData)
	fmt.Println("this is the init json string %v", jsonString)

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
