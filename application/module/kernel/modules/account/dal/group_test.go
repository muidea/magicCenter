package dal

import (
	"testing"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCommon/model"
)

func TestGroup(t *testing.T) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	g1 := model.GroupDetail{}
	g1.Name = "test1"
	g1.Description = "Description"
	g1.Catalog = model.AdminGroup

	SaveGroup(helper, g1)

	g2 := model.GroupDetail{}
	g2.Name = "test2"
	g2.Description = "Desc"
	g2.Catalog = model.CommonGroup
	SaveGroup(helper, g2)

	groups := QueryAllGroup(helper)
	if len(groups) < 2 {
		t.Error("SaveGroup failed")
		return
	}

	g11, found := QueryGroupByName(helper, "test1")
	if !found {
		t.Errorf("QueryGroupByName failed, name=%s", "test1")
		return
	}
	if !g11.AdminGroup() {
		t.Error("QueryGroupByName return invalid groupType")
		return
	}

	g111, found := QueryGroupByID(helper, g11.ID)
	if !found {
		t.Errorf("QueryGroupByID failed, id=%d", g11.ID)
		return
	}

	if g111.Name != "test1" {
		t.Error("QueryGroupByID return invalid groupName")
		return
	}

	ok := DeleteGroup(helper, g11.ID)
	if !ok {
		t.Errorf("DeleteGroup failed, id=%d", g11.ID)
		return
	}

	g22, found := QueryGroupByName(helper, "test2")
	if !found {
		t.Errorf("QueryGroupByName failed, name=%s", "test2")
		return
	}

	DeleteGroup(helper, g22.ID)
}
