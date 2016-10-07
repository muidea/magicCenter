package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/configuration"
	"magiccenter/kernel/modules/dashboard/modulemanage/bll"
	"net/http"
	"strings"
)

// ModuleManageView Module管理视图内容
type ModuleManageView struct {
	Modules       []model.Module
	DefaultModule string
}

// ModuleList 模块列表
type ModuleList struct {
	Modules []model.Module
}

// ModuleManageViewHandler Module管理视图处理器
func ModuleManageViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModuleManageViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/module.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ModuleManageView{}
	view.Modules = bll.QueryAllModules()
	view.DefaultModule, _ = configuration.GetOption(configuration.SysDefaultModule)

	t.Execute(w, view)
}

// AjaxModuleActionHandler 更新Module
func AjaxModuleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxModuleActionHandler")

	result := common.Result{}
	for {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}

		enableModuleList := r.FormValue("enable-list")
		defaultModule := r.FormValue("default-module")

		moduleIds := strings.Split(enableModuleList, ",")
		_, ok := bll.EnableModules(moduleIds)
		if ok {
			ok = configuration.SetOption(configuration.SysDefaultModule, defaultModule)
		} else {
			result.ErrCode = 1
			result.Reason = "启动Moule失败"
		}

		if ok {
			result.ErrCode = 0
			result.Reason = "更新Module信息成功'"
		} else {
			result.ErrCode = 1
			result.Reason = "设置默认模块失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// ModuleListActionHandler 获取模块列表处理器
func ModuleListActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModuleManageViewHandler")

	result := ModuleList{}
	result.Modules = bll.QueryAllModules()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
