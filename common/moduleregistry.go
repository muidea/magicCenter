package common

import "muidea.com/magicCommon/model"

// ModuleRegistryHandler 模块仓库管理接口
type ModuleRegistryHandler interface {
	GetModuleList() []model.Module

	GetModuleDetailList() []model.ModuleDetail

	QueryModuleByID(id string) (model.ModuleDetail, bool)
}
