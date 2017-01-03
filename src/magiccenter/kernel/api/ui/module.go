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

// SingleModule 单模块
type SingleModule struct {
	common.Result
	Module model.Module
}

// ModuleList 模块列表
type ModuleList struct {
	common.Result
	ModuleList []model.Module
}

// GetModuleActionHandler 获取Module列表
func GetModuleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetModuleActionHandler")

	params := util.SplitParam(r.URL.RawQuery)
	mid, found := params["id"]
	if !found {
		result := ModuleList{}

		result.ModuleList = bll.QueryAllModules()
		result.ErrCode = 0

		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	} else {
		result := SingleModule{}

		for true {
			found := false
			result.Module, found = bll.QueryModule(mid)
			if found {
				result.ErrCode = 0
			} else {
				result.ErrCode = 1
				result.Reason = "无效参数"
			}

			break
		}

		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	}
}

// PostModuleActionHandler 新建Module
func PostModuleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostModuleActionHandler")
	result := SingleModule{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("module-name")
		description := r.FormValue("module-description")
		url := r.FormValue("module-url")
		mType := r.FormValue("module-type")
		
		mStatus := r.FormValue("module-status")
		if len(account) == 0 || len(email) == 0 {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		ret := false
		result.User, ret = commonbll.CreateUser(account, email)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "创建用户失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// ModuleBlock 模块结构
type ModuleBlock struct {
	common.Result
	Module    model.Module
	BlockList []model.Block
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

// BlockContent 功能块内容
type BlockContent struct {
	common.Result
	Block    model.Block
	ItemList []model.Item
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

// ModuleAuthorityGroup 模块管理组
type ModuleAuthorityGroup struct {
	common.Result
	Module    model.Module
	GroupList []model.Group
}

// GetModuleAuthorityGroupActionHandler 获取Module授权分组
func GetModuleAuthorityGroupActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetModuleAuthorityGroupActionHandler")

	result := ModuleAuthorityGroup{}
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
			result.GroupList, found = bll.GetModuleAuthGroup(id)
			if !found {
				result.ErrCode = 1
				result.Reason = "无效请求参数"
			} else {
				result.ErrCode = 0
			}
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
