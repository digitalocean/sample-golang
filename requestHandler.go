package main

import (
	"context"
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
	fmt.Println("-------------------------------------- this is submit request")
	defer r.Body.Close()
	// Convert the bytes to string and print it
	bodyString := string(bodyBytes)
	fmt.Printf("this iis the body this %v \n", bodyString)
	fmt.Println("======================================")
	// You must close the original body

	// Unmarshal the JSON into your struct
	response, err := HandlePreoncallCanvasSubmitAction(context.Background(), bodyString)
	if err != nil {
		////log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
		fmt.Println("Error occurred during marshaling. Error: %s", err.Error()
		return
	}

	canvas, err := json.Marshal(response.Canvas)
	if err != nil {
		fmt.Println("Error marshalling response %V", http.StatusInternalServerError)
		return
	}

	fmt.Printf("+ ++++ this is the response of submit handler %v", string(canvas))

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling response %V", http.StatusInternalServerError)
		return
	}
	fmt.Printf("this is the response of submit handler %v", string(jsonResponse))

	w.Write(jsonResponse)

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
	response := HandlePreoncallInitializationAction(context.Background())
	fmt.Printf("response %v\n", response)

	// Convert the byte slice to a string and print it

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling response %V", http.StatusInternalServerError)
		return
	}
	fmt.Printf("this is the response of initialize handler %v \n", string(jsonResponse))

	w.Write(jsonResponse)
}
