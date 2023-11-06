package main

import (
	"encoding/json"
	"fmt"
)

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

type InputValues struct {
	TicketOption string `json:"pre-oncall-ticket-option"`
}

type InitializeRequest struct {
	CurrentCanvas Canvas      `json:"current_canvas"`
	InputValues   InputValues `json:"input_values"`
}

func InitPreOncallCanvas() []Components {
	initTicketCanvasBody := []Components{
		{
			Type:  "single-select",
			ID:    "pre-oncall-ticket-option",
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

	return initTicketCanvasBody
}

func GetCreateTicketComponents() []Components {
	jsonData := `[
		{
			"type": "single-select",
			"id": "ticket-option",
			"label": "Pre-Oncall Ticket",
			"options": [
				{
					"type": "option",
					"id": "option-1",
					"text": "Related Ticket"
				},
				{
					"type": "option",
					"id": "option-2",
					"text": "Create Ticket"
				}
			]
		},
		{
			"type": "text",
			"text": "*Create a ticket*",
			"style": "header"
		},
		{
			"type": "dropdown",
			"id": "bussiness-line",
			"label": "Bussiness Line",
			"options": [
				{
					"type": "option",
					"id": "bug",
					"text": "Bug"
				},
				{
					"type": "option",
					"id": "feedback",
					"text": "Feedback"
				}
			]
		},
		{
			"type": "single-select",
			"id": "ticket-option",
			"label": "Priority",
			"options": [
				{
					"type": "option",
					"id": "option-1",
					"text": "P0 Urgent"
				},
				{
					"type": "option",
					"id": "option-1",
					"text": "P1 High"
				},
				{
					"type": "option",
					"id": "option-1",
					"text": "P2 Medium"
				}
			]
		},
		{
			"type": "single-select",
			"id": "ticket-option",
			"label": "Issue Target User",
			"options": [
				{
					"type": "option",
					"id": "option-1",
					"text": "Select from TK"
				},
				{
					"type": "option",
					"id": "option-2",
					"text": "Enter Mannually"
				}
			]
		},
		{
			"type": "spacer",
			"size": "l"
		},
		{
			"type": "spacer",
			"size": "l"
		},
		{
			"type": "button",
			"id": "submit-issue-form",
			"label": "Create issue",
			"style": "primary",
			"action": {
				"type": "submit"
			}
		}
	]`

	var components []Components
	err := json.Unmarshal([]byte(jsonData), &components)
	if err != nil {
		fmt.Println("error in unmarshall %v", err.Error())
	}

	return components
}

func GetInitTicketCanvasBody() CanvasResponse {
	return CanvasResponse{
		Canvas: Canvas{
			Content: Content{
				Components: InitPreOncallCanvas(),
			},
		},
	}
}

func GetCreateTicketCanvasBody() CanvasResponse {
	return CanvasResponse{
		Canvas: Canvas{
			Content: Content{
				Components: GetCreateTicketComponents(),
			},
		},
	}
}
