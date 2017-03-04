package dal

import (
	"fmt"
	"magiccenter/common"
	"magiccenter/common/model"
)

// InsertBlock 新建一条Block
func InsertBlock(helper common.DBHelper, name, tag string, style int, owner string) (model.Block, bool) {
	ret := false

	block := model.Block{}
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

// UpdateBlock 更新一条Block
func UpdateBlock(helper common.DBHelper, block model.Block) (model.Block, bool) {
	ret := false

	sql := fmt.Sprintf("update block set name ='%s', tag = '%s', style= %d where id = %d", block.Name, block.Tag, block.Style, block.ID)
	num, ret := helper.Execute(sql)

	return block, num == 1 && ret
}

// DeleteBlock 删除一条Block记录
func DeleteBlock(helper common.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from block where id=%d", id)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

// QueryBlock 查询一条BlockDetail
func QueryBlock(helper common.DBHelper, id int) (model.Block, bool) {
	block := model.Block{}
	ret := false

	sql := fmt.Sprintf("select id, name, tag, style, owner from block where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&block.ID, &block.Name, &block.Tag, &block.Style, &block.Owner)
		ret = true
	}

	return block, ret
}

// QueryBlockContent 查询一条BlockContent
func QueryBlockContent(helper common.DBHelper, id int) (model.BlockContent, bool) {
	block := model.BlockContent{}
	ret := false

	sql := fmt.Sprintf("select id, name, tag, style, owner from block where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&block.ID, &block.Name, &block.Tag, &block.Style, &block.Owner)

		block.Content = QueryItems(helper, block.ID)
		ret = true
	}

	return block, ret
}

// QueryBlocks 查询指定Module拥有的Block
func QueryBlocks(helper common.DBHelper, owner string) []model.Block {
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

// QueryBlockContents 查询指定类型Block的详情
func QueryBlockContents(helper common.DBHelper, owner string) []model.BlockContent {
	blockList := []model.BlockContent{}
	sql := fmt.Sprintf("select id, name, tag, style, owner from block where owner='%s'", owner)
	helper.Query(sql)

	for helper.Next() {
		b := model.BlockContent{}
		helper.GetValue(&b.ID, &b.Name, &b.Tag, &b.Style, &b.Owner)

		blockList = append(blockList, b)
	}

	// TODO 这里如果直接取可能会存在问题
	for _, b := range blockList {
		b.Content = QueryItems(helper, b.ID)
	}

	return blockList
}
