package handler

import (
	"encoding/json"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/net"
)

// CreateFileRegistryHandler 新建FileRegistryHandler
func CreateFileRegistryHandler(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) common.FileRegistryHandler {
	uploadPath, _ := cfg.GetOption(model.UploadPath)

	i := impl{uploadPath: uploadPath, sessionRegistry: sessionRegistry}

	return &i
}

type impl struct {
	uploadPath      string
	sessionRegistry common.SessionRegistry
}

type uploadFileResult struct {
	common.Result
	FilePath string
}

type deleteFileResult struct {
	common.Result
}

func (s *impl) FindFile(filePath string) (string, bool) {
	return "", true
}

func (s *impl) UploadFile(res http.ResponseWriter, req *http.Request) {
	result := uploadFileResult{}
	for true {
		if req.Method != common.POST {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		keyName := req.URL.Query().Get("key-name")
		if len(keyName) == 0 {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		err := req.ParseMultipartForm(0)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		dstFile, err := net.MultipartFormFile(req, keyName, s.uploadPath)
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

	res.Write(b)
}

func (s *impl) DeleteFile(res http.ResponseWriter, req *http.Request) {
	result := deleteFileResult{}
	for true {
		if req.Method != common.DELETE {
			result.ErrCode = 1
			result.Reason = "非法请求"
			break
		}

		//_, id := net.SplitRESTAPI(req.URL.Path)

		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}
