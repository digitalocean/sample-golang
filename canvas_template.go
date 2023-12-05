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
