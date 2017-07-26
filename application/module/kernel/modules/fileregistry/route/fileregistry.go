package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendFileRegistryRoute 追加FileRegistry路由
func AppendFileRegistryRoute(routes []common.Route, uploadPath string) []common.Route {
	route := createUploadFileRoute(uploadPath)

	routes = append(routes, route)
	return routes
}

func createUploadFileRoute(uploadPath string) common.Route {
	return &uploadFileRoute{uploadPath: uploadPath}
}

type uploadFileRoute struct {
	uploadPath string
}

type uploadFileResult struct {
	common.Result
	FilePath string
}

func (i *uploadFileRoute) Method() string {
	return common.POST
}

func (i *uploadFileRoute) Pattern() string {
	return net.JoinURL(def.URL, "/fileregistry/")
}

func (i *uploadFileRoute) Handler() interface{} {
	return i.uploadFileHandler
}

func (i *uploadFileRoute) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("uploadFileHandler")

	result := uploadFileResult{}
	for true {
		keyName := r.URL.Query().Get("key-name")
		if len(keyName) == 0 {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		err := r.ParseMultipartForm(0)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		dstFile, err := net.MultipartFormFile(r, keyName, i.uploadPath)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "处理出错"
			break
		}

		result.ErrCode = 0
		result.FilePath = dstFile
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
