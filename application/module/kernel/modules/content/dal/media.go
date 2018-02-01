package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllMedia 查询所有图像
func QueryAllMedia(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}
	sql := fmt.Sprintf(`select id, name, createdate,creater from content_media`)
	helper.Query(sql)

	for helper.Next() {
		media := model.Summary{}
		helper.GetValue(&media.ID, &media.Name, &media.CreateDate, &media.Creater)

		summaryList = append(summaryList, media)
	}

	for index, value := range summaryList {
		media := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryMediaByCatalog 查询指定分类的图像
func QueryMediaByCatalog(helper dbhelper.DBHelper, id int) []model.Summary {
	summaryList := []model.Summary{}

	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.MEDIA)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name,createdate,creater from content_media where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			media := model.Summary{}
			helper.GetValue(&media.ID, &media.Name, &media.CreateDate, &media.Creater)
			summaryList = append(summaryList, media)
		}
	}

	for index, value := range summaryList {
		media := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryMediaByID 查询指定的图像
func QueryMediaByID(helper dbhelper.DBHelper, id int) (model.MediaDetail, bool) {
	media := model.MediaDetail{}

	sql := fmt.Sprintf(`select id, name, url, description,createdate, creater from content_media where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Desc, &media.CreateDate, &media.Creater)
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
	sql := fmt.Sprintf(`delete from content_media where id =%d`, id)
	num, result := helper.Execute(sql)
	if num > 0 && result {
		img := resource.CreateSimpleRes(id, model.MEDIA, "", "", -1)
		result = resource.DeleteResource(helper, img)
	}

	return result
}

// CreateMedia 新建文件
func CreateMedia(helper dbhelper.DBHelper, name, url, desc, createdate string, uID int, catalogs []int) (model.Summary, bool) {
	media := model.Summary{}
	media.Name = name
	media.Catalog = catalogs
	media.CreateDate = createdate
	media.Creater = uID

	// insert
	sql := fmt.Sprintf(`insert into content_media (name,url, description, createdate, creater) values ('%s','%s','%s','%s',%d)`, name, url, desc, createdate, uID)
	num, result := helper.Execute(sql)
	if num != 1 || !result {
		return media, false
	}

	sql = fmt.Sprintf(`select id from content_media where url= '%s' and creater=%d`, url, uID)
	helper.Query(sql)
	result = false
	if helper.Next() {
		helper.GetValue(&media.ID)
		result = true
	}

	if result {
		res := resource.CreateSimpleRes(media.ID, model.MEDIA, media.Name, media.CreateDate, media.Creater)
		for _, c := range media.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "", "", -1)
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return media, result
}

// SaveMedia 保存文件
func SaveMedia(helper dbhelper.DBHelper, media model.MediaDetail) (model.Summary, bool) {
	// modify
	sql := fmt.Sprintf(`update content_media set name='%s', url ='%s', description='%s', createdate='%s', creater=%d where id=%d`, media.Name, media.URL, media.Desc, media.CreateDate, media.Creater, media.ID)
	num, result := helper.Execute(sql)
	if num == 1 && result {
		res := resource.CreateSimpleRes(media.ID, model.MEDIA, media.Name, media.CreateDate, media.Creater)
		for _, c := range media.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "", "", -1)
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	} else {
		result = false
	}

	return model.Summary{ID: media.ID, Name: media.Name, Catalog: media.Catalog, CreateDate: media.CreateDate, Creater: media.Creater}, result
}
