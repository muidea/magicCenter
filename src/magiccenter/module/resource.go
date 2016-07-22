package module

import "net/http"

// Resource 模块资源
type Resource interface {
	// Get 获取资源
	Get(rsp http.ResponseWriter, req *http.Request)
	// Put 更新资源
	Put(rsp http.ResponseWriter, req *http.Request)
	// Post 新增资源
	Post(rsp http.ResponseWriter, req *http.Request)
	// Delete 删除资源
	Delete(rsp http.ResponseWriter, req *http.Request)
}
