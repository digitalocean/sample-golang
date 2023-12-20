package preoncall_service

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCreateTicketCanvasBody(t *testing.T) {
	inputValues := map[string]string{
		"category_single_select": "create_ticket",
		"business_single_select": "bid_urso2fcl",
		"priority_single_select": "P0",
		"stack_single_select":    "pc",
		"region_single_select":   "India",
		"version_single_select":  "7.4.8",
		"app_single_select":      "Lark",
		"external_single_select": "字节内",
		"source_single_select":   "客服渠道",
		"reporter_single_select": "1110",
		"remarks_text":           "this is 1 intercom ticket",
		"channel_type":           "intercom",
		"biz_ticket_id":          "11111",
	}

	intercomConversationID := "12345"
	assigneeID := 12345
	value := "create_ticket"
	canvas := IntercomCanvasReceiver{
		Content: IntercomContent{
			Components: []IntercomComponent{},
		},
	}

	response := GetCreateTicketCanvasBody(context.Background(), inputValues, intercomConversationID, assigneeID, value, canvas)

	fmt.Println(response)
	//log..Infof("response: %v", response)
}
