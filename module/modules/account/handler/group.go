package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/account/dal"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/model"
)

type groupActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *groupActionHandler) getGroupCount() int {
	return dal.QueryGroupCount(i.dbhelper)
}

func (i *groupActionHandler) getAllGroup() []model.Group {
	return dal.QueryAllGroup(i.dbhelper)
}

func (i *groupActionHandler) getAllGroupDetail(filter *common_def.PageFilter) []model.GroupDetail {
	return dal.QueryAllGroupDetail(i.dbhelper, filter)
}

func (i *groupActionHandler) getGroups(ids []int) []model.Group {
	return dal.QueryGroups(i.dbhelper, ids)
}

func (i *groupActionHandler) findGroupByID(id int) (model.GroupDetail, bool) {
	return dal.QueryGroupByID(i.dbhelper, id)
}

func (i *groupActionHandler) findSubGroup(id int) []model.Group {
	return dal.QuerySubGroups(i.dbhelper, id)
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
