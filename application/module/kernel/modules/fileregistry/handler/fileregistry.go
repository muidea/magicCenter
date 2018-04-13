package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCommon/model"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry/dal"
	"muidea.com/magicCenter/foundation/net"
	"muidea.com/magicCenter/foundation/util"
)

// CreateFileRegistryHandler 新建FileRegistryHandler
func CreateFileRegistryHandler(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) common.FileRegistryHandler {
	staticPath, _ := cfg.GetOption(model.StaticPath)
	uploadPath, _ := cfg.GetOption(model.UploadPath)

	dbhelper, _ := dbhelper.NewHelper()

	i := impl{dbhelper: dbhelper, uploadPath: path.Join(staticPath, uploadPath), sessionRegistry: sessionRegistry}

	return &i
}

type impl struct {
	dbhelper        dbhelper.DBHelper
	uploadPath      string
	sessionRegistry common.SessionRegistry
}

type uploadFileResult struct {
	common.Result
	AccessToken string `json:"accessToken"`
}

type downloadFileResult struct {
	common.Result
	RedirectURL string `json:"redirectUrl"`
}

type deleteFileResult struct {
	common.Result
}

func (s *impl) FindFile(accessToken string) (string, model.FileSummary, bool) {
	fileSummary, ok := dal.FindFileSummary(s.dbhelper, accessToken)
	return s.uploadPath, fileSummary, ok
}

func (s *impl) UploadFile(res http.ResponseWriter, req *http.Request) {
	result := uploadFileResult{}
	for true {
		if req.Method != common.POST {
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		keyName := req.URL.Query().Get("key-name")
		if len(keyName) == 0 {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		err := req.ParseMultipartForm(0)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "无效请求数据"
			break
		}

		tempPath := "./"
		dstFile, err := net.MultipartFormFile(req, keyName, tempPath)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "上传文件出错"
			break
		}

		accessToken := strings.ToLower(util.RandomAlphanumeric(32))
		_, fileName := path.Split(dstFile)
		finalFilePath := path.Join(s.uploadPath, accessToken)
		_, err = os.Stat(finalFilePath)
		if err != nil {
			err = os.MkdirAll(finalFilePath, os.ModePerm)
		}
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "处理文件出错"
			break
		}

		finalFilePath = path.Join(finalFilePath, fileName)
		err = os.Rename(dstFile, finalFilePath)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "处理文件出错"
			break
		}

		filePath := path.Join(accessToken, fileName)
		fileSummary := model.FileSummary{FileName: fileName, FilePath: filePath}
		fileSummary.AccessToken = accessToken
		fileSummary.UploadDate = time.Now().Format("2006-01-02 15:04:05")

		ret := dal.SaveFileSummary(s.dbhelper, fileSummary)
		if ret {
			result.AccessToken = fileSummary.AccessToken
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
			result.Reason = "保存文件信息失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}

func (s *impl) DownloadFile(res http.ResponseWriter, req *http.Request) {
	result := downloadFileResult{}
	for true {
		if req.Method != common.GET {
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		_, id := net.SplitRESTAPI(req.URL.Path)
		_, ok := dal.FindFileSummary(s.dbhelper, id)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "指定文件不存在"
			break
		}

		result.ErrorCode = common.Success
		result.RedirectURL = fmt.Sprintf("/static/?source=%s", id)
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
			result.ErrorCode = common.Failed
			result.Reason = "非法请求"
			break
		}

		_, id := net.SplitRESTAPI(req.URL.Path)
		fileSummary, ok := dal.FindFileSummary(s.dbhelper, id)
		if ok {
			dal.RemoveFileSummary(s.dbhelper, id)
			finalFilePath := path.Join(s.uploadPath, fileSummary.FilePath)
			_, err := os.Stat(finalFilePath)
			if err == nil {
				os.Remove(finalFilePath)

				filePath, _ := path.Split(finalFilePath)
				os.Remove(filePath)
			}
		}

		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}
