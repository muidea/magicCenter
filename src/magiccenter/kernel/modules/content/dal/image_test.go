package dal

import (
	"log"
	"magiccenter/common/model"
	"magiccenter/util/dbhelper"
	"testing"
)

func TestImage(t *testing.T) {
	log.Print("------------------TestImage--------------------")

	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	img := model.ImageDetail{}
	img.Name = "test image"
	img.URL = "test image url"
	img.Desc = "test image descr"
	img.Creater = 10
	img.Catalog = append(img.Catalog, 10)

	ret := SaveImage(helper, img)
	if !ret {
		t.Error("SaveImage failed")
		return
	}

	imgList := QueryImageByCatalog(helper, 10)
	if len(imgList) != 1 {
		t.Error("QueryImageByCatalog failed")
		return
	}

	curImg, found := QueryImageByID(helper, imgList[0].ID)
	if !found {
		t.Error("QueryImageByID failed")
		return
	}

	if curImg.URL != "test image url" {
		t.Error("QueryImageByID failed")
		return
	}

	curImg.Desc = "tttt"
	ret = SaveImage(helper, curImg)
	if !ret {
		t.Error("SaveImage failed")
		return
	}

	imgList = QueryAllImage(helper)
	if len(imgList) != 1 {
		t.Error("QueryAllImage failed")
		return
	}

	ret = DeleteImageByID(helper, curImg.ID)
	if !ret {
		t.Error("DeleteImageByID failed")
	}
}
