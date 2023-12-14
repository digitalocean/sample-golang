package main

import (
	"fmt"
)

// Component interface for all UI components.
type Component interface {
	Render() string
}

// Action defines the action to be taken when the button is clicked.
type Action struct {
	Type string `json:"type"` // Can be "submit", "URL", or "sheet"
	// Additional fields can be added here based on the action's requirements
}

// NewAction is a constructor for Action
func NewAction(actionType string) Action {
	return Action{Type: actionType}
}

// Button represents a UI button component.
type Button struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Label    string `json:"label"`
	Action   Action `json:"action"`
	Style    string `json:"style,omitempty"`    // Primary, Secondary, Link
	Disabled bool   `json:"disabled,omitempty"` // Default is false
}

// NewButton creates a new button with the given parameters.
func NewButton(id, label string, action Action, style string, disabled bool) *Button {
	return &Button{
		Type:     "button",
		ID:       id,
		Label:    label,
		Action:   action,
		Style:    style,
		Disabled: disabled,
	}
}

// Render method for Button
func (b *Button) Render() string {
	return fmt.Sprintf("Button ID: %s, Label: %s, Action: %s, Style: %s, Disabled: %v", b.ID, b.Label, b.Action.Type, b.Style, b.Disabled)
}

// Text component
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Style string `json:"style,omitempty"`
}

// NewText is a constructor for Text
func NewText(text, style string) *Text {
	return &Text{Type: "text", Text: text, Style: style}
}

// Render method for Text
func (t *Text) Render() string {
	return fmt.Sprintf("Text: %s, Style: %s", t.Text, t.Style)
}

// Input component
type Input struct {
	Type        string  `json:"type"`
	ID          string  `json:"id"`
	Label       string  `json:"label"`
	Placeholder string  `json:"placeholder"`
	Value       *string `json:"value,omitempty"`
}

func (i Input) Render() string {
	return ""
}

// NewInput is a constructor for Input
func NewInput(id, label, placeholder string, value *string) *Input {
	return &Input{Type: "input", ID: id, Label: label, Placeholder: placeholder, Value: value}
}

// TextArea component
type TextArea struct {
	Type        string
	ID          string
	Label       string
	Placeholder string
}

// NewTextArea is a constructor for TextArea
func NewTextArea(id, label, placeholder string) *TextArea {
	return &TextArea{Type: "textarea", ID: id, Label: label, Placeholder: placeholder}
}

// Render method for TextArea
func (ta *TextArea) Render() string {
	return fmt.Sprintf("TextArea ID: %s, Label: %s, Placeholder: %s", ta.ID, ta.Label, ta.Placeholder)
}

// Option for Dropdown and SingleSelect
type Option struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

// NewOption is a constructor for Option
func NewOption(id, text string) *Option {
	return &Option{Type: "option", ID: id, Text: text}
}

// Dropdown component
type Dropdown struct {
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Options []Option `json:"options"`
	Value   *string  `json:"value,omitempty"`
}

// NewDropdown is a constructor for Dropdown
func NewDropdown(id, label string, options []Option, value *string) *Dropdown {
	return &Dropdown{Type: "dropdown", ID: id, Label: label, Options: options, Value: value}
}

// Render method for Dropdown
func (d *Dropdown) Render() string {
	return fmt.Sprintf("Dropdown ID: %s, Label: %s, Options: %v", d.ID, d.Label, d.Options)
}

// SingleSelect component
type SingleSelect struct {
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Options []Option `json:"options"`
	Action  *Action  `json:"action"`
	Value   *string  `json:"value,omitempty"`
}

func NewSingleSelect(id, selectType, label string, options []Option, action *Action, value *string) *SingleSelect {
	return &SingleSelect{Type: selectType, ID: id, Label: label, Options: options, Action: action, Value: value}
}

// Render method for SingleSelect
func (ss *SingleSelect) Render() string {
	return fmt.Sprintf("SingleSelect ID: %s, Label: %s, Options: %v", ss.ID, ss.Label, ss.Options)
}

// Spacer component
type Spacer struct {
	Type string
	Size string
}

func (s Spacer) Render() string {
	return ""
}

// NewSpacer is a constructor for Spacer
func NewSpacer(size string) *Spacer {
	return &Spacer{Type: "spacer", Size: size}
}

// NewCanvas is a constructor for Canvas
func Newcontent(components []Component) *Content {
	return &Content{Components: components}
}

// AddComponent adds a component to the canvas
func (c *Content) AddComponent(component Component) {
	c.Components = append(c.Components, component)
}

