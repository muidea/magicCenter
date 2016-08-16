package dal

import (
	"magiccenter/util/dbhelper"
	"testing"
)

func TestResource(t *testing.T) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	res := CreateSimpleRes(0, "test", "test")

	catalog := CreateSimpleRes(9, "catalog", "")

	res.AppendRelative(catalog)

	ret := SaveResource(helper, res)
	if !ret {
		t.Errorf("Save resource failed")
		return
	}

	res1, found := QueryResource(helper, 0, "test")
	if !found {
		t.Error("Query resource failed")
		return
	}

	if res1.RName() != "test" {
		t.Error("invalid resource name")
		return
	}

	if res1.URL() != "test/id=0" {
		t.Error("invalid URL func")
		return
	}

	rres := res1.Relative()
	if len(rres) != 1 {
		t.Error("fetch relative catalog failed")
		return
	}

	ret = DeleteResource(helper, res)
	if !ret {
		t.Error("Delete resource failed")
		return
	}
}
