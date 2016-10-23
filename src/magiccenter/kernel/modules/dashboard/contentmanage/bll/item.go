package bll

import (
	"magiccenter/common/model"
	"magiccenter/kernel/modules/dashboard/contentmanage/dal"
	"magiccenter/system"
)

// GetBlockItem 查询指定Item
func GetBlockItem(id int) (model.Item, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryItem(helper, id)
}

// GetBlockItems 查询属于指定Block的Item
func GetBlockItems(id int) []model.Item {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryItems(helper, id)
}

// AppendBlockItem 追加Item
func AppendBlockItem(item model.Item) (model.Item, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.AddItem(helper, item)
}

// DeleteBlockItem 删除一个Module的Block
func DeleteBlockItem(id int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.RemoveItem(helper, id)
}

// ClearBlockItem 清空指定Block的Item
func ClearBlockItem(owner int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.ClearItems(helper, owner)
}
