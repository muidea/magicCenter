package dal

import (
	"log"
	"magiccenter/common/model"
	"magiccenter/util/dbhelper"
	"testing"
)

func TestMedia(t *testing.T) {
	log.Print("------------------TestMedia--------------------")

	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	img := model.MediaDetail{}
	img.Name = "test image"
	img.URL = "test image url"
	img.Desc = "test image descr"
	img.Creater = 10
	img.Catalog = append(img.Catalog, 10)

	ret := SaveMedia(helper, img)
	if !ret {
		t.Error("SaveMedia failed")
		return
	}

	imgList := QueryMediaByCatalog(helper, 10)
	if len(imgList) != 1 {
		t.Error("QueryMediaByCatalog failed")
		return
	}

	curImg, found := QueryMediaByID(helper, imgList[0].ID)
	if !found {
		t.Error("QueryMediaByID failed")
		return
	}

	if curImg.URL != "test image url" {
		t.Error("QueryMediaByID failed")
		return
	}

	curImg.Desc = "tttt"
	ret = SaveMedia(helper, curImg)
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
