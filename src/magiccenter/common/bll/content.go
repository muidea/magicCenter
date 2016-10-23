package bll

import (
	"magiccenter/common"
	commonmodel "magiccenter/common/model"
	"magiccenter/system"
)

// ContentModuleID ID
const ContentModuleID = "ffe1c37f-4fea-4a03-a7c3-55a331f5995f"

// QueryModuleBlockRequest 查询指定Module拥有的功能块列表请求
type QueryModuleBlockRequest struct {
	ID string
}

// QueryModuleBlockResponse 查询指定Module拥有的功能块列表响应
type QueryModuleBlockResponse struct {
	Blocks []commonmodel.Block
}

// QueryModulePageRequest 查询指定Module的页面请求
type QueryModulePageRequest struct {
	ID  string
	URL string
}

// QueryModulePageResponse 查询指定Module的页面响应
type QueryModulePageResponse struct {
	Result common.Result
	Page   []commonmodel.Page
}

// QueryBlockItemRequest 查询指定Block拥有的条目请求
type QueryBlockItemRequest struct {
	ID int
}

// QueryBlockItemResponse 查询指定Block拥有的条目响应
type QueryBlockItemResponse struct {
	Result common.Result
	Items  []commonmodel.Item
}

// QueryModuleBlock 查询指定Module的功能块
func QueryModuleBlock(id string) ([]commonmodel.Block, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryModuleBlockRequest{}
	request.ID = id

	response := QueryModuleBlockResponse{}
	result := contentModule.Invoke(&request, &response)

	return response.Blocks, result
}

// QueryModulePage 查询指定Module的页面
func QueryModulePage(id string) ([]commonmodel.Page, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryModulePageRequest{}
	request.ID = id

	response := QueryModulePageResponse{}
	result := contentModule.Invoke(&request, &response)
	if result {
		result = response.Result.Success()
	}

	return response.Page, result
}

// QueryBlockItem 查询指定Block拥有的Items
func QueryBlockItem(id int) ([]commonmodel.Item, bool) {
	moduleHub := system.GetModuleHub()
	contentModule, found := moduleHub.FindModule(ContentModuleID)
	if !found {
		panic("can't find account module")
	}

	request := QueryBlockItemRequest{}
	request.ID = id

	response := QueryBlockItemResponse{}
	result := contentModule.Invoke(&request, &response)
	if result {
		result = response.Result.Success()
	}

	return response.Items, result
}
