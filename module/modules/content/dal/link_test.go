package dal

import (
	"log"
	"testing"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/model"
)

func TestLink(t *testing.T) {
	log.Print("------------------TestLink--------------------")

	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	lnk := model.LinkDetail{}
	lnk.ID = 13
	lnk.Name = "test Link"
	lnk.URL = "test url"
	lnk.Logo = "test link logo"
	lnk.Creater = 10
	lnk.Catalog = append(lnk.Catalog, model.CatalogUnit{ID: 8, Type: "catalog"})

	_, ret := SaveLink(helper, lnk)
	if !ret {
		t.Error("SaveLink failed")
		return
	}

	lnkList := QueryLinkByCatalog(helper, model.CatalogUnit{ID: 8, Type: "catalog"})
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
	_, ret = SaveLink(helper, curLnk)
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
