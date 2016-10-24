package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/dashboard/content/bll"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// SingleBlock 单项Block
type SingleBlock struct {
	Result common.Result
	Block  model.Block
}

// SingleBlockDetail 单项Block
type SingleBlockDetail struct {
	Result common.Result
	Block  model.BlockDetail
}

func getSingleBlock(w http.ResponseWriter, r *http.Request) {
	result := SingleBlock{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		result.Block, found = bll.GetModuleBlock(aid)
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Result.ErrCode = 0
		result.Result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

func getSingleBlockDetail(w http.ResponseWriter, r *http.Request) {
	result := SingleBlockDetail{}

	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		result.Block, found = bll.GetModuleBlockDetail(aid)
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Result.ErrCode = 0
		result.Result.Reason = "查询成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// GetBlockActionHandler 查询Page拥有的Block信息
func GetBlockActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetBlockActionHandler")

	params := util.SplitParam(r.URL.RawQuery)
	action, found := params["action"]
	if !found || action != "detail" {
		getSingleBlock(w, r)
	} else {
		getSingleBlockDetail(w, r)
	}
}

// PostBlockActionHandler 新建Block
func PostBlockActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostBlockActionHandler")

	result := SingleBlock{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		result.Block.Name = r.FormValue("name")
		result.Block.Tag = r.FormValue("tag")
		result.Block.Owner = r.FormValue("module")
		style := r.FormValue("style")
		result.Block.Style, err = strconv.Atoi(style)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
		}

		ret := false
		result.Block, ret = bll.AppendModuleBlock(result.Block)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "保存失败"
			break
		}

		result.Result.ErrCode = 0
		result.Result.Reason = "保存成功"
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteBlockActionHandler 删除Block
func DeleteBlockActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteBlockActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["module"]
		if !found {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}
		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		ret := bll.DeleteModuleBlock(aid)
		if !ret {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		result.ErrCode = 0
		result.Reason = "删除成功"
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// PutBlockActionHandler 更新Block
func PutBlockActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutBlockActionHandler")

	result := SingleBlock{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		result.Block.Name = r.FormValue("name")
		result.Block.Tag = r.FormValue("tag")
		result.Block.Owner = r.FormValue("module")
		id := r.FormValue("id")
		result.Block.ID, err = strconv.Atoi(id)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
		}
		style := r.FormValue("style")
		result.Block.Style, err = strconv.Atoi(style)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
		}

		ret := bll.UpdateModuleBlock(result.Block)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "保存失败"
			break
		}

		result.Result.ErrCode = 0
		result.Result.Reason = "保存成功"
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
