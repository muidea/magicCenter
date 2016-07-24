package dal

import (
	"fmt"
	"magiccenter/kernel/dashboard/module/model"
	contentModel "magiccenter/kernel/modules/content/model"
	"magiccenter/util/modelhelper"
)

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
			helper.GetValue(&block.Id)
			block.Name = name
			block.Tag = tag
			block.Style = style
			block.Owner = owner
			ret = true
		}
	}

	return block, ret
}

func DeleteBlock(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from block where id=%d", id)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

func QueryBlock(helper modelhelper.Model, id int) (model.BlockDetail, bool) {
	block := model.BlockDetail{}
	ret := false

	sql := fmt.Sprintf("select id, name, tag, style, owner from block where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&block.Id, &block.Name, &block.Tag, &block.Style, &block.Owner)

		block.Article = QueryItems(helper, block.Id, contentModel.ARTICLE)
		block.Catalog = QueryItems(helper, block.Id, contentModel.CATALOG)
		block.Link = QueryItems(helper, block.Id, contentModel.LINK)
		ret = true
	}

	return block, ret
}

func QueryBlocks(helper modelhelper.Model, owner string) []model.Block {
	blockList := []model.Block{}
	sql := fmt.Sprintf("select id, name, tag, style, owner from block where owner='%s'", owner)
	helper.Query(sql)

	for helper.Next() {
		b := model.Block{}
		helper.GetValue(&b.Id, &b.Name, &b.Tag, &b.Style, &b.Owner)

		blockList = append(blockList, b)
	}

	return blockList
}

func QueryBlockDetails(helper modelhelper.Model, owner string) []model.BlockDetail {
	blockList := []model.BlockDetail{}
	sql := fmt.Sprintf("select id, name, tag, style, owner from block where owner='%s'", owner)
	helper.Query(sql)

	for helper.Next() {
		b := model.BlockDetail{}
		helper.GetValue(&b.Id, &b.Name, &b.Tag, &b.Style, &b.Owner)

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

func QueryBlockView(helper modelhelper.Model, uri string, id int) (model.BlockView, bool) {
	block := model.BlockView{}
	sql := fmt.Sprintf("select id, name, tag, style, owner from block where id=%d", id)
	helper.Query(sql)

	found := false
	if helper.Next() {
		helper.GetValue(&block.Id, &block.Name, &block.Tag, &block.Style, &block.Owner)
		found = true
	}

	if found {
		block.Items = QueryItemViews(helper, block.Id, uri)
	}

	return block, found
}
