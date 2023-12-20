package main

import (
	"context"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"strings"
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
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, nil)

	content := newContent([]Component{singleSelect})
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}

func InitRelatedTicketCanvas(ctx context.Context, oncallTickets TickeInfotResponse) CanvasReponse {
	////log. utils.GetLoggerWithMethod(ctx, "InitRelatedTicketCanvas")
	////log.Infof("InitRelatedTicketCanvas oncallTickets %v", larkcore.Prettify(oncallTickets))

	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, nil)

	components := []Component{}
	components = append(components, singleSelect)
	// TODO: Replace with the real ticket info
	for i := 1; i <= 2; i++ {
		ticketBannerTitle := NewText("Related Ticket", "header")
		components = append(components, ticketBannerTitle)

		ticketID := NewText("Ticket id", "paragraph")
		components = append(components, ticketID)

		bizLine := NewText("Biz Line", "paragraph")
		components = append(components, bizLine)

		TicketTitle := NewText("Ticket Title", "paragraph")
		components = append(components, TicketTitle)

		ReportBy := NewText("Report By", "paragraph")
		components = append(components, ReportBy)

		Assignee := NewText("Assignee", "paragraph")
		components = append(components, Assignee)

		CreateTime := NewText("Create Time", "paragraph")
		components = append(components, CreateTime)

		UpdateTime := NewText("Update Time", "paragraph")
		components = append(components, UpdateTime)

		AdditionalInfo := NewText("Additional Info", "paragraph")
		components = append(components, AdditionalInfo)

		GroupLink := NewText("Group Link", "paragraph")
		components = append(components, GroupLink)

	}

	content := newContent(components)
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}

// InitCreateOncalTicketCanvas is a constructor for Create Create-Ticket CanvasReponse
func InitCreateOncalTicketCanvas(bizLines []string, regions []string, stackNames []string, selectedValues map[string]string, validInput bool) CanvasReponse {
	if selectedValues == nil {
		selectedValues = make(map[string]string)
	}
	////log.Infof("InitCreateOncalTicketCanvas selectedValues %v", selectedValues)
	////log.Infof("InitCreateOncalTicketCanvas bizLines %v, regions %v, stackNames %v", bizLines, regions, stackNames)
	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
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
		////log.Infof("bizLine %v", bizLine)
		bizLineDropDownOptions = append(bizLineDropDownOptions, *NewOption(bizLine, bizLine))
	}

	bizLineSearchDropDown := NewDropdown(BizLineSearchDropdownID, BizLineSearchDropdownLabel, bizLineDropDownOptions, getValuePtr(BizLineSearchDropdownID, selectedValues))

	// ticket title
	ticketTitleText := NewText("Ticket Title", "header")

	ticketTitleInput := NewInput(TicketTitleInputID, TicketTitleLabel, "Briefly describe the problem", getValuePtr(TicketTitleInputID, selectedValues))

	// region search
	regionSearchText := NewText("Region Search", "header")
	regionSearchInput := NewInput(RegionSearchInputID, RegionSearchLabel, "Enter input here", nil)
	regionSearchBtn := NewButton(RegionSearchButtonID, RegionSearchButtonLabel, action, "primary", false)

	regionDropDownOptions := []Option{}
	for _, region := range regions {
		regionDropDownOptions = append(regionDropDownOptions, *NewOption(region, region))
	}

	regionSearchDropDown := NewDropdown(RegionSearchDropdownID, RegionSearchDropdownLabel, regionDropDownOptions, getValuePtr(RegionSearchDropdownID, selectedValues))

	// stack search
	stackSearchText := NewText("Stack Search", "header")
	stackDropDownOptions := []Option{}
	for _, stackOption := range stackNames {
		stackDropDownOptions = append(stackDropDownOptions, *NewOption(stackOption, stackOption))
	}

	stackSearchDropDown := NewDropdown(StackSearchDropdownID, StackSearchDropdownLabel, stackDropDownOptions, getValuePtr(StackSearchDropdownID, selectedValues))
	//
	// priority
	priorityText := NewText("Priority", "header")
	prioritySingleSelectOptions := []Option{}
	priorityList := []string{P0, P1, P2}
	for _, priority := range priorityList {
		prioritySingleSelectOptions = append(prioritySingleSelectOptions, *NewOption(priority, priority))
	}

	prioritySingleSelect := NewSingleSelect(PrioritySingleSelectID, "single-select", PrioritySingleSelectLabel, prioritySingleSelectOptions, nil, getValuePtr(PrioritySingleSelectID, selectedValues))

	// create group
	createGroupText := NewText("Create Group", "header")
	createGroupSingleSelectOptions := []Option{}
	createGroupList := []string{AutoCreateGroup, AssociateGroup, NotCreateGroup}
	for _, createGroup := range createGroupList {
		createGroupSingleSelectOptions = append(createGroupSingleSelectOptions, *NewOption(createGroup, createGroup))
	}

	createGroupSingleSelect := NewSingleSelect(CreateGroupSingleSelectID, "single-select", CreateGroupSingleSelectLabel, createGroupSingleSelectOptions, nil, getValuePtr(CreateGroupSingleSelectID, selectedValues))

	// user id

	userIDText := NewText("User ID", "header")
	userIDInput := NewInput(userIDInputID, userIDInputLabel, "type in user id", getValuePtr(userIDInputID, selectedValues))

	// tenant id
	tenantIDText := NewText("Tenant ID", "header")
	tenantIDInput := NewInput(tenantIDInputID, tenantIDInputLabel, "type in tenant id", getValuePtr(tenantIDInputID, selectedValues))

	// lark version
	larkVersionText := NewText("Lark Version", "header")
	larkVersionInput := NewInput(LarkVersionInputID, LarkVersionInputLabel, "type in lark version", getValuePtr(LarkVersionInputID, selectedValues))

	// Create button to submit ticket
	submitTicketBtn := NewButton(SubmitTicketButtonID, SubmitTicketLabel, action, "primary", false)

	content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn,
		bizLineSearchDropDown, ticketTitleText, ticketTitleInput, regionSearchText, regionSearchInput, regionSearchBtn, regionSearchDropDown,
		stackSearchText, stackSearchDropDown, priorityText, prioritySingleSelect,
		createGroupText, createGroupSingleSelect, userIDText, userIDInput,
		tenantIDText, tenantIDInput,
		larkVersionText, larkVersionInput,
		submitTicketBtn})

	//content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})

	//content := newContent([]Component{singleSelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})
	canvasResp := newCanvasReponse(*content)
	//fmt.Println(" InitCreateOncalTicketCanvas canvasResp %v", larkcore.Prettify(canvasResp))
	return *canvasResp
}

