package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// FileRegistryHandler 文件管理处理器
type FileRegistryHandler interface {
	// FindFile 查找指定文件
	FindFile(accessToken string) (string, model.FileSummary, bool)
	// UploadFile 上传文件
	UploadFile(res http.ResponseWriter, req *http.Request)
	// DownloadFile 下载文件
	DownloadFile(res http.ResponseWriter, req *http.Request)
	// DeleteFile 删除文件
	DeleteFile(res http.ResponseWriter, req *http.Request)
}
