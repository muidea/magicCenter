package bll

/*
管理一个模块的Block

对于一个Module来说，拥有的Block时固定的，定义了一个Module时就确定了其拥有的Block

*/

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/dashboard/content/dal"
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

// GetModuleBlockDetail 查询指定Module拥有的Block
func GetModuleBlockDetail(id int) (model.BlockDetail, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryBlockDetail(helper, id)
}

// AppendModuleBlock 增加一个Module的Block
func AppendModuleBlock(block model.Block) (model.Block, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.InsertBlock(helper, block)
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
func UpdateModuleBlock(block model.Block) bool {
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