// Content represents the content field within canvas.
type Content struct {
	Components []Component `json:"components"`
}

// Canvas represents the top-level canvas field in your JSON.
type Canvas struct {
	Content Content `json:"content"`
}

// Root structure to encapsulate the Canvas
type CanvasReponse struct {
	Canvas Canvas `json:"canvas"`
}

func newContent(components []Component) *Content {
	return &Content{Components: components}
}

func newCanvas(content Content) *Canvas {
	return &Canvas{Content: content}
}

func newCanvasReponse(content Content) *CanvasReponse {
	canvas := newCanvas(content)
	return &CanvasReponse{Canvas: *canvas}
}

//func CreateDemoCanvas() *CanvasReponse {
//	// Creating components using constructor functions
//	text := NewText("*Create a ticket*", "header")
//	input := NewInput("title", "Title", "Enter a title for your issue...")
//	textArea := NewTextArea("description", "Description", "Enter a description of the issue...")
//	option1 := NewOption("bug", "Bug")
//	option2 := NewOption("feedback", "Feedback")
//	dropdown := NewDropdown("label", "Label", []Option{*option1, *option2})
//	option3 := NewOption("low", "Low")
//	option4 := NewOption("medium", "Medium")
//	option5 := NewOption("high", "High")
//	singleSelect := NewSingleSelect("priority", "Priority", []Option{*option3, *option4, *option5})
//	spacer := NewSpacer("s")
//	action := NewAction("submit")
//	button := NewButton("submit", "Submit", action, "primary", false)
//
//	// Creating a canvas and adding components
//	content := newContent([]Component{text, input, textArea, dropdown, singleSelect, spacer, button})
//	canvasResp := newCanvasReponse(*content)
//
//	// Marshalling struct back to JSON
//	marshalledData, err := json.MarshalIndent(canvasResp, "", "  ")
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	fmt.Println("Marshalled Data:")
//	fmt.Println(string(marshalledData))
//	// Optionally, marshal to JSON for demonstration
//	return canvasResp
//}

func getValuePtr(key string, selectedValues map[string]string) *string {
	var selectedValue *string
	if val, exist := selectedValues[key]; exist && val != "" {
		selectedValue = &val
	}
	fmt.Println("getValuePtr selectedValue %v", selectedValue)
	return selectedValue
}

func InitPreOncallCanvas() CanvasReponse {
	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(SummitTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, nil)

	content := newContent([]Component{singleSelect})
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}

