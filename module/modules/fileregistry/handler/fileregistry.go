package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/fileregistry/dal"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
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

func (s *impl) FindFile(fileToken string) (string, model.FileSummary, bool) {
	fileSummary, ok := dal.FindFileSummary(s.dbhelper, fileToken)
	return s.uploadPath, fileSummary, ok
}

func (s *impl) RemoveFile(fileToken string) {
	fileSummary, ok := dal.FindFileSummary(s.dbhelper, fileToken)
	if ok {
		dal.RemoveFileSummary(s.dbhelper, fileToken)

		fullPath := path.Join(fileSummary.FilePath)

		os.Remove(fullPath)

		filePath, _ := path.Split(fullPath)
		_, err := os.Stat(filePath)
		if err == nil {
			os.Remove(filePath)
		}
	}
}

func (s *impl) UploadFile(res http.ResponseWriter, req *http.Request) {
	result := common_def.UploadFileResult{}
	for {
		if req.Method != common.POST {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		keyName := req.URL.Query().Get("key-name")
		if len(keyName) == 0 {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		// max file size
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Printf("ParseMultipartForm failed, err:%s", err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "无效请求数据"
			break
		}

		tempPath := "./"
		dstFile, err := net.MultipartFormFile(req, keyName, tempPath)
		if err != nil {
			log.Printf("net.MultipartFormFile failed, err:%s", err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "上传文件出错"
			break
		}

		fileToken := strings.ToLower(util.RandomAlphanumeric(32))
		_, fileName := path.Split(dstFile)
		finalFilePath := path.Join(s.uploadPath, fileToken)
		_, err = os.Stat(finalFilePath)
		if err != nil {
			err = os.MkdirAll(finalFilePath, os.ModePerm)
		}
		if err != nil {
			log.Printf("Stat file failed, filePath:%s, err:%s", finalFilePath, err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "处理文件出错"
			break
		}

		finalFilePath = path.Join(finalFilePath, fileName)
		err = os.Rename(dstFile, finalFilePath)
		if err != nil {
			log.Printf("rename file failed, rawFile:%s, filePath:%s, err:%s", dstFile, finalFilePath, err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "处理文件出错"
			break
		}

		filePath := path.Join(s.uploadPath, fileToken, fileName)
		fileSummary := model.FileSummary{FileName: fileName, FilePath: filePath}
		fileSummary.FileToken = fileToken
		fileSummary.UploadDate = time.Now().Format("2006-01-02 15:04:05")

		ret := dal.SaveFileSummary(s.dbhelper, fileSummary)
		if ret {
			result.FileToken = fileSummary.FileToken
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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
	result := common_def.DownloadFileResult{}
	for {
		if req.Method != common.GET {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法请求"
			break
		}

		fileToken := req.URL.Query().Get("fileToken")
		if len(fileToken) == 0 {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法请求"
			break
		}
		fileSummary, ok := dal.FindFileSummary(s.dbhelper, fileToken)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "指定文件不存在"
			break
		}

		result.ErrorCode = common_def.Success
		result.RedirectURL = fileSummary.FilePath
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}

func (s *impl) DeleteFile(res http.ResponseWriter, req *http.Request) {
	result := common_def.DeleteFileResult{}
	for true {
		if req.Method != common.DELETE {
			result.ErrorCode = common_def.Failed
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

		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	res.Write(b)
}
