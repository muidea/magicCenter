package dal

import (
	"log"
	"testing"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

func TestCatalog(t *testing.T) {
	log.Print("------------------TestCatalog--------------------")

	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ca := model.CatalogDetail{}
	ca.Name = "testCatalog"
	ca.Creater = 3
	ca.Parent = append(ca.Parent, 10)
	catalog, ret := SaveCatalog(helper, ca)
	if !ret {
		t.Error("SaveCatalog failed")
		return
	}

	ca, found := QueryCatalogByID(helper, catalog.ID)
	if !found {
		t.Error("QueryCatalogByID failed")
	}
	if ca.Creater != 3 {
		t.Error("QueryCatalogByID failed")
	}

	ca.Parent = append(ca.Parent, 8)
	ca.Parent = append(ca.Parent, 9)

	catalog, ret = SaveCatalog(helper, ca)
	if !ret {
		t.Error("SaveCatalog failed")
	}

	ca, found = QueryCatalogByID(helper, catalog.ID)
	if !found {
		t.Error("QueryCatalogByID failed")
	}

	if len(ca.Parent) != 3 {
		t.Error("QueryCatalogByID failed")
	}

	ret = DeleteCatalog(helper, ca.ID)
	if !ret {
		t.Error("DeleteCatalog failed")
	}

	catalogs := QueryAllCatalog(helper)
	if len(catalogs) != 3 {
		t.Error("QueryAllCatalog failed")
	}

	catalogDetails := QueryAllCatalogDetail(helper)
	if len(catalogDetails) != 3 {
		t.Error("QueryAllCatalogDetail")
	}

	catalogs = QueryAvalibleParentCatalog(helper, 10)
	if len(catalogs) != 2 {
		t.Error("QueryAvalibleParentCatalog failed")
	}
}
