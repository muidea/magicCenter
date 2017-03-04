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

// SingleItem 单个Item项
type SingleItem struct {
	Result common.Result
	Item   model.Item
}

// GetItemActionHandler 查询Block拥有的Item信息
func GetItemActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetItemActionHandler")

	result := SingleItem{}
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

		result.Item, found = bll.GetBlockItem(aid)
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

// PostItemActionHandler 添加Item
func PostItemActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostItemActionHandler")

	result := SingleItem{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		rid := r.FormValue("rid")
		result.Item.Rid, err = strconv.Atoi(rid)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}
		result.Item.Rtype = r.FormValue("type")
		owner := r.FormValue("owner")
		result.Item.Owner, err = strconv.Atoi(owner)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		ret := false
		result.Item, ret = bll.AppendBlockItem(result.Item)
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

// DeleteItemActionHandler 删除Item
func DeleteItemActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteItemActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		id, found := params["id"]
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

		ret := bll.DeleteBlockItem(aid)
		if !ret {
			result.ErrCode = 1
			result.Reason = "删除出错"
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

// PutItemActionHandler 更新Item
func PutItemActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutItemActionHandler")

	result := common.Result{}
	result.ErrCode = 1
	result.Reason = "操作不支持"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
