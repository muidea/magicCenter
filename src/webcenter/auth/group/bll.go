package group

import (
	"fmt"
	"webcenter/modelhelper"
)


type GroupInfo struct {
	Id int
	Name string
	UserCount int
	Catalog int	
}

func QueryAllGroup(model modelhelper.Model) []GroupInfo {
	groupInfoList := []GroupInfo{}
	sql := fmt.Sprintf("select id, name, 0 count, catalog from `group`")
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		group := GroupInfo{}
		model.GetValue(&group.Id, &group.Name, &group.UserCount, &group.Catalog)
		
		groupInfoList = append(groupInfoList, group)
	}

	for index, _ := range groupInfoList {
		info := &groupInfoList[index]
		
		sql = fmt.Sprintf("select count(id) from `user` where `group` like '%d%%,%%' or `group` like '%%,%d%%' or `group` like '%d'", info.Id, info.Id, info.Id)
		if model.Query(sql) {
			for model.Next() {
				model.GetValue(&info.UserCount)
			}
		}
	}

	return groupInfoList
}


