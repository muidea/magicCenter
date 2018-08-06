package dal

import (
	"log"
	"testing"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/model"
)

func TestMedia(t *testing.T) {
	log.Print("------------------TestMedia--------------------")

	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	img := model.MediaDetail{}
	img.ID = 32
	img.Name = "test image"
	img.FileToken = "test image fileToken"
	img.Description = "test image descr"
	img.Creater = 10
	img.Catalog = append(img.Catalog, model.CatalogUnit{ID: 10, Type: "catalog"})

	_, ret := SaveMedia(helper, img)
	if !ret {
		t.Error("SaveMedia failed")
		return
	}

	imgList := QueryMediaByCatalog(helper, model.CatalogUnit{ID: 10, Type: "catalog"})
	if len(imgList) != 1 {
		t.Error("QueryMediaByCatalog failed")
		return
	}

	curImg, found := QueryMediaByID(helper, imgList[0].ID)
	if !found {
		t.Error("QueryMediaByID failed")
		return
	}

	if curImg.FileToken != "test image fileToken" {
		t.Error("QueryMediaByID failed")
		return
	}

	curImg.Description = "tttt"
	_, ret = SaveMedia(helper, curImg)
	if !ret {
		t.Error("SaveMedia failed")
		return
	}

	imgList = QueryAllMedia(helper)
	if len(imgList) != 1 {
		t.Error("QueryAllMedia failed")
		return
	}

	ret = DeleteMediaByID(helper, curImg.ID)
	if !ret {
		t.Error("DeleteMediaByID failed")
	}
}
