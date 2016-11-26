package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/api/bll"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// ModuleList 模块列表
type ModuleList struct {
	ModuleList []model.Module
}

// ModuleBlock 模块结构
type ModuleBlock struct {
	common.Result
	Module    model.Module
	BlockList []model.Block
}

// BlockContent 功能块内容
type BlockContent struct {
	common.Result
	Block    model.Block
	ItemList []model.Item
}

// ModuleAuthorityGroup 模块管理组
type ModuleAuthorityGroup struct {
	common.Result
	Module    model.Module
	GroupList []model.Group
}

// GetModuleListActionHandler 获取Module列表
func GetModuleListActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetModuleListActionHandler")

	result := ModuleList{}

	result.ModuleList = bll.QueryAllModules()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GetModuleBlockActionHandler 获取Module 功能块信息
func GetModuleBlockActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetModuleBlockActionHandler")

	result := ModuleBlock{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "非法请求数据"
			break
		}

		result.Module, found = bll.QueryModule(id)
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求参数"
		} else {
			result.ErrCode = 0
			result.BlockList = bll.QueryModuleBlocks(id)
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GetBlockItemActionHandler 获取Module Content信息
func GetBlockItemActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetBlockItemActionHandler")

	result := BlockContent{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "非法请求数据"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法请求参数"
		}

		result.Block, found = bll.GetModuleBlock(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求参数"
		} else {
			result.ItemList = bll.GetBlockItems(aid)
			result.ErrCode = 0
		}

		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GetModuleAuthorityGroupActionHandler 获取Module授权分组
func GetModuleAuthorityGroupActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetModuleAuthorityGroupActionHandler")

	result := ModuleAuthorityGroup{}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
