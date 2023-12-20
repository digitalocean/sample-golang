package main

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

	if err == nil {

		return nil
	}

	return err
}

// Generic function to execute API calls
func executePreOncallAPIRequest(ctx context.Context, client *http.Client, method, url, token string, requestBody interface{}, responseStruct interface{}) error {

	// Marshal the request body into JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return preOncallAPIError(ctx, err)
	}

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

	client := &http.Client{}

	// Define the request body
	requestBody := MetaInfoApiRequest{
		Business: true,
		Stack:    true,
		Region:   true,
	}

	// Define the response struct
	var responseStruct MetaInfoApiResponse

	// Call the generic function
	url := preOncallPrefix + "/openapi/ticket/v1/getMetaInfo"

	if err := executePreOncallAPIRequest(ctx, client, "POST", url, preOncallToken, requestBody, &responseStruct); err != nil {
		return responseStruct, preOncallAPIError(ctx, err)
	}

	return responseStruct, nil
}

func SubmitPreOncallTicket(ctx context.Context, ticketRequest TicketSubmitRequest) (TicketSubmitResponse, error) {

	client := &http.Client{}

	// Define the response struct
	var responseStruct TicketSubmitResponse

	url := preOncallPrefix + "/openapi/ticket/v1/createTicket"

	if err := executePreOncallAPIRequest(ctx, client, "POST", url, preOncallToken, ticketRequest, &responseStruct); err != nil {
		return responseStruct, preOncallAPIError(ctx, err)
	}

	return responseStruct, nil
}

func GetPreOncallTicket(ctx context.Context, bizTicketID string, channelType string) (TickeInfotResponse, error) {

	client := &http.Client{}

	// Define the request body
	// Construct the URL with query parameters
	url := fmt.Sprintf("%s/openapi/ticket/v1/getTicketsByChannelType?channelType=%s&bizTicketId=%s", preOncallPrefix, channelType, bizTicketID)

	// Call the generic executeAPIRequest function
	var response TickeInfotResponse
	err := executePreOncallAPIRequest(ctx, client, "GET", url, preOncallToken, nil, &response)
	if err != nil {
		return response, err
	}

	return response, nil

}
