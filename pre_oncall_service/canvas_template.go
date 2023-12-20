package preoncall_service

import (
	"fmt"
)

// ==============Pre Oncall Intercom Struct==========================================================================

// IntercomCanvasRequest is the struct for the canvas request body from Intercom
type IntercomCanvasRequest struct {
	WorkspaceID     string                 `json:"workspace_id"`
	WorkspaceRegion string                 `json:"workspace_region"`
	Conversation    IntercomConversation   `json:"conversation"`
	InputValues     map[string]string      `json:"input_values"`
	ComponentID     string                 `json:"component_id"`
	CurrentCanvas   IntercomCanvasReceiver `json:"current_canvas"`
}

type IntercomConversation struct {
	ConversationID  string `json:"id"`
	AdminAssigneeID int    `json:"admin_assignee_id"`
}

// Component interface for all UI components.
type Component interface {
	Render() string
	GetID() string
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

func (b *Button) GetID() string {
	return b.ID
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

func (t *Text) GetID() string {
	return ""
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

func (i Input) GetID() string {
	return i.ID
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

func (ta *TextArea) GetID() string {
	return ta.ID
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

func (d *Dropdown) GetID() string {
	return d.ID
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

func (ss *SingleSelect) GetID() string {
	return ss.ID
}

// Spacer component
type Spacer struct {
	Type string
	Size string
}

func (s Spacer) Render() string {
	return ""
}

func (s Spacer) GetID() string {
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
