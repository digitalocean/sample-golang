package pre_oncall

// ==============Pre Oncall Meta Info API Struct==========================================================================

// Define a struct to match the JSON request body
type MetaInfoApiRequest struct {
	Business      bool `json:"business"`
	Stack         bool `json:"stack"`
	Type          bool `json:"type,omitempty"`
	Priority      bool `json:"priority,omitempty"`
	App           bool `json:"app,omitempty"`
	CreateChatWay bool `json:"createChatWay,omitempty"`
	External      bool `json:"external,omitempty"`
	Source        bool `json:"source,omitempty"`
	Region        bool `json:"region,omitempty"`
	Status        bool `json:"status,omitempty"`
}

// Define the structs to match the JSON response structure
type MetaInfoApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	BusinessList      []Business     `json:"businessList"`
	TypeList          []CodeNamePair `json:"typeList"`
	PriorityList      []CodeNamePair `json:"priorityList"`
	AppList           []CodeNamePair `json:"appList"`
	CreateChatWayList []CodeNamePair `json:"createChatWayList"`
	ExternalList      []CodeNamePair `json:"externalList"`
	SourceList        []CodeNamePair `json:"sourceList"`
	RegionList        []CodeNamePair `json:"regionList"`
	StatusList        []CodeNamePair `json:"statusList"`
}

type Business struct {
	Bid      string   `json:"bid"`
	Enabled  bool     `json:"enabled"`
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	ParentId string   `json:"parentId"`
	Inherit  bool     `json:"inherit"`
	Stacks   []string `json:"stacks"`
}

type CodeNamePair struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ==============Pre Oncall Ticket API Struct==========================================================================
// TicketRequest represents the JSON structure for the API request
type TicketSubmitRequest struct {
	Title         string `json:"title"`
	Business      string `json:"business"`
	Priority      string `json:"priority"`
	Stack         string `json:"stack"`
	Region        string `json:"region"`
	UserId        string `json:"userId"`
	Version       string `json:"version"`
	CreateChatWay string `json:"createChatWay"`
	Type          string `json:"type"`
	App           string `json:"app"`
	External      string `json:"external"`
	Source        string `json:"source"`
	Reporter      string `json:"reporter"`
	Remarks       string `json:"remarks"`
	ChannelType   string `json:"channelType"`
	BizTicketId   string `json:"bizTicketId"`
}

// TicketResponse represents the JSON structure for the API response
type TicketSubmitResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TicketId string `json:"ticketId"`
	} `json:"data"`
}

// ==============Pre Oncall Ticket Request Struct==========================================================================
type TickeInfotResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []Ticket `json:"data"`
}

type Ticket struct {
	TicketId     string   `json:"ticketId"`
	Title        string   `json:"title"`
	Business     string   `json:"business"`
	BusinessName string   `json:"businessName"`
	Assignee     []string `json:"assignee"`
	Status       string   `json:"status"`
	Reporter     string   `json:"reporter"`
	OpenChatId   string   `json:"openChatId"`
	Remarks      string   `json:"remarks"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
	GroupLink    string   `json:"groupLink"`
	TicketLink   string   `json:"ticketLink"`
	ChannelType  string   `json:"channelType"`
	BizTicketId  string   `json:"bizTicketId"`
}
