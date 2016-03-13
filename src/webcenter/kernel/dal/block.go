package dal

import (
	"fmt"
	"webcenter/util/modelhelper"
)

type Block struct {
	Id int
	Name string
	Owner string
	Items []Item
}


func InsertBlock(helper modelhelper.Model, name, owner string) (Block, bool) {
	block := Block{}
	block.Items = []Item{}
	ret := false
	
	sql := fmt.Sprintf("insert into block (name,owner) values('%s','%s')", name, owner)
	_, ret = helper.Execute(sql)
	if ret {
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

func QueryBlock(helper modelhelper.Model, id int) (Block, bool) {
	block := Block{}
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

func QueryBlocks(helper modelhelper.Model, owner string) []Block {	
	blockList := []Block{}
	sql := fmt.Sprintf("select id,name,owner from block where owner='%s'", owner)
	helper.Query(sql)
	
	for helper.Next() {
		b := Block{}
		helper.GetValue(&b.Id, &b.Name, &b.Owner)
		
		blockList = append(blockList, b)
	}
	
	for i, _ := range blockList {
		b := &blockList[i]
		b.Items = QueryItems(helper, b.Id)
	}
	
	return blockList
}



