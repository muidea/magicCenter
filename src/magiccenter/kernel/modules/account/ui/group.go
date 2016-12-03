package ui

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/bll"
	"magiccenter/system"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// QueryAllGroup 查询所有分组
func QueryAllGroup(request *commonbll.QueryAllGroupRequest, response *commonbll.QueryAllGroupResponse) bool {
	response.Result.ErrCode = 0
	response.Groups = bll.QueryAllGroup()

	return true
}

// QueryGroups 查询指定分组
func QueryGroups(request *commonbll.QueryGroupsRequest, response *commonbll.QueryGroupsResponse) bool {
	response.Result.ErrCode = 0
	response.Groups = bll.QueryGroups(request.Ids)

	return true
}

// CreateGroup 新建分组
func CreateGroup(request *commonbll.CreateGroupRequest, response *commonbll.CreateGroupResponse) bool {
	response.Result.ErrCode = 0
	ret := false
	response.Group, ret = bll.CreateGroup(request.Name, request.Description)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "新建分组失败"
	}

	return true
}

// UpdateGroup 更新分组
func UpdateGroup(request *commonbll.UpdateGroupRequest, response *commonbll.UpdateGroupResponse) bool {
	response.Result.ErrCode = 0
	ret := false
	response.Group, ret = bll.SaveGroup(request.Group)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "更新分组失败"
	}

	return true
}

// DeleteGroup 更新分组
func DeleteGroup(request *commonbll.DeleteGroupRequest, response *commonbll.DeleteGroupResponse) bool {
	response.Result.ErrCode = 0
	ret := false
	ret = bll.DeleteGroup(request.ID)
	if !ret {
		response.Result.ErrCode = 1
		response.Result.Reason = "删除分组失败"
	}

	return true
}

// ManageGroupView 分组管理视图
type ManageGroupView struct {
	Groups []model.Group
}

// AllGroupList 所有分组结果
type AllGroupList struct {
	Groups []model.Group
}

// SingleGroup 当个分组结果
type SingleGroup struct {
	common.Result
	Group model.Group
}

// ManageGroupViewHandler 分组管理视图处理器
func ManageGroupViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageGroupViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/account/group.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageGroupView{}
	view.Groups = bll.QueryAllGroup()

	t.Execute(w, view)
}

// QueryAllGroupActionHandler 查询所有分组信息处理器
func QueryAllGroupActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllGroupActionHandler")

	result := AllGroupList{}
	result.Groups = bll.QueryAllGroup()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// QueryGroupActionHandler 查询单个分组信息处理器
func QueryGroupActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryGroupActionHandler")

	result := SingleGroup{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		gid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.Group, found = bll.QueryGroupByID(gid)
		if !found {
			result.ErrCode = 1
			result.Reason = "指定Group不存在"
			break
		}

		result.ErrCode = 0
		result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteGroupActionHandler 删除分组处理器
func DeleteGroupActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteGroupActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		gid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		ok := bll.DeleteGroup(gid)
		if !ok {
			result.ErrCode = 1
			result.Reason = "删除分组失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// SaveGroupActionHandler 保存分组信息处理器
func SaveGroupActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SaveGroupActionHandler")

	result := common.Result{}
	for true {
		err := r.ParseForm()
		if err != nil {
			log.Print("paseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id := r.FormValue("group-id")
		name := r.FormValue("group-name")
		desc := r.FormValue("group-description")

		gid := -1
		if len(id) > 0 {
			gid, err = strconv.Atoi(id)
			if err != nil {
				log.Printf("parse id failed, id:%s", id)
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}
		}

		group := model.Group{}
		group.ID = gid
		group.Name = name
		group.Description = desc
		_, ok := bll.SaveGroup(group)
		if !ok {
			result.ErrCode = 1
			result.Reason = "保存分组失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "保存分组成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