func GetInitTicketCanvasBody() CanvasReponse {
	return InitPreOncallCanvas()
}

func GetRelatedTicketCanvasBody(ctx context.Context, intercomConversationID string) CanvasReponse {
	////log. utils.GetLoggerWithMethod(ctx, "GetRelatedTicketCanvasBody")
	// We use the intercomConversationID to get the tickets external id via pre-oncall api
	////log.Infof("GetRelatedTicketCanvasBody intercomConversation %v", intercomConversationID)

	// TODO we use the intercomConversation to get the tickets via pre-oncall api
	oncallTickets := TickeInfotResponse{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
	return InitRelatedTicketCanvas(ctx, oncallTickets)
}

func searchBusinessLine(ctx context.Context, keyword string, bizLines []Business) []string {
	////log. utils.GetLoggerWithMethod(ctx, "searchBusinessLine")
	////log.Infof("searchBusinessLine keyword %v", keyword)
	result := make([]string, 0)
	for _, biz := range bizLines {
		if strings.Contains(strings.ToLower(biz.Name), strings.ToLower(keyword)) || keyword == "" {
			result = append(result, biz.Name+"-"+biz.Name)
		}
	}

	////log.Infof("searchBusinessLine result %v", result)
	return result
}

func searchRegion(ctx context.Context, keyword string, regions []CodeNamePair) []string {
	////log. utils.GetLoggerWithMethod(ctx, "searchRegion")
	////log.Infof("searchRegion keyword %v", keyword)
	result := make([]string, 0)
	for _, region := range regions {
		if strings.Contains(strings.ToLower(region.Name), strings.ToLower(keyword)) || keyword == "" {
			result = append(result)
		}
	}

	////log.Infof("searchRegion result %v", result)
	return result
}

func extractBizlinesFromCurrentCanvas(ctx context.Context, currentCanvas IntercomCanvasReceiver) []string {
	////log. utils.GetLoggerWithMethod(ctx, "extractBizlinesFromCurrentCanvas")
	//log.Infof("extractBizlinesFromCurrentCanvas currentCanvas %v", larkcore.Prettify(currentCanvas))
	bizLines := make([]string, 0)
	for _, component := range currentCanvas.Content.Components {
		if component.ID == BizLineSearchDropdownID {
			////log.Infof("extractBizlinesFromCurrentCanvas find bizline dropdown %v", larkcore.Prettify(component))
			existingOptions := component.Options
			for _, option := range existingOptions {
				bizLines = append(bizLines, option.Text)
			}
		}
	}

	////log.Infof("extractBizlinesFromCurrentCanvas bizLines %v", bizLines)

	return bizLines
}

func extractRegionsFromCurrentCanvas(ctx context.Context, currentCanvas IntercomCanvasReceiver) []string {
	////log. utils.GetLoggerWithMethod(ctx, "extractRegionsFromCurrentCanvas")
	//log.Infof("extractRegionsFromCurrentCanvas currentCanvas %v", larkcore.Prettify(currentCanvas))
	regions := make([]string, 0)
	for _, component := range currentCanvas.Content.Components {
		if component.ID == RegionSearchDropdownID {
			//log.Infof("extractRegionsFromCurrentCanvas find region dropdown %v", larkcore.Prettify(component))

			existingOptions := component.Options
			for _, option := range existingOptions {
				regions = append(regions, option.Text)
			}
		}
	}

	//log.Infof("extractRegionsFromCurrentCanvas regions %v", regions)
	return regions
}

func GetCreateTicketCanvasBody(ctx context.Context, inputValues map[string]string, intercomConversationID string, assigneeID int, buttonClick string, currentCanvas IntercomCanvasReceiver) CanvasReponse {
	//log. utils.GetLoggerWithMethod(ctx, "GetCreateTicketCanvasBody")
	fmt.Println("GetCreateTicketCanvasBody buttonClick %v, selectedValue %v, intercom convID %v, assigneeID %v, canvas %v", buttonClick, inputValues, intercomConversationID, assigneeID, larkcore.Prettify(currentCanvas))

	//metaInfoResp, err := GetPreOncallMetaInfo(ctx, true, true)
	//if err != nil {
	//	//log.Errorf("GetCreateTicketCanvasBody err %v", err)
	//	return InitPreOncallCanvas()
	//}

	bizLines := extractBizlinesFromCurrentCanvas(ctx, currentCanvas)
	regions := extractRegionsFromCurrentCanvas(ctx, currentCanvas)
	stackNames := make([]string, 0)
	fmt.Println("GetCreateTicketCanvasBody buttonClick %v", bizLines)
	fmt.Println("GetCreateTicketCanvasBody buttonClick %v", regions)
	if buttonClick == CreateTicketOptionID {
		//log.Infof("GetCreateTicketCanvasBody create ticket option")
		resp, err := GetPreOncallMetaInfo(ctx, true, true)
		if err != nil {
			//log.Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			fmt.Println("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			return InitPreOncallCanvas()
		}
		//log.Infof("GetCreateTicketCanvasBody resp %v", larkcore.Prettify(resp))
		bizList := resp.Data.BusinessList
		for _, biz := range bizList {
			bizLines = append(bizLines, biz.Name+"-"+biz.Name)
		}
		//log.Infof("GetCreateTicketCanvasBody bizLines %v", bizLines)

		regionList := resp.Data.RegionList
		for _, region := range regionList {
			regions = append(regions, region.Name)
		}

		//log.Infof("GetCreateTicketCanvasBody regions %v", regions)

		larkVersion := ""
		tenantID := ""
		userID := ""

		//convSrv := service.GetLarkAssistantConversationRecordService()
		//conversation, err := convSrv.GetConversationByIntercomConversationIDWithRetry(ctx, intercomConversationID)
		//
		//if err == nil {
		//	larkConverstionID := conversation.LarkConversationID
		//	//log.Infof("GetCreateTicketCanvasBody find larkConverstionID %v", larkConverstionID)
		//
		//	// get user info related to this conversation
		//	userSrv := service.GetLarkAssistantUserRecordService()
		//	user, userErr := userSrv.GetLarkAssistantUserByConversationIDWithRetry(ctx, larkConverstionID)
		//
		//	if userErr == nil {
		//		//log.Infof("GetCreateTicketCanvasBody find user %v", larkcore.Prettify(user))
		//		// Prefill the user id and tenant id
		//		tenantID = user.TenantID
		//		userID = user.UserID
		//		userOpenID := user.UserOpenID
		//
		//		// Begin to fetch the lark version
		//		redisCli := cache.GetClient()
		//		segmentID, cacheErr := intercom.GetSegmentID(ctx, redisCli, userOpenID)
		//
		//		if cacheErr == nil {
		//			segSrv := service.GetLarkAssistantSegmentInfoRecordService()
		//			//log.Infof("GetCreateTicketCanvasBody find segment info by segmentID %v and userOpenID %v", segmentID)
		//			segmentInfo, segErr := segSrv.GetBySegmentIDOrUserID(ctx, segmentID, userOpenID)
		//			if segmentInfo != nil && segErr == nil {
		//				// Prefill the lark version
		//				larkVersion = segmentInfo.ExtraInfo.Version
		//				//log.Infof("GetCreateTicketCanvasBody find lark version %v", larkVersion)
		//			}
		//		}
		//
		//	}
		//
		//}
		inputValues[userIDInputID] = userID
		inputValues[tenantIDInputID] = tenantID
		inputValues[LarkVersionInputID] = larkVersion
	}

	if buttonClick == BizLineSearchButtonID {
		var bizLineSearchKeyword string
		if v, ok := inputValues[BizLineSearchInputID]; ok {
			bizLineSearchKeyword = v
		}

		resp, err := GetPreOncallMetaInfo(ctx, true, true)
		if err != nil {
			//log.Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			return InitPreOncallCanvas()
		}
		bussinessList := resp.Data.BusinessList
		bizLines = searchBusinessLine(ctx, bizLineSearchKeyword, bussinessList)
	}

	if buttonClick == RegionSearchButtonID {
		var regionSearchKeyword string
		if v, ok := inputValues[RegionSearchInputID]; ok {
			regionSearchKeyword = v
		}

		resp, err := GetPreOncallMetaInfo(ctx, true, true)
		if err != nil {
			//log.Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			return InitPreOncallCanvas()
		}
		regionList := resp.Data.RegionList
		regions = searchRegion(ctx, regionSearchKeyword, regionList)
	}

	return InitCreateOncalTicketCanvas(bizLines, stackNames, stackNames, inputValues, true)
}
