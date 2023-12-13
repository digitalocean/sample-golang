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
	WorkspaceID     string       `json:"workspace_id"`
	WorkspaceRegion string       `json:"workspace_region"`
	Conversation    Conversation `json:"conversation"`
	InputValues     string       `json:"input_values"`
}

type Conversation struct {
	ConversationID  string `json:"id"`
	AdminAssigneeID string `json:"admin_assignee_id"`
}

func GetInitTicketCanvasBody() CanvasReponse {
	return InitPreOncallCanvas()
}

func GetCreateTicketCanvasBody(bizLines []string, regions []string, stackNames []string) CanvasReponse {
	return InitCreateOncalTicketCanvas(bizLines, regions, stackNames)
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
