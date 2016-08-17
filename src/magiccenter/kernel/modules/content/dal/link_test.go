package dal

import (
	"log"
	"magiccenter/kernel/modules/content/model"
	"magiccenter/util/dbhelper"
	"testing"
)

func TestLink(t *testing.T) {
	log.Print("------------------TestLink--------------------")

	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	lnk := model.Link{}
	lnk.Name = "test Link"
	lnk.URL = "test url"
	lnk.Logo = "test link logo"
	lnk.Creater = 10
	lnk.Catalog = append(lnk.Catalog, 8)

	ret := SaveLink(helper, lnk)
	if !ret {
		t.Error("SaveLink failed")
		return
	}

	lnkList := QueryLinkByCatalog(helper, 8)
	if len(lnkList) != 1 {
		t.Error("QueryLinkByCatalog failed")
		return
	}

	curLnk, found := QueryLinkByID(helper, lnkList[0].ID)
	if !found {
		t.Error("QueryLinkByID failed")
		return
	}

	if curLnk.URL != "test url" {
		t.Error("QueryLinkByID failed")
		return
	}

	curLnk.Logo = "logo"
	ret = SaveLink(helper, curLnk)
	if !ret {
		t.Error("SaveLink failed")
		return
	}

	lnkList = QueryAllLink(helper)
	if len(lnkList) != 1 {
		t.Error("QueryAllLink failed")
		return
	}

	ret = DeleteLinkByID(helper, curLnk.ID)
	if !ret {
		t.Error("DeleteLinkByID failed")
	}

	lnkList = QueryAllLink(helper)
	if len(lnkList) != 0 {
		t.Error("DeleteLinkByID failed")
		return
	}
}
