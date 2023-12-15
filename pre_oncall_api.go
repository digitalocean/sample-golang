package main

// Define a struct to match the JSON request body
type MetaInfoApiRequestBody struct {
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

//func call_metainfo_api() MetaInfoApiResponse {
//	// Create an instance of the request body
//	body := requestBody{
//		Business: true,
//		Stack:    true,
//		Region:   true,
//	}
//
//	// Marshal the request body into JSON
//	jsonBody, err := json.Marshal(body)
//	if err != nil {
//		fmt.Println("Error marshaling JSON:", err)
//		return MetaInfoApiResponse{}
//	}
//
//	// Define the API endpoint and token
//	url := "https://lark-oncall.bytedance.net/openapi/ticket/v1/getMetaInfo"
//	token := "26ad213fcdc54e0da3a6e7fc79e99b75"
//
//	// Create a new HTTP client
//	client := &http.Client{}
//
//	// Create a new request
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
//	if err != nil {
//		fmt.Println("Error creating request:", err)
//		return MetaInfoApiResponse{}
//	}
//
//	// Add headers
//	req.Header.Set("Authorization", "Basic "+token)
//	req.Header.Set("Content-Type", "application/json")
//
//	// Send the request
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Error sending request:", err)
//		return MetaInfoApiResponse{}
//	}
//	defer resp.Body.Close()
//
//	// Read the response body
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Error reading response body:", err)
//		return MetaInfoApiResponse{}
//	}
//
//	// Print the response body
//	fmt.Println("Response:", string(respBody))
//
//	var response MetaInfoApiResponse
//	err = json.Unmarshal(respBody, &response)
//	if err != nil {
//		fmt.Println("Error unmarshaling JSON:", err)
//		return MetaInfoApiResponse{}
//	}
//
//	// Print the response struct to verify
//	fmt.Printf("%+v\n", response)
//
//	return response
//}
