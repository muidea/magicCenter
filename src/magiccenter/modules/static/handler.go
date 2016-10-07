package static

import (
	"encoding/json"
	"html/template"
	"log"
	"magiccenter/common"
	"net/http"
)

func viewArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	t, err := template.ParseFiles("template/html/blog/view.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(w, nil)
}

// MaintainViewHandler 管理维护视图处理器
func MaintainViewHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("MaintainViewHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/modules/static/maintain.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, nil)
}

// MaintainActionHandler 管理维护信息处理器
func MaintainActionHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("MaintainActionHandler")

	result := common.Result{}

	for {
		err := req.ParseForm()
		if err != nil {
			log.Print("parseform failed")

			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		result.ErrCode = 0
		result.Reason = "更新成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}
