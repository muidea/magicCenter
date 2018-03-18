package resource

import (
	"testing"

	"muidea.com/magicCenter/application/common/dbhelper"
)

func TestResource(t *testing.T) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	catalog1 := CreateSimpleRes(9, "catalog", "ca1", "2018-03-10 00:00:00", 0)
	catalog2 := CreateSimpleRes(10, "catalog", "ca2", "2018-03-10 00:00:00", 0)
	catalog3 := CreateSimpleRes(11, "catalog", "ca3", "2018-03-10 00:00:00", 0)

	CreateResource(helper, catalog1)
	CreateResource(helper, catalog2)
	CreateResource(helper, catalog3)

	res := CreateSimpleRes(0, "test", "test", "2018-03-10 00:00:00", 0)

	res.AppendRelative(catalog1)
	res.AppendRelative(catalog2)

	ret := CreateResource(helper, res)
	if !ret {
		t.Errorf("Create resource failed")
		return
	}

	res1, found := QueryResource(helper, res.RId(), res.RType())
	if !found {
		t.Error("Query resource failed")
		return
	}

	if res1.RName() != "test" {
		t.Error("invalid resource name")
		return
	}

	rres := res1.Relative()
	if len(rres) != 2 {
		t.Error("fetch relative catalog failed")
		return
	}

	res1.AppendRelative(catalog3)
	ret = SaveResource(helper, res1)
	if !ret {
		t.Error("Save resouce failed")
	}

	DeleteResource(helper, catalog1)
	DeleteResource(helper, catalog2)
	DeleteResource(helper, catalog3)
	ret = DeleteResource(helper, res)
	if !ret {
		t.Error("Delete resource failed")
		return
	}
}
