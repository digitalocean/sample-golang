package main

import (
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"log"
)

func HandlePreoncallCanvasSubmitAction(body string) (CanvasReponse, error) {
	var initReq Workspace
	if err := json.Unmarshal([]byte(body), &initReq); err != nil {
		fmt.Println("meet parse err", err)

		return CanvasReponse{}, err
	}
	fmt.Printf("pretty's request %v \n", larkcore.Prettify(initReq))
	fmt.Println("======================================")
	response := GetCreateTicketCanvasBody([]string{"bizline1", "bizline2"}, []string{"region1", "region2"}, []string{"stack1", "stack2"}, initReq.InputValues)
	//fmt.Println("GetCreateTicketCanvasBody response %v", larkcore.Prettify(response))
	// Send the response as JSON
	// Marshal the struct to a JSON byte slice
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	// Convert the byte slice to a string and print it
	jsonString := string(jsonData)
	fmt.Printf("this is the submit json string %v\n", jsonString)

	return response, nil
}

func HandlePreoncallInitializationAction() CanvasReponse {
	// Construct the response object
	response := GetInitTicketCanvasBody()

	return response
}
