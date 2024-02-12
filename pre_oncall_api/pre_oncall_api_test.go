package pre_oncall

import (
	"context"
	"fmt"
	"testing"
)

// Test function
func TestCallMetaInfoAPI(t *testing.T) {

	resp, err := GetPreOncallMetaInfo(context.Background(), true, true)
	if err != nil {
		t.Errorf("call_metainfo_api() returned an error: %v", err)
	}

	fmt.Println(resp)
}

func TestPreOncallTicketCreation(t *testing.T) {
	ticketCreationReq := TicketSubmitRequest{
		Title:         "A Test Junyu",
		Business:      "bid_urso2fcl",
		Priority:      "P0",
		Stack:         "pc",
		Region:        "India",
		UserId:        "12344321",
		Version:       "7.4.8",
		CreateChatWay: "autoCreate",
		Type:          "Pre-Oncall",
		App:           "Lark",
		External:      "字节内",
		Source:        "客服渠道",
		Reporter:      "1110",
		Remarks:       "this is 1 intercom ticket",
		ChannelType:   "intercom",
		BizTicketId:   "11111",
	}

	resp, err := SubmitPreOncallTicket(context.Background(), ticketCreationReq)
	if err != nil {
		t.Errorf("call_metainfo_api() returned an error: %v", err)
	}

	// ticket_17029332969470773
	fmt.Println(resp)
}

func TestGetPreOncallMetaInfo(t *testing.T) {
	resp, err := GetPreOncallTicket(context.Background(), "397", "intercom")
	if err != nil {
		t.Errorf("call_metainfo_api() returned an error: %v", err)
	}

	fmt.Println(resp)
}
