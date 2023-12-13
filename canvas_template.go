package main

// Define the structures based on the JSON input
const (
	RelatedTicketID = "related-ticket"
	CreateTicketID  = "create-ticket"
	SubmitTicketID  = "submit-ticket"
)

type InputValues struct {
	TicketOption string `json:"pre-oncall-ticket-option"`
}

type InitializeRequest struct {
	CurrentCanvas Canvas      `json:"current_canvas"`
	InputValues   InputValues `json:"input_values"`
}

type Workspace struct {
	WorkspaceID     string `json:"workspace_id"`
	WorkspaceRegion string `json:"workspace_region"`
	// ... include other fields as necessary
}

func GetInitTicketCanvasBody() CanvasReponse {
	return InitPreOncallCanvas()
}

//func GetCreateTicketCanvasBody() CanvasResponse {
//	return CanvasResponse{
//		Canvas: Canvas{
//			Content: Content{
//				Components: GetCreateTicketComponents(),
//			},
//		},
//	}
//}
