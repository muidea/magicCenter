package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllMedia 查询所有图像
func QueryAllMedia(helper dbhelper.DBHelper) []model.MediaDetail {
	mediaList := []model.MediaDetail{}
	sql := fmt.Sprintf(`select id, name, url, description, creater from media`)
	helper.Query(sql)

	for helper.Next() {
		media := model.MediaDetail{}
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Desc, &media.Creater)

		mediaList = append(mediaList, media)
	}

	for index, value := range mediaList {
		media := &mediaList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return mediaList
}

// QueryMediaByCatalog 查询指定分类的图像
func QueryMediaByCatalog(helper dbhelper.DBHelper, id int) []model.MediaDetail {
	mediaList := []model.MediaDetail{}

	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.MEDIA)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name, url, description, creater from media where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			media := model.MediaDetail{}
			helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Desc, &media.Creater)
			mediaList = append(mediaList, media)
		}
	}

	for index, value := range mediaList {
		media := &mediaList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return mediaList
}

// QueryMediaByID 查询指定的图像
func QueryMediaByID(helper dbhelper.DBHelper, id int) (model.MediaDetail, bool) {
	media := model.MediaDetail{}

	sql := fmt.Sprintf(`select id, name, url, description, creater from media where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Desc, &media.Creater)
		result = true
	}

	if result {
		ress := resource.QueryRelativeResource(helper, id, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return media, result
}

// DeleteMediaByID 删除图像
func DeleteMediaByID(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from media where id =%d`, id)
	num, result := helper.Execute(sql)
	if num > 0 && result {
		img := resource.CreateSimpleRes(id, model.MEDIA, "")
		result = resource.DeleteResource(helper, img)
	}

	return result
}

// CreateMedia 新建文件
func CreateMedia(helper dbhelper.DBHelper, name, url, desc string, uID int, catalogs []int) (model.MediaDetail, bool) {
	media := model.MediaDetail{}
	media.Name = name
	media.URL = url
	media.Desc = desc
	media.Creater = uID
	media.Catalog = catalogs

	// insert
	sql := fmt.Sprintf(`insert into media (name,url, description,creater) values ('%s','%s','%s',%d)`, media.Name, media.URL, media.Desc, media.Creater)
	num, result := helper.Execute(sql)
	if num != 1 || !result {
		return media, false
	}

	sql = fmt.Sprintf(`select id from media where url= '%s' and creater=%d`, media.URL, media.Creater)
	helper.Query(sql)
	result = false
	if helper.Next() {
		helper.GetValue(&media.ID)
		result = true
	}

	if result {
		res := resource.CreateSimpleRes(media.ID, model.MEDIA, media.Name)
		for _, c := range media.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return media, result
}

// SaveMedia 保存文件
func SaveMedia(helper dbhelper.DBHelper, media model.MediaDetail) (model.MediaDetail, bool) {
	// modify
	sql := fmt.Sprintf(`update media set name='%s', url ='%s', description='%s', creater=%d where id=%d`, media.Name, media.URL, media.Desc, media.Creater, media.ID)
	num, result := helper.Execute(sql)
	if num != 1 || !result {
		return media, false
	}

	if result {
		res := resource.CreateSimpleRes(media.ID, model.MEDIA, media.Name)
		for _, c := range media.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return media, result
}
