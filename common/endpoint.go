package common

import (
	"muidea.com/magicCommon/model"
)

// EndpointHandler 鉴权处理器
type EndpointHandler interface {
	// 查询所有Endpoint
	QueryAllEndpoint() []model.Endpoint
	// 查询指定的Endpoint
	QueryEndpointByID(id string) (model.Endpoint, bool)
	// 新增Endpoint
	InsertEndpoint(id, name, description string, user []int, status int, accessToken string) (model.Endpoint, bool)
	// 更新Endpoint
	UpdateEndpoint(endpoint model.Endpoint) (model.Endpoint, bool)
	// 删除Endpoint
	DeleteEndpoint(id string) bool

	// 查询授权信息摘要
	GetSummary() model.EndpointSummary
}
