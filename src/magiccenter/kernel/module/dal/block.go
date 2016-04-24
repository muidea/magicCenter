package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/module/model"
	contentModel "magiccenter/kernel/content/model"
)

func InsertBlock(helper modelhelper.Model, name string, style int, owner string) (model.Block, bool) {
	block := model.Block{}
	ret := false
	
	sql := fmt.Sprintf("insert into block (name, style, owner) values('%s', %d, '%s')", name, style, owner)
	num, ret := helper.Execute(sql)
	if num == 1 && ret {
		ret = false
		sql = fmt.Sprintf("select id from block where name='%s' and owner='%s'", name, owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&block.Id)
			block.Name = name
			block.Style = style
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
	
	sql := fmt.Sprintf("select id, name, style, owner from block where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&block.Id, &block.Name, &block.Style, &block.Owner)
		
		block.Article = QueryItems(helper, block.Id, contentModel.ARTICLE)
		block.Catalog = QueryItems(helper, block.Id, contentModel.CATALOG)
		block.Link = QueryItems(helper, block.Id, contentModel.LINK)		
		ret = true
	}
	
	return block, ret
}

func QueryBlocks(helper modelhelper.Model, owner string) []model.Block {	
	blockList := []model.Block{}
	sql := fmt.Sprintf("select id, name, style, owner from block where owner='%s'", owner)
	helper.Query(sql)
	
	for helper.Next() {
		b := model.Block{}
		helper.GetValue(&b.Id, &b.Name, &b.Style, &b.Owner)
		
		blockList = append(blockList, b)
	}
		
	return blockList
}


func QueryBlockDetails(helper modelhelper.Model, owner string) []model.BlockDetail {	
	blockList := []model.BlockDetail{}
	sql := fmt.Sprintf("select id, name, style, owner from block where owner='%s'", owner)
	helper.Query(sql)
	
	for helper.Next() {
		b := model.BlockDetail{}
		helper.GetValue(&b.Id, &b.Name, &b.Style, &b.Owner)
		
		blockList = append(blockList, b)
	}
		
	for i, _ := range blockList {
		b := &blockList[i]
		b.Article = QueryItems(helper, contentModel.ARTICLE, b.Id)
		b.Catalog = QueryItems(helper, contentModel.CATALOG, b.Id)
		b.Link = QueryItems(helper, contentModel.LINK, b.Id)
	}
	
	return blockList
}



