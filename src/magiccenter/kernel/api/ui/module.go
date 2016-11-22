package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/api/bll"
	"magiccenter/system"
	"net/http"

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

// ModuleContent 模块内容
type ModuleContent struct {
	common.Result
	Module      model.Module
	ContentList []model.Item
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

	modulehub := system.GetModuleHub()
	modules := modulehub.QueryAllModule()
	for _, m := range modules {
		mod := model.Module{}
		mod.ID = m.ID()
		mod.Name = m.Name()
		mod.Description = m.Description()
		mod.URL = m.URL()
		mod.Type = m.Type()
		mod.Status = m.Status()

		result.ModuleList = append(result.ModuleList, mod)
	}

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
			result.Reason = "无效请求数据"
			break
		}

		modulehub := system.GetModuleHub()
		mod, found := modulehub.FindModule(id)
		if found {
			result.ErrCode = 0
			result.Module.ID = mod.ID()
			result.Module.Name = mod.Name()
			result.Module.Description = mod.Description()
			result.Module.URL = mod.URL()
			result.Module.Type = mod.Type()
			result.Module.Status = mod.Status()
		} else {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GetModuleContentActionHandler 获取Module Content信息
func GetModuleContentActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetModuleContentActionHandler")

	result := ModuleContent{}

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
