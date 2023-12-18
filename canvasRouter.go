package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func HandlePreoncallCanvasSubmitAction(body string) (CanvasReponse, error) {
	var canvasReq IntercomCanvasRequest
	if err := json.Unmarshal([]byte(body), &canvasReq); err != nil {
		fmt.Println("meet parse err", err)

		return CanvasReponse{}, err
	}
	ctx := context.Background()
	var response CanvasReponse
	// TODO: call intercom pre ocall api to create ticket
	switch canvasReq.ComponentID {
	case CategorySingleSelectID:
		fmt.Printf("HandlePreoncallCanvasSubmitAction single select \n")
		if value, ok := canvasReq.InputValues[CategorySingleSelectID]; ok {
			fmt.Printf("HandlePreoncallCanvasSubmitAction single select value %v \n", value)
			if value == SummitTicketOptionID {
				response = GetCreateTicketCanvasBody(ctx, canvasReq.InputValues)
			} else if value == RelatedTicketOptionID {

			}
		}
	}

	fmt.Printf("HandlePreoncallCanvasSubmitAction vanvas response %v \n", response)

	return response, nil
}

func HandlePreoncallInitializationAction() CanvasReponse {
	// Construct the response object
	response := GetInitTicketCanvasBody()

	return response
}
