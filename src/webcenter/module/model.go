package module

import (
	"fmt"
	"webcenter/util/modelhelper"
)

func addPageBlock(helper modelhelper.Model, url string, block int) bool {
	sql := fmt.Sprintf("insert into page_block (url,block) values('%s',%d)", url, block)	
	return helper.Execute(sql)
}

func removePageBlock(helper modelhelper.Model, url string, block int) {
	sql := fmt.Sprintf("delete page_block where url='%s' and block=%d", url, block)	
	helper.Execute(sql)
}

func queryPageBlock(helper modelhelper.Model, url string) []Block {
	blockList := []Block{}
	sql := fmt.Sprintf("select id,name,owner from module_block where id in (select id from page_block where url='%s')", url)
	helper.Query(sql)
	
	bList := []*block{}
	for helper.Next() {
		b := &block{}
		helper.GetValue(&b.id, &b.name, &b.owner)
		
		bList = append(bList, b)
	}
	
	for i, _ := range bList {
		b := bList[i]
		b.items = queryBlockItems(helper, b.id)
		
		blockList = append(blockList, b)
	}
	
	return blockList
}

func addBlockItem(helper modelhelper.Model, name,url string, owner int) bool {
	sql := fmt.Sprintf("insert into block_item (name,url,owner) values('%s','%s',%d)", name, url,owner)
	return helper.Execute(sql)
}

func removeBlockItem(helper modelhelper.Model, id int) {
	sql := fmt.Sprintf("delete block_item where id=%d", id)
	helper.Execute(sql)
}

func queryBlockItems(helper modelhelper.Model, owner int) []Item {
	itemList := []Item{}
	
	sql := fmt.Sprintf("select id,name,url,owner from block_item where owner=%d", owner)
	helper.Query(sql)
	for helper.Next() {
		i := &item{}
		helper.GetValue(&i.id, &i.name, &i.url, &i.owner)
		
		itemList = append(itemList, i)
	}
	
	return itemList
}

func insertModuleBlock(helper modelhelper.Model, name,owner string) (Block, bool) {
	b := &block{}
	
	result := false
	sql := fmt.Sprintf("insert into module_block (name,owner) values('%s','%s')", name, owner)
	if helper.Execute(sql) {
		sql = fmt.Sprintf("select id from module_block where name='%s' and owner='%s'", name, owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&b.id)
			b.name = name
			b.owner = owner
			result = true
		}
	}
	
	return b,result
}

func deleteModuleBlock(helper modelhelper.Model, id int) {
	sql := fmt.Sprintf("delete from module_block where id=%d", id)
	helper.Execute(sql)
}

func queryModuleBlocks(helper modelhelper.Model, owner string) []Block {	
	blockList := []Block{}
	sql := fmt.Sprintf("select id,name,owner from module_block where owner='%s'", owner)
	helper.Query(sql)
	
	bList := []*block{}
	for helper.Next() {
		b := &block{}
		helper.GetValue(&b.id, &b.name, &b.owner)
		
		bList = append(bList, b)
	}
	
	for i, _ := range bList {
		b := bList[i]
		b.items = queryBlockItems(helper, b.id)
		
		blockList = append(blockList, b)
	}
	
	return blockList
}

func deleteModule(helper modelhelper.Model, id string) {
	helper.BeginTransaction()
	
	sql := fmt.Sprintf("delete from module_block where owner='%s'", id)
	if helper.Execute(sql) {
		sql = fmt.Sprintf("delete from module where id='%s'", id)
		if helper.Execute(sql) {
			helper.Commit()
		} else {
			helper.Rollback()			
		}
	} else {
		helper.Rollback()
	}
	
}

func queryModule(helper modelhelper.Model, id string) (Module, bool) {
	m := &module{}
	sql := fmt.Sprintf("select id, name, description, enableflag, defaultflag, styleflag from module where id='%s'", id)	
	helper.Query(sql)
	
	result := false
	if helper.Next() {
		helper.GetValue(&m.id, &m.name, &m.description, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		result = true
	}
	
	if result {
		m.blocks = queryModuleBlocks(helper, m.id)
	}
		
	return m, result	
}

func queryAllModule(helper modelhelper.Model) []Module {
	moduleList := []Module{}
	
	sql := fmt.Sprintf("select id, name, description, enableflag, defaultflag, styleflag from module order by styleflag")	
	helper.Query(sql)
	
	mList := []*module{}
	for helper.Next() {
		m := &module{}
		helper.GetValue(&m.id, &m.name, &m.description, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		
		mList = append(mList, m)
	}
	
	for i, _ := range mList {
		m := mList[i]
		m.blocks = queryModuleBlocks(helper, m.id)
		
		moduleList = append(moduleList, m)
	}
	
	return moduleList
}

func saveModule(helper modelhelper.Model, m Module) bool {
	result := false
	_, found := queryModule(helper, m.ID())
	if found {
		sql := fmt.Sprintf("update module set Name ='%s', Description ='%s', enableflag =%d, defaultflag =%d where Id='%s'", m.Name(), m.Description(), m.EnableStatus(), m.DefaultStatus(), m.ID())
		result = helper.Execute(sql)
	} else {
		sql := fmt.Sprintf("insert into module(id, name, description, enableflag, defaultflag, styleflag) values ('%s','%s','%s',%d,%d,%d)", m.ID(), m.Name(), m.Description(), m.EnableStatus(), m.DefaultStatus(), m.StyleFlag())
		result = helper.Execute(sql)
	}
	
	return result
}






