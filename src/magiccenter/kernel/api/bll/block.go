package bll

/*
管理一个模块的Block

对于一个Module来说，拥有的Block时固定的，定义了一个Module时就确定了其拥有的Block
*/

import (
	"magiccenter/common/model"
	"magiccenter/kernel/api/dal"
	"magiccenter/system"
)

// GetModuleBlock 查询指定Module拥有的Block
func GetModuleBlock(id int) (model.Block, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryBlock(helper, id)
}

// GetModuleBlockContent 查询指定Module拥有的Block
func GetModuleBlockContent(id int) (model.BlockContent, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryBlockContent(helper, id)
}

// InsertModuleBlock 增加一个Module的Block
func InsertModuleBlock(name, tag string, style int, owner string) (model.Block, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.InsertBlock(helper, name, tag, style, owner)
}

// DeleteModuleBlock 删除一个Module的Block
func DeleteModuleBlock(id int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteBlock(helper, id)
}

// UpdateModuleBlock 更新指定的Block
func UpdateModuleBlock(block model.Block) (model.Block, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.UpdateBlock(helper, block)
}

// QueryModuleBlocks 查询一个Module拥有的Block
func QueryModuleBlocks(owner string) []model.Block {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryBlocks(helper, owner)
}
