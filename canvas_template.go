package main

// Define the structures based on the JSON input
const (
	RelatedTicketID = "related-ticket"
	CreateTicketID  = "create-ticket"
	SubmitTicketID  = "submit-ticket"
)

type Action struct {
	Type string `json:"type,omitempty"`
}

type Option struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Components struct {
	Type    string   `json:"type"`
	ID      string   `json:"id,omitempty"`
	Label   string   `json:"label,omitempty"`
	Text    string   `json:"text,omitempty"`
	Style   string   `json:"style,omitempty"`
	Size    string   `json:"size,omitempty"`
	Action  Action   `json:"action,omitempty"`
	Options []Option `json:"options,omitempty"`
}

type Content struct {
	Components []Components `json:"components"`
}

type Canvas struct {
	Content Content `json:"content"`
}

type CanvasResponse struct {
	Canvas Canvas `json:"canvas"`
}

var InitTicketCanvasBody = []Components{
	{
		Type:  "single-select",
		ID:    "ticket-option",
		Label: "Pre-Oncall Ticket",
		Action: Action{
			Type: "submit",
		},
		Options: []Option{
			{Type: "option", ID: RelatedTicketID, Text: "Related Ticket"},
			{Type: "option", ID: SubmitTicketID, Text: "Create Ticket"},
		},
	},
}

func GetInitTicketCanvasBody() CanvasResponse {
	return CanvasResponse{
		Canvas: Canvas{
			Content: Content{
				Components: InitTicketCanvasBody,
			},
		},
	}
}
