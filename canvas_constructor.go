package main

import (
	"context"
	"fmt"
)

// getValuePtr is a helper function to get the pointer of the value
// To fill in the canvas fields with the previous input values
func getValuePtr(key string, selectedValues map[string]string) *string {
	var selectedValue *string
	if val, exist := selectedValues[key]; exist && val != "" {
		selectedValue = &val
	}
	fmt.Printf("getValuePtr selectedValue %v \n", selectedValue)
	return selectedValue
}

// InitPreOncallCanvas is a constructor for Init CanvasReponse
func InitPreOncallCanvas() CanvasReponse {
	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(SummitTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, nil)

	content := newContent([]Component{singleSelect})
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}

// InitCreateOncalTicketCanvas is a constructor for Create Create-Ticket CanvasReponse
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

	ticketTitleInput := NewInput(TicketTitleInputID, TicketTitleLabel, "Briefly describe the problem", getValuePtr(TicketTitleInputID, selectedValues))

	// region search
	//regionSearchText := NewText("Region Search", "header")
	//regionSearchInput := NewInput(RegionSearchInputID, RegionSearchLabel, "Enter input here", nil)
	//regionSearchBtn := NewButton(RegionSearchButtonID, RegionSearchButtonLabel, action, "primary", false)
	//
	//regionDropDownOptions := []Option{}
	//for _, region := range regions {
	//	regionDropDownOptions = append(regionDropDownOptions, *NewOption(region, region))
	//}
	//
	//regionSearchDropDown := NewDropdown(RegionSearchDropdownID, RegionSearchDropdownLabel, regionDropDownOptions, getValuePtr(RegionSearchDropdownID, selectedValues))
	//
	//// stack search
	//stackSearchText := NewText("Stack Search", "header")
	//stackDropDownOptions := []Option{}
	//for _, stackOption := range stackNames {
	//	stackDropDownOptions = append(stackDropDownOptions, *NewOption(stackOption, stackOption))
	//}
	//
	//stackSearchDropDown := NewDropdown(StackSearchDropdownID, StackSearchDropdownLabel, stackDropDownOptions, getValuePtr(StackSearchDropdownID, selectedValues))
	////
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

	content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn,
		bizLineSearchDropDown, ticketTitleText, ticketTitleInput,
		//regionSearchText, regionSearchInput, regionSearchBtn, regionSearchDropDown,
		//stackSearchText, stackSearchDropDown,
		//priorityText, prioritySingleSelect,
		//createGroupText, createGroupSingleSelect,
		//userIDText, userIDInput,
		//tenantIDText, tenantIDInput,
		//larkVersionText, larkVersionInput,
		//submitTicketBtn,
	})

	//content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})

	//content := newContent([]Component{singleSelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})
	canvasResp := newCanvasReponse(*content)
	//fmt.Println(" InitCreateOncalTicketCanvas canvasResp %v", larkcore.Prettify(canvasResp))
	return *canvasResp
}

func GetInitTicketCanvasBody() CanvasReponse {
	return InitPreOncallCanvas()
}

func GetCreateTicketCanvasBody(ctx context.Context, selectedValue map[string]string) CanvasReponse {

	fmt.Printf("GetCreateTicketCanvasBody selectedValue %v \n", selectedValue)

	//metaInfoResp, err := GetPreOncallMetaInfo(ctx, true, true)
	//if err != nil {
	//	fmt.Printf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v \n", err)
	//	return InitPreOncallCanvas()
	//}

	//bizLines := make([]string, 0)
	//regions := make([]string, 0)
	//stackNames := make([]string, 0)
	//
	//businessList := metaInfoResp.Data.BusinessList
	//for idx, _ := range businessList {
	//	bussiness := businessList[idx]
	//	bizLines = append(bizLines, bussiness.Name)
	//	stacks := bussiness.Stacks
	//	for _, stack := range stacks {
	//		stackNames = append(stackNames, bussiness.Name+"-"+stack)
	//	}
	//}
	//
	//regionList := metaInfoResp.Data.RegionList
	//for idx, _ := range regionList {
	//	region := regionList[idx]
	//	regions = append(regions, region.Name)
	//}
	//bizLines := make([]string, 0)
	//bizLines = append(bizLines, "test1")
	//regions := make([]string, 0)
	//regions = append(regions, "test2")
	//stackNames := make([]string, 0)
	//stackNames = append(stackNames, "test3")

	//fmt.Printf("GetCreateTicketCanvasBody bizLines %v \n", bizLines)
	//fmt.Printf("GetCreateTicketCanvasBody regions %v \n", regions)
	//fmt.Printf("GetCreateTicketCanvasBody stackNames %v \n", stackNames)
	bizLines := make([]string, 0)
	bizLines = append(bizLines, "abc")
	bizLines = append(bizLines, "bbm")
	bizLines = append(bizLines, "cc")

	return InitCreateOncalTicketCanvas(bizLines, []string{"1", "2", "3"}, []string{"t", "e", "l"}, selectedValue, true)
}
