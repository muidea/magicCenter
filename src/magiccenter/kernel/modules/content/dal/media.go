package dal

import (
	"fmt"
	"magiccenter/common"
	"magiccenter/common/model"
	resdal "magiccenter/resource/dal"
)

// QueryAllMedia 查询所有图像
func QueryAllMedia(helper common.DBHelper) []model.MediaDetail {
	mediaList := []model.MediaDetail{}
	sql := fmt.Sprintf(`select id, name, url, type, description, creater from media`)
	helper.Query(sql)

	for helper.Next() {
		media := model.MediaDetail{}
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Type, &media.Desc, &media.Creater)

		mediaList = append(mediaList, media)
	}

	for index, _ := range mediaList {
		media := &mediaList[index]
		ress := resdal.QueryRelativeResource(helper, media.ID, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return mediaList
}

// QueryMediaByCatalog 查询指定分类的图像
func QueryMediaByCatalog(helper common.DBHelper, id int) []model.MediaDetail {
	mediaList := []model.MediaDetail{}

	resList := resdal.QueryReferenceResource(helper, id, model.CATALOG, model.MEDIA)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name, url, type, description, creater from media where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			media := model.MediaDetail{}
			helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Type, &media.Desc, &media.Creater)
			mediaList = append(mediaList, media)
		}
	}

	for _, media := range mediaList {
		ress := resdal.QueryRelativeResource(helper, id, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return mediaList
}

// QueryMediaByRang 查询指定范围的图像
func QueryMediaByRang(helper common.DBHelper, begin int, offset int) []model.MediaDetail {
	mediaList := []model.MediaDetail{}
	sql := fmt.Sprintf(`select id, name, url, type, description, creater from media order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		media := model.MediaDetail{}
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Type, &media.Desc, &media.Creater)

		mediaList = append(mediaList, media)
	}

	for _, media := range mediaList {
		ress := resdal.QueryRelativeResource(helper, media.ID, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return mediaList
}

// QueryMediaByID 查询指定的图像
func QueryMediaByID(helper common.DBHelper, id int) (model.MediaDetail, bool) {
	media := model.MediaDetail{}

	sql := fmt.Sprintf(`select id, name, url, type, description, creater from media where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Type, &media.Desc, &media.Creater)
		result = true
	}

	if result {
		ress := resdal.QueryRelativeResource(helper, id, model.MEDIA)
		for _, r := range ress {
			media.Catalog = append(media.Catalog, r.RId())
		}
	}

	return media, result
}

// DeleteMediaByID 删除图像
func DeleteMediaByID(helper common.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from media where id =%d`, id)
	num, result := helper.Execute(sql)
	if num > 0 && result {
		img := resdal.CreateSimpleRes(id, model.MEDIA, "")
		result = resdal.DeleteResource(helper, img)
	}

	return result
}

// CreateMedia 新建文件
func CreateMedia(helper common.DBHelper, name, url, mediaType, desc string, uID int, catalogs []int) (model.MediaDetail, bool) {
	media := model.MediaDetail{}
	media.Name = name
	media.URL = url
	media.Type = mediaType
	media.Desc = desc
	media.Creater = uID
	media.Catalog = catalogs

	// insert
	sql := fmt.Sprintf(`insert into media (name,url, type, description,creater) values ('%s','%s','%s','%s',%d)`, media.Name, media.URL, media.Type, media.Desc, media.Creater)
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
		res := resdal.CreateSimpleRes(media.ID, model.MEDIA, media.Name)
		for _, c := range media.Catalog {
			ca := resdal.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resdal.SaveResource(helper, res)
	}

	return media, result
}

// SaveMedia 保存文件
func SaveMedia(helper common.DBHelper, media model.MediaDetail) (model.MediaDetail, bool) {
	// modify
	sql := fmt.Sprintf(`update media set name='%s', url ='%s', description='%s', creater=%d where id=%d`, media.Name, media.URL, media.Desc, media.Creater, media.ID)
	num, result := helper.Execute(sql)
	if num != 1 || !result {
		return media, false
	}

	if result {
		res := resdal.CreateSimpleRes(media.ID, model.MEDIA, media.Name)
		for _, c := range media.Catalog {
			ca := resdal.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resdal.SaveResource(helper, res)
	}

	return media, result
}
