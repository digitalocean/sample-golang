package pre_oncall

import (
	"context"
	"time"
)

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

func GetFakePreOncallTicket(ctx context.Context, bizTicketID string, channelType string) (TickeInfotResponse, error) {
	// Mock data for the response
	data := []Ticket{
		{
			TicketId:     "TICKET001",
			Title:        "Server Downtime",
			Business:     "IT",
			BusinessName: "IT Services",
			Assignee:     []string{"John Doe"},
			Status:       "Open",
			Reporter:     "Alice Johnson",
			OpenChatId:   "CHAT001",
			Remarks:      "Server is down since 3 AM",
			CreatedAt:    time.Now().Add(-48 * time.Hour).Format(time.RFC3339), // 2 days ago
			UpdatedAt:    time.Now().Format(time.RFC3339),
			GroupLink:    "http://example.com/group/001",
			TicketLink:   "http://example.com/ticket/TICKET001",
			ChannelType:  "Email",
			BizTicketId:  "BIZ001",
		},
		{
			TicketId:     "TICKET002",
			Title:        "Database Connectivity Issue",
			Business:     "Database Team",
			BusinessName: "DB Services",
			Assignee:     []string{"Jane Smith", "Emily White"},
			Status:       "In Progress",
			Reporter:     "Bob Brown",
			OpenChatId:   "CHAT002",
			Remarks:      "Intermittent connectivity issues observed",
			CreatedAt:    time.Now().Add(-24 * time.Hour).Format(time.RFC3339), // 1 day ago
			UpdatedAt:    time.Now().Format(time.RFC3339),
			GroupLink:    "http://example.com/group/002",
			TicketLink:   "http://example.com/ticket/TICKET002",
			ChannelType:  "Slack",
			BizTicketId:  "BIZ002",
		},
		{
			TicketId:     "TICKET003",
			Title:        "Email Service Disruption",
			Business:     "Communications",
			BusinessName: "Communication Services",
			Assignee:     []string{"Michael Green"},
			Status:       "Resolved",
			Reporter:     "Clara Davis",
			OpenChatId:   "CHAT003",
			Remarks:      "Emails were not being sent out",
			CreatedAt:    time.Now().Add(-72 * time.Hour).Format(time.RFC3339), // 3 days ago
			UpdatedAt:    time.Now().Format(time.RFC3339),
			GroupLink:    "http://example.com/group/003",
			TicketLink:   "http://example.com/ticket/TICKET003",
			ChannelType:  "Phone",
			BizTicketId:  "BIZ003",
		},
	}

	return TickeInfotResponse{
		Code: 200,
		Msg:  "Success",
		Data: data,
	}, nil
}
