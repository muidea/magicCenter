package bll

import (
	"magiccenter/kernel/modules/content/model"
	"magiccenter/util/modelhelper"
)

// QueryAllImage 查询全部图像
func QueryAllImage() []model.ImageDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAllImage(helper)
}

// QueryImageByID 查询指定图像
func QueryImageByID(id int) (model.ImageDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryImageById(helper, id)
}

// DeleteImageByID 删除图像
func DeleteImageByID(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteImageById(helper, id)
}

// QueryImageByCatalog 查询指定分类的图像
func QueryImageByCatalog(id int) []model.ImageDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryImageByCatalog(helper, id)
}

// QueryImageByRang 查询指定范围图像
func QueryImageByRang(begin, offset int) []model.ImageDetail {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryImageByRang(helper, begin, offset)
}

// SaveImage 保存图像
func SaveImage(id int, name, url, desc string, uID int, catalogs []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	image := model.ImageDetail{}
	image.Id = id
	image.Name = name
	image.Url = url
	image.Desc = desc
	image.Creater = uID
	image.Catalog = catalogs

	return dal.SaveImage(helper, image)
}
