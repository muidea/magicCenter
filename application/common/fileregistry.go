package common

import "net/http"

// FileRegistryHandler 文件管理处理器
type FileRegistryHandler interface {
	// FindFile 查找指定文件
	FindFile(filePath string) (string, bool)
	// UploadFile 上传文件
	UploadFile(res http.ResponseWriter, req *http.Request)
	// DeleteFile 删除文件
	DeleteFile(res http.ResponseWriter, req *http.Request)
}
