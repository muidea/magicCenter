package dal

import (
	"fmt"
	"magiccenter/kernel/modules/content/model"
	resdal "magiccenter/resource/dal"
	"magiccenter/util/modelhelper"
)

// QueryAllImage 查询所有图像
func QueryAllImage(helper modelhelper.Model) []model.ImageDetail {
	imageList := []model.ImageDetail{}
	sql := fmt.Sprintf(`select id, name, url, description, creater from image`)
	helper.Query(sql)

	for helper.Next() {
		image := model.ImageDetail{}
		helper.GetValue(&image.ID, &image.Name, &image.URL, &image.Desc, &image.Creater)

		imageList = append(imageList, image)
	}

	for _, image := range imageList {
		ress := resdal.QueryRelativeResource(helper, image.ID, model.IMAGE)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.ID = r.RId()
			catalog.Name = r.RName()

			image.Catalog = append(image.Catalog, catalog)
		}
	}

	return imageList
}

// QueryImageByCatalog 查询指定分类的图像
func QueryImageByCatalog(helper modelhelper.Model, id int) []model.ImageDetail {
	imageList := []model.ImageDetail{}

	resList := resdal.QueryReferenceResource(helper, id, model.CATALOG, model.IMAGE)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name, url, description, creater from image where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			image := model.ImageDetail{}
			helper.GetValue(&image.ID, &image.Name, &image.URL, &image.Desc, &image.Creater)
			imageList = append(imageList, image)
		}
	}

	for , image := range imageList {
		ress := resdal.QueryRelativeResource(helper, id, model.IMAGE)
		for _, r := range ress {
			image.Catalog = append(image.Catalog, r.RId())
		}
	}

	return imageList
}

// QueryImageByRang 查询指定范围的图像
func QueryImageByRang(helper modelhelper.Model, begin int, offset int) []model.ImageDetail {
	imageList := []model.ImageDetail{}
	sql := fmt.Sprintf(`select id, name, url, description, creater from image order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		image := model.ImageDetail{}
		helper.GetValue(&image.ID, &image.Name, &image.URL, &image.Desc, &image.Creater)

		imageList = append(imageList, image)
	}

	for , image := range imageList {
		ress := resdal.QueryRelativeResource(helper, image.ID, model.IMAGE)
		for _, r := range ress {
			image.Catalog = append(image.Catalog, r.RId())
		}
	}

	return imageList
}

// QueryImageById 查询指定的图像
func QueryImageById(helper modelhelper.Model, id int) (model.ImageDetail, bool) {
	image := model.ImageDetail{}

	sql := fmt.Sprintf(`select id, name, url, description, creater from image where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&image.ID, &image.Name, &image.URL, &image.Desc, &image.Creater)
		result = true
	}

	if result {
		ress := resdal.QueryRelativeResource(helper, id, model.IMAGE)
		for _, r := range ress {
			image.Catalog = append(image.Catalog, r.RId())
		}
	}

	return image, result
}

// DeleteImageById 删除图像
func DeleteImageById(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf(`delete from image where id =%d`, id)
	num, result := helper.Execute(sql)
	if num > 0 && result {
		img := model.ImageDetail{}
		img.ID = id
		result = resdal.DeleteResource(helper, &img)
	}

	return result
}

// SaveImage 保存图像
func SaveImage(helper modelhelper.Model, image model.ImageDetail) bool {
	sql := fmt.Sprintf(`select id from image where id=%d`, image.ID)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into image (name,url,description,creater) values ('%s','%s','%s',%d)`, image.Name, image.URL, image.Desc, image.Creater)
		_, result = helper.Execute(sql)
		sql = fmt.Sprintf(`select id from image where url='%s' and description ='%s' and creater=%d`, image.URL, image.Desc, image.Creater)

		helper.Query(sql)
		result = false
		if helper.Next() {
			helper.GetValue(&image.ID)
			result = true
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update image set name='%s', url ='%s', description='%s', creater=%d where id=%d`, image.Name, image.URL, image.Desc, image.Creater, image.ID)
		_, result = helper.Execute(sql)
	}

	if result {
		result = resdal.SaveResource(helper, &image)
	}

	return result
}
