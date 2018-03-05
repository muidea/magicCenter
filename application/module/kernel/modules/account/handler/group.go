package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/account/dal"
)

type groupActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *groupActionHandler) getAllGroups() []model.GroupDetail {
	return dal.QueryAllGroup(i.dbhelper)
}

func (i *groupActionHandler) getGroups(ids []int) []model.Group {
	return dal.QueryGroups(i.dbhelper, ids)
}

func (i *groupActionHandler) findGroupByID(id int) (model.GroupDetail, bool) {
	return dal.QueryGroupByID(i.dbhelper, id)
}

func (i *groupActionHandler) findGroupByName(name string) (model.GroupDetail, bool) {
	return dal.QueryGroupByName(i.dbhelper, name)
}

func (i *groupActionHandler) createGroup(name, description string, catalog int) (model.GroupDetail, bool) {
	return dal.CreateGroup(i.dbhelper, name, description, catalog)
}

func (i *groupActionHandler) saveGroup(group model.GroupDetail) (model.GroupDetail, bool) {
	return dal.SaveGroup(i.dbhelper, group)
}

func (i *groupActionHandler) destroyGroup(id int) bool {
	return dal.DeleteGroup(i.dbhelper, id)
}
