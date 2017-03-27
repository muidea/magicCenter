package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/account/dal"
)

type groupActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *groupActionHandler) getAllGroups() []model.Group {
	return dal.QueryAllGroup(i.dbhelper)
}

func (i *groupActionHandler) findGroupByID(id int) (model.Group, bool) {
	return dal.QueryGroupByID(i.dbhelper, id)
}

func (i *groupActionHandler) findGroupByName(name string) (model.Group, bool) {
	return dal.QueryGroupByName(i.dbhelper, name)
}

func (i *groupActionHandler) createGroup(name, description string, catalog int) (model.Group, bool) {
	return dal.CreateGroup(i.dbhelper, name, description, catalog)
}

func (i *groupActionHandler) saveGroup(group model.Group) (model.Group, bool) {
	return dal.SaveGroup(i.dbhelper, group)
}

func (i *groupActionHandler) destroyGroup(id int) bool {
	return dal.DeleteGroup(i.dbhelper, id)
}
