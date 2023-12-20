package pre_oncall

import "context"

func GetFakePreOncallMetaInfo(ctx context.Context, business bool, stack bool) (MetaInfoApiResponse, error) {
	// Create dummy data for the response
	fakeResponse := MetaInfoApiResponse{
		Code: 200,       // Example success code
		Msg:  "Success", // Success message
		Data: Data{
			BusinessList: []Business{
				{
					Bid:      "BID123",
					Enabled:  true,
					Name:     "Dummy Business 1",
					Desc:     "Description for Dummy Business 1",
					ParentId: "PID123",
					Inherit:  true,
					Stacks:   []string{"Stack1", "Stack2"},
				},
				{
					Bid:      "BID456",
					Enabled:  false,
					Name:     "Dummy Business 2",
					Desc:     "Description for Dummy Business 2",
					ParentId: "PID456",
					Inherit:  false,
					Stacks:   []string{"Stack3", "Stack4"},
				},
				{
					Bid:      "BID789",
					Enabled:  true,
					Name:     "Dummy Business 3",
					Desc:     "Description for Dummy Business 3",
					ParentId: "PID789",
					Inherit:  true,
					Stacks:   []string{"Stack5", "Stack6"},
				},
			},
			RegionList: []CodeNamePair{
				{
					Code: "R1",
					Name: "Region 1",
				},
				{
					Code: "R2",
					Name: "Region 2",
				},
				{
					Code: "R3",
					Name: "Region 3",
				},
				{
					Code: "R4",
					Name: "Region 4",
				},
			},
		},
	}
	return fakeResponse, nil
}

func SubmitFakePreOncallTicket(ctx context.Context, ticketRequest TicketSubmitRequest) (TicketSubmitResponse, error) {
	// Create dummy data for the TicketSubmitResponse
	fakeResponse := TicketSubmitResponse{
		Code: 200,                            // Example success code
		Msg:  "Ticket submission successful", // Success message
		Data: struct {
			TicketId string `json:"ticketId"`
		}{
			TicketId: "TICKET12345", // Dummy Ticket ID
		},
	}

	// You can //log. the fake response if needed
	// //log..Infof("submit_fake_pre_oncall_ticket response: %v", larkcore.Prettify(fakeResponse))

	// Return the fake response with no error
	return fakeResponse, nil
}
