package preoncall_service

import (
	"context"
	"encoding/json"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

// Intercom Canvas Receiver
type IntercomCanvasReceiver struct {
	Content IntercomContent `json:"content"`
}

type IntercomContent struct {
	Components []IntercomComponent `json:"components"`
}

type IntercomComponent struct {
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Options []Option `json:"options"`
	Value   *string  `json:"value,omitempty"`
}

func HandlePreoncallCanvasSubmitAction(ctx context.Context, body string) (CanvasReponse, error) {
	//log. := utils.Get//log.gerWithMethod(ctx, "HandlePreoncallCanvasSubmitAction")
	//log..Infof("HandlePreoncallCanvasSubmitAction request body: %v", body)
	var canvasReq IntercomCanvasRequest
	if err := json.Unmarshal([]byte(body), &canvasReq); err != nil {
		fmt.Println("meet parse err", err)

		return CanvasReponse{}, err
	}

	intercomConversationID := canvasReq.Conversation.ConversationID
	assigneeID := canvasReq.Conversation.AdminAssigneeID
	inputValues := canvasReq.InputValues

	//log..Infof("HandlePreoncallCanvasSubmitAction pretty's request %v \n", larkcore.Prettify(canvasReq))
	fmt.Printf("HandlePreoncallCanvasSubmitAction pretty's request %v \n", larkcore.Prettify(canvasReq))
	var response CanvasReponse
	// TODO: call intercom pre ocall api to create ticket
	//log..Infof("HandlePreoncallCanvasSubmitAction component id %v", canvasReq.ComponentID)
	fmt.Printf("HandlePreoncallCanvasSubmitAction component id %v", canvasReq.ComponentID)
	switch canvasReq.ComponentID {
	case CategorySingleSelectID:
		//log..Infof("HandlePreoncallCanvasSubmitAction single select ")
		if value, ok := canvasReq.InputValues[CategorySingleSelectID]; ok {
			//log..Infof("HandlePreoncallCanvasSubmitAction single select value %v", value)
			fmt.Printf("HandlePreoncallCanvasSubmitAction single select value %v \n", value)
			if value == CreateTicketOptionID || value == BizLineSearchButtonID || value == RegionSearchButtonID {
				response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, value, canvasReq.CurrentCanvas)
			} else if value == RelatedTicketOptionID {
				response = GetRelatedTicketCanvasBody(ctx, intercomConversationID)
			} else if value == SubmitTicketButtonID {

			}
		}
	}

	//log..Infof("HandlePreoncallCanvasSubmitAction vanvas response %v", larkcore.Prettify(response))

	return response, nil
}

func HandlePreoncallInitializationAction(ctx context.Context) CanvasReponse {
	//log. := utils.Get//log.gerWithMethod(ctx, "HandlePreoncallInitializationAction")
	response := GetInitTicketCanvasBody()

	//log..Infof("HandlePreoncallInitializationAction response %v", larkcore.Prettify(response))

	return response
}
