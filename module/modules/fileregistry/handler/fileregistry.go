package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/module/modules/fileregistry/dal"
	common_def "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/net"
	"github.com/muidea/magicCommon/foundation/util"
	"github.com/muidea/magicCommon/model"
)

// CreateFileRegistryHandler 新建FileRegistryHandler
func CreateFileRegistryHandler(cfg common.Configuration, sessionRegistry common.SessionRegistry, modHub common.ModuleHub) common.FileRegistryHandler {
	staticPath, _ := cfg.GetOption(model.StaticPath)
	uploadPath, _ := cfg.GetOption(model.UploadPath)

	i := impl{uploadPath: path.Join(staticPath, uploadPath), sessionRegistry: sessionRegistry}

	return &i
}

type impl struct {
	uploadPath      string
	sessionRegistry common.SessionRegistry
}

func (s *impl) FindFile(fileToken string) (string, model.FileSummary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	fileSummary, ok := dal.FindFileSummary(dbhelper, fileToken)
	return s.uploadPath, fileSummary, ok
}

func (s *impl) RemoveFile(fileToken string) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	fileSummary, ok := dal.FindFileSummary(dbhelper, fileToken)
	if ok {
		dal.RemoveFileSummary(dbhelper, fileToken)

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
		wd, err := os.Getwd()
		if err != nil {
			log.Printf("get current working path failed, err:%s", err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "上传文件出错"
			break
		}
		finalFilePath := path.Join(s.uploadPath, fileToken)

		finalFullPath := path.Join(wd, finalFilePath)
		_, err = os.Stat(finalFullPath)
		if err != nil {
			err = os.MkdirAll(finalFullPath, os.ModePerm)
		}
		if err != nil {
			log.Printf("Stat file failed, filePath:%s, err:%s", finalFullPath, err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "处理文件出错"
			break
		}

		finalFullPath = path.Join(finalFullPath, fileName)
		cmd := exec.Command("mv", dstFile, finalFullPath)
		err = cmd.Run()
		if err != nil {
			log.Printf("move file failed, rawFile:%s, filePath:%s, err:%s", dstFile, finalFullPath, err.Error())
			result.ErrorCode = common_def.Failed
			result.Reason = "处理文件出错"
			break
		}

		filePath := path.Join(s.uploadPath, fileToken, fileName)
		fileSummary := model.FileSummary{FileName: fileName, FilePath: filePath}
		fileSummary.FileToken = fileToken
		fileSummary.UploadDate = time.Now().Format("2006-01-02 15:04:05")

		dbhelper, err := dbhelper.NewHelper()
		if err != nil {
			panic(err)
		}
		defer dbhelper.Release()
		ret := dal.SaveFileSummary(dbhelper, fileSummary)
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

		dbhelper, err := dbhelper.NewHelper()
		if err != nil {
			panic(err)
		}
		defer dbhelper.Release()
		fileSummary, ok := dal.FindFileSummary(dbhelper, fileToken)
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
		dbhelper, err := dbhelper.NewHelper()
		if err != nil {
			panic(err)
		}
		defer dbhelper.Release()

		fileSummary, ok := dal.FindFileSummary(dbhelper, id)
		if ok {
			dal.RemoveFileSummary(dbhelper, id)
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
