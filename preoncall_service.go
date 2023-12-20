package main

import (
	"context"
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

func HandlePreoncallCanvasSubmitAction(ctx context.Context, body string) (CanvasReponse, error) {
	////log. utils.GetLoggerWithMethod(ctx, "HandlePreoncallCanvasSubmitAction")
	////log.Infof("HandlePreoncallCanvasSubmitAction request body: %v", body)
	var canvasReq IntercomCanvasRequest
	if err := json.Unmarshal([]byte(body), &canvasReq); err != nil {
		fmt.Println("meet parse err", err)

		return CanvasReponse{}, err
	}

	intercomConversationID := canvasReq.Conversation.ConversationID
	assigneeID := canvasReq.Conversation.AdminAssigneeID
	inputValues := canvasReq.InputValues

	////log.Infof("HandlePreoncallCanvasSubmitAction pretty's request %v \n", larkcore.Prettify(canvasReq))
	var response CanvasReponse
	// TODO: call intercom pre ocall api to create ticket
	////log.Infof("HandlePreoncallCanvasSubmitAction component id %v", canvasReq.ComponentID)
	switch canvasReq.ComponentID {
	case CategorySingleSelectID:
		////log.Infof("HandlePreoncallCanvasSubmitAction single select ")
		if value, ok := canvasReq.InputValues[CategorySingleSelectID]; ok {
			////log.Infof("HandlePreoncallCanvasSubmitAction single select value %v", value)
			if value == CreateTicketOptionID || value == BizLineSearchButtonID || value == RegionSearchButtonID {
				response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, value, canvasReq.CurrentCanvas)
				fmt.Printf("====== response of  GetCreateTicketCanvasBody%v", larkcore.Prettify(response))
			}
		}
	}

	////log.Infof("HandlePreoncallCanvasSubmitAction vanvas response %v", larkcore.Prettify(response))
	fmt.Printf(" ------response of  GetCreateTicketCanvasBody%v", larkcore.Prettify(response))
	return response, nil
}

func HandlePreoncallInitializationAction(ctx context.Context) CanvasReponse {
	////log. utils.GetLoggerWithMethod(ctx, "HandlePreoncallInitializationAction")
	response := GetInitTicketCanvasBody()

	////log.Infof("HandlePreoncallInitializationAction response %v", larkcore.Prettify(response))

	return response
}