func InitCreateOncalTicketCanvas(bizLines []string, regions []string, stackNames []string, selectedValues map[string]string, validInput bool) CanvasReponse {
	if selectedValues == nil {
		selectedValues = make(map[string]string)
	}

	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(SummitTicketOptionID, "Create Ticket")
	action := NewAction("submit")

	categorySelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, getValuePtr(CategorySingleSelectID, selectedValues))

	// bizline
	bizLineText := NewText("Business Line Search", "header")

	var bizLineSearchValue *string
	if val, exist := selectedValues[BizLineSearchInputID]; exist {
		bizLineSearchValue = &val
	}

	bizLineSearchInput := NewInput(BizLineSearchInputID, BizLineSearchLabel, "Enter input here", bizLineSearchValue)
	bizLineSearchBtn := NewButton(BizLineSearchButtonID, BizLineSearchButtonLabel, action, "primary", false)
	bizLineDropDownOptions := []Option{}
	for _, bizLine := range bizLines {
		bizLineDropDownOptions = append(bizLineDropDownOptions, *NewOption(bizLine, bizLine))
	}

	bizLineSearchDropDown := NewDropdown(BizLineSearchDropdownID, BizLineSearchDropdownLabel, bizLineDropDownOptions, getValuePtr(BizLineSearchDropdownID, selectedValues))

	// ticket title
	ticketTitleText := NewText("Ticket Title", "header")

	var ticketTitleValue *string
	if val, exist := selectedValues[TicketTitleInputID]; exist {
		ticketTitleValue = &val
	}
	ticketTitleInput := NewInput(TicketTitleInputID, TicketTitleLabel, "Briefly describe the problem", ticketTitleValue)

	// region search
	regionSearchText := NewText("Region Search", "header")
	regionSearchInput := NewInput(RegionSearchInputID, RegionSearchLabel, "Enter input here", nil)
	regionSearchBtn := NewButton(RegionSearchButtonID, RegionSearchButtonLabel, action, "primary", false)

	regionDropDownOptions := []Option{}
	for _, region := range regions {
		regionDropDownOptions = append(regionDropDownOptions, *NewOption(region, region))
	}

	var regionSelectedValue *string
	if val, exist := selectedValues[RegionSearchDropdownID]; exist {
		regionSelectedValue = &val
	}
	regionSearchDropDown := NewDropdown(RegionSearchDropdownID, RegionSearchDropdownLabel, regionDropDownOptions, regionSelectedValue)

	// stack search
	stackSearchText := NewText("Stack Search", "header")
	stackDropDownOptions := []Option{}
	for _, stackOption := range stackNames {
		stackDropDownOptions = append(stackDropDownOptions, *NewOption(stackOption, stackOption))
	}

	var stackSelectedValue *string
	if val, exist := selectedValues[StackSearchDropdownID]; exist {
		stackSelectedValue = &val
	}
	stackSearchDropDown := NewDropdown(StackSearchDropdownID, StackSearchDropdownLabel, stackDropDownOptions, stackSelectedValue)
	//
	//// priority
	//priorityText := NewText("Priority", "header")
	//prioritySingleSelectOptions := []Option{}
	//priorityList := []string{P0, P1, P2}
	//for _, priority := range priorityList {
	//	prioritySingleSelectOptions = append(prioritySingleSelectOptions, *NewOption(priority, priority))
	//}
	//
	//var prioritySelectedValue *string
	//if val, exist := selectedValues[PrioritySingleSelectID]; exist {
	//	prioritySelectedValue = &val
	//}
	//prioritySingleSelect := NewSingleSelect(PrioritySingleSelectID, "single-select", PrioritySingleSelectLabel, prioritySingleSelectOptions, nil, prioritySelectedValue)
	//
	//// create group
	//createGroupText := NewText("Create Group", "header")
	//createGroupSingleSelectOptions := []Option{}
	//createGroupList := []string{AutoCreateGroup, AssociateGroup, NotCreateGroup}
	//for _, createGroup := range createGroupList {
	//	createGroupSingleSelectOptions = append(createGroupSingleSelectOptions, *NewOption(createGroup, createGroup))
	//}
	//
	//var createGroupSelectedValue *string
	//if val, exist := selectedValues[CreateGroupSingleSelectID]; exist {
	//	createGroupSelectedValue = &val
	//}
	//createGroupSingleSelect := NewSingleSelect(CreateGroupSingleSelectID, "single-select", CreateGroupSingleSelectLabel, createGroupSingleSelectOptions, nil, createGroupSelectedValue)
	//
	//// user id
	//var userIDInputValue *string
	//if val, exist := selectedValues[userIDInputID]; exist {
	//	userIDInputValue = &val
	//}
	//userIDText := NewText("User ID", "header")
	//userIDInput := NewInput(userIDInputID, userIDInputLabel, "type in user id", userIDInputValue)
	//
	//// tenant id
	//var tenantIDInputValue *string
	//if val, exist := selectedValues[tenantIDInputID]; exist {
	//	tenantIDInputValue = &val
	//}
	//tenantIDText := NewText("Tenant ID", "header")
	//tenantIDInput := NewInput(tenantIDInputID, tenantIDInputLabel, "type in tenant id", tenantIDInputValue)
	//
	//// lark version
	//var larkVersionInputValue *string
	//if val, exist := selectedValues[LarkVersionInputID]; exist {
	//	larkVersionInputValue = &val
	//}
	//larkVersionText := NewText("Lark Version", "header")
	//larkVersionInput := NewInput(LarkVersionInputID, LarkVersionInputLabel, "type in lark version", larkVersionInputValue)
	//
	//// Create button to submit ticket
	//submitTicketBtn := NewButton(SubmitTicketButtonID, SubmitTicketLabel, action, "primary", false)
	//
	content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn,
		bizLineSearchDropDown, ticketTitleText, ticketTitleInput,
		regionSearchText, regionSearchInput, regionSearchBtn, regionSearchDropDown,
		stackSearchText, stackSearchDropDown,
		//priorityText, prioritySingleSelect,
		//createGroupText, createGroupSingleSelect,
		//userIDText, userIDInput,
		//tenantIDText, tenantIDInput,
		//larkVersionText, larkVersionInput,
		//submitTicketBtn
	})

	//content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})

	//content := newContent([]Component{singleSelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})
	canvasResp := newCanvasReponse(*content)
	//fmt.Println(" InitCreateOncalTicketCanvas canvasResp %v", larkcore.Prettify(canvasResp))
	return *canvasResp
}
