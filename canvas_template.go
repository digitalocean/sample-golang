package main

// Define the structures based on the JSON input
const (
	RelatedTicketID = "related-ticket"
	SubmitTicketID  = "submit-ticket"
)

type Workspace struct {
	WorkspaceID     string            `json:"workspace_id"`
	WorkspaceRegion string            `json:"workspace_region"`
	Conversation    Conversation      `json:"conversation"`
	InputValues     map[string]string `json:"input_values"`
}

type Conversation struct {
	ConversationID  string `json:"id"`
	AdminAssigneeID int    `json:"admin_assignee_id"`
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
