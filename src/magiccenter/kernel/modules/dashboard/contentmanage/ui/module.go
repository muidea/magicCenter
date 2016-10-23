package ui

/*

管理一个Module的内容信息，获取Module拥有的Page 已经支持的Block
*/

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/dashboard/contentmanage/bll"
	"magiccenter/module"
	"net/http"

	"muidea.com/util"
)

// ModuleView Module视图
// Blocks Module定义的功能块列表
// Pages Module定义的页面URL
type ModuleView struct {
	Result common.Result
	Blocks []model.Block
	Pages  []string
}

// ModuleViewHandler Content管理视图处理器
func ModuleViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModuleViewHandler")

	result := ModuleView{}
	params := util.SplitParam(r.URL.RawQuery)
	for true {
		owner, found := params["module"]
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		contentModule, found := module.FindModule(owner)
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		url := contentModule.URL()
		routes := contentModule.Routes()
		for _, rt := range routes {
			pageURL := util.JoinURL(url, rt.Pattern())
			result.Pages = append(result.Pages, pageURL)
		}

		result.Blocks = bll.QueryModuleBlocks(owner)

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
