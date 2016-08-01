package dal

import (
	"fmt"
	"magiccenter/kernel/dashboard/module/model"
	"magiccenter/util/modelhelper"
)

// InsertBlock 新建一条Block记录
func InsertBlock(helper modelhelper.Model, name, tag string, style int, owner string) (model.Block, bool) {
	block := model.Block{}
	ret := false

	sql := fmt.Sprintf("insert into block (name, tag, style, owner) values('%s','%s', %d, '%s')", name, tag, style, owner)
	num, ret := helper.Execute(sql)
	if num == 1 && ret {
		ret = false
		sql = fmt.Sprintf("select id from block where name='%s' and owner='%s'", name, owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&block.ID)
			block.Name = name
			block.Tag = tag
			block.Style = style
			block.Owner = owner
			ret = true
		}
	}

	return block, ret
}

// DeleteBlock 删除一条Block记录
func DeleteBlock(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from block where id=%d", id)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

// QueryBlock 查询一条Block
func QueryBlock(helper modelhelper.Model, id int) (model.BlockDetail, bool) {
	block := model.BlockDetail{}
	ret := false

	sql := fmt.Sprintf("select id, name, tag, style, owner from block where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&block.ID, &block.Name, &block.Tag, &block.Style, &block.Owner)

		block.Items = QueryItems(helper, block.ID)
		ret = true
	}

	return block, ret
}

// QueryBlocks 查询指定类型的Block
func QueryBlocks(helper modelhelper.Model, owner string) []model.Block {
	blockList := []model.Block{}
	sql := fmt.Sprintf("select id, name, tag, style, owner from block where owner='%s'", owner)
	helper.Query(sql)

	for helper.Next() {
		b := model.Block{}
		helper.GetValue(&b.ID, &b.Name, &b.Tag, &b.Style, &b.Owner)

		blockList = append(blockList, b)
	}

	return blockList
}

// QueryBlockDetails 查询指定类型Block的详情
func QueryBlockDetails(helper modelhelper.Model, owner string) []model.BlockDetail {
	blockList := []model.BlockDetail{}
	sql := fmt.Sprintf("select id, name, tag, style, owner from block where owner='%s'", owner)
	helper.Query(sql)

	for helper.Next() {
		b := model.BlockDetail{}
		helper.GetValue(&b.ID, &b.Name, &b.Tag, &b.Style, &b.Owner)

		blockList = append(blockList, b)
	}

	// TODO 这里如果直接取可能会存在问题
	for _, b := range blockList {
		//b := &blockList[i]
		b.Items = QueryItems(helper, b.ID)
	}

	return blockList
}
