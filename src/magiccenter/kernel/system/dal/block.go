package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/system/model"
)

func InsertBlock(helper modelhelper.Model, name, owner string) (model.Block, bool) {
	block := model.Block{}
	ret := false
	
	sql := fmt.Sprintf("insert into block (name,owner) values('%s','%s')", name, owner)
	num, ret := helper.Execute(sql)
	if num == 1 && ret {
		ret = false
		sql = fmt.Sprintf("select id from block where name='%s' and owner='%s'", name, owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&block.Id)
			block.Name = name
			block.Owner = owner
			ret = true
		}
	}
	
	return block,ret
}

func DeleteBlock(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from block where id=%d", id)
	num ,ret := helper.Execute(sql)
	return num == 1 && ret
}

func QueryBlock(helper modelhelper.Model, id int) (model.BlockDetail, bool) {
	block := model.BlockDetail{}
	ret := false
	
	sql := fmt.Sprintf("select id,name,owner from block where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&block.Id, &block.Name, &block.Owner)
		
		block.Items = QueryItems(helper, id)
		ret = true
	}
	
	return block, ret
}

func QueryBlocks(helper modelhelper.Model, owner string) []model.Block {	
	blockList := []model.Block{}
	sql := fmt.Sprintf("select id,name,owner from block where owner='%s'", owner)
	helper.Query(sql)
	
	for helper.Next() {
		b := model.Block{}
		helper.GetValue(&b.Id, &b.Name, &b.Owner)
		
		blockList = append(blockList, b)
	}
		
	return blockList
}



