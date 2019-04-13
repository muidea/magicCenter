package handler

import (
	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/module/modules/account/dal"
	common_util "github.com/muidea/magicCommon/foundation/util"
	"github.com/muidea/magicCommon/model"
)

type groupActionHandler struct {
}

func (i *groupActionHandler) getGroupCount() int {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryGroupCount(dbhelper)
}

func (i *groupActionHandler) getAllGroup() []model.Group {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllGroup(dbhelper)
}

func (i *groupActionHandler) getAllGroupDetail(filter *common_util.PageFilter) []model.GroupDetail {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllGroupDetail(dbhelper, filter)
}

func (i *groupActionHandler) getGroups(ids []int) []model.Group {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryGroups(dbhelper, ids)
}

func (i *groupActionHandler) findGroupByID(id int) (model.GroupDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryGroupByID(dbhelper, id)
}

func (i *groupActionHandler) findSubGroup(id int) []model.Group {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QuerySubGroups(dbhelper, id)
}

func (i *groupActionHandler) findGroupByName(name string) (model.GroupDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryGroupByName(dbhelper, name)
}

func (i *groupActionHandler) createGroup(name, description string, catalog int) (model.GroupDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.CreateGroup(dbhelper, name, description, catalog)
}

func (i *groupActionHandler) saveGroup(group model.GroupDetail) (model.GroupDetail, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveGroup(dbhelper, group)
}

func (i *groupActionHandler) destroyGroup(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteGroup(dbhelper, id)
}
