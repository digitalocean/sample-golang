package pre_oncall

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	preOncallToken  = "26ad213fcdc54e0da3a6e7fc79e99b75"
	preOncallPrefix = "https://lark-oncall.bytedance.net"
)

func preOncallAPIError(ctx context.Context, err error) error {
	////log. := utils.Get//log.gerWithMethod(ctx, "preOncallError")
	if err == nil {
		////log..Infof("preOncallError: %v", errors.New("unknown error"))
		return nil
	}
	////log..Errorf("preOncallError: %v", err.Error())

	return err
}

// Generic function to execute API calls
func executePreOncallAPIRequest(ctx context.Context, client *http.Client, method, url, token string, requestBody interface{}, responseStruct interface{}) error {
	////log. := utils.Get//log.gerWithMethod(ctx, "executePreOncallAPIRequest")
	// Marshal the request body into JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return preOncallAPIError(ctx, err)
	}

	////log..Infof("executePreOncallAPIRequest request body: %v", larkcore.Prettify(jsonBody))

	// Create a new request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return preOncallAPIError(ctx, err)
	}

	// Add headers
	req.Header.Set("Authorization", "Basic "+token)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return preOncallAPIError(ctx, err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return preOncallAPIError(ctx, err)
	}

	// Unmarshal the response into the provided response struct
	if err := json.Unmarshal(respBody, &responseStruct); err != nil {
		return preOncallAPIError(ctx, err)
	}

	return nil
}

func GetPreOncallMetaInfo(ctx context.Context, business bool, stack bool) (MetaInfoApiResponse, error) {
	////log. := utils.Get//log.gerWithMethod(ctx, "get_pre_oncall_meta_info")

	client := &http.Client{}

	// Define the request body
	requestBody := MetaInfoApiRequest{
		Business: true,
		Stack:    true,
		Region:   true,
	}

	////log..Infof("get_pre_oncall_meta_info request body: %v", larkcore.Prettify(requestBody))

	// Define the response struct
	var responseStruct MetaInfoApiResponse

	// Call the generic function
	url := preOncallPrefix + "/openapi/ticket/v1/getMetaInfo"

	if err := executePreOncallAPIRequest(ctx, client, "POST", url, preOncallToken, requestBody, &responseStruct); err != nil {
		return responseStruct, preOncallAPIError(ctx, err)
	}

	////log..Infof("get_pre_oncall_meta_info response code: %v, msg: %v", responseStruct.Code, responseStruct.Msg)
	////log..Infof("get_pre_oncall_meta_info response body: %v", larkcore.Prettify(responseStruct))

	return responseStruct, nil
}

func SubmitPreOncallTicket(ctx context.Context, ticketRequest TicketSubmitRequest) (TicketSubmitResponse, error) {
	////log. := utils.Get//log.gerWithMethod(ctx, "submit_pre_oncall_ticket")

	client := &http.Client{}

	//log..Infof("submit_pre_oncall_ticket request body: %v", larkcore.Prettify(ticketRequest))

	// Define the response struct
	var responseStruct TicketSubmitResponse

	url := preOncallPrefix + "/openapi/ticket/v1/createTicket"

	if err := executePreOncallAPIRequest(ctx, client, "POST", url, preOncallToken, ticketRequest, &responseStruct); err != nil {
		return responseStruct, preOncallAPIError(ctx, err)
	}

	//log..Infof("submit_pre_oncall_ticket response code: %v, msg: %v", responseStruct.Code, responseStruct.Msg)
	//log..Infof("submit_pre_oncall_ticket response body: %v", larkcore.Prettify(responseStruct))

	return responseStruct, nil
}

func GetPreOncallTicket(ctx context.Context, bizTicketID string, channelType string) (TickeInfotResponse, error) {
	//log. := utils.Get//log.gerWithMethod(ctx, "get_pre_oncall_ticket")

	client := &http.Client{}

	// Define the request body
	// Construct the URL with query parameters
	url := fmt.Sprintf("%s/openapi/ticket/v1/getTicketsByChannelType?channelType=%s&bizTicketId=%s", preOncallPrefix, channelType, bizTicketID)
	//log..Infof("get_pre_oncall_ticket request url: %v", url)
	// Call the generic executeAPIRequest function
	var response TickeInfotResponse
	err := executePreOncallAPIRequest(ctx, client, "GET", url, preOncallToken, nil, &response)
	if err != nil {
		return response, err
	}

	return response, nil

}

func GetFakePreOncallTicket(ctx context.Context, bizTicketID string, channelType string) (TickeInfotResponse, error) {
	// Create dummy data for the TickeInfotResponse
	fakeResponse := TickeInfotResponse{
		Code: 200,       // Example success code
		Msg:  "Success", // Success message
		Data: []Ticket{
			{
				TicketId:     "TICKET101",
				Title:        "Issue with Service A",
				Business:     "BUS123",
				BusinessName: "Business A",
				Assignee:     []string{"User1", "User2"},
				Status:       "Open",
				Reporter:     "ReporterA",
				OpenChatId:   "CHATID101",
				Remarks:      "Urgent issue with high priority",
				CreatedAt:    "2023-12-18T15:04:05Z",
				UpdatedAt:    "2023-12-18T16:00:00Z",
				GroupLink:    "http://example.com/group/101",
				TicketLink:   "http://example.com/ticket/101",
				ChannelType:  channelType,
				BizTicketId:  bizTicketID,
			},
			{
				TicketId:     "TICKET102",
				Title:        "Request for Configuration Update",
				Business:     "BUS456",
				BusinessName: "Business B",
				Assignee:     []string{"User3"},
				Status:       "Resolved",
				Reporter:     "ReporterB",
				OpenChatId:   "CHATID102",
				Remarks:      "Configuration updated successfully",
				CreatedAt:    "2023-12-17T10:20:30Z",
				UpdatedAt:    "2023-12-17T11:15:45Z",
				GroupLink:    "http://example.com/group/102",
				TicketLink:   "http://example.com/ticket/102",
				ChannelType:  "ServiceRequest",
				BizTicketId:  "REQ456",
			},
		},
	}
	// //log..Infof("get_fake_pre_oncall_ticket response: %v", larkcore.Prettify(fakeResponse))

	// Return the fake response with no error
	return fakeResponse, nil
}
