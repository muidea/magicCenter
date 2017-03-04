package common

import "net/http"

// EndPoint 终端对象
// 所有对外接口都通过EndPoint来操作
type EndPoint interface {
	// Get 获取资源
	Get(rsp http.ResponseWriter, req *http.Request)
	// Put 更新资源
	Put(rsp http.ResponseWriter, req *http.Request)
	// Post 新增资源
	Post(rsp http.ResponseWriter, req *http.Request)
	// Delete 删除资源
	Delete(rsp http.ResponseWriter, req *http.Request)
}
