package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCommon/model"
	"muidea.com/magicCenter/application/common/resource"
	"muidea.com/magicCenter/foundation/util"
)

func loadMediaID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from content_media`)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

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

// QueryMedias 查询指定文章
func QueryMedias(helper dbhelper.DBHelper, ids []int) []model.Media {
	mediaList := []model.Media{}

	if len(ids) == 0 {
		return mediaList
	}

	sql := fmt.Sprintf(`select id, name from content_media where id in(%s)`, util.IntArray2Str(ids))
	helper.Query(sql)

	for helper.Next() {
		media := model.Media{}
		helper.GetValue(&media.ID, &media.Name)

		mediaList = append(mediaList, media)
	}

	return mediaList
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
		helper.GetValue(&media.ID, &media.Name, &media.URL, &media.Description, &media.CreateDate, &media.Creater)
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
	result := false
	helper.BeginTransaction()

	for {
		sql := fmt.Sprintf(`delete from content_media where id =%d`, id)
		_, result = helper.Execute(sql)
		if result {
			res, ok := resource.QueryResource(helper, id, model.MEDIA)
			if ok {
				result = resource.DeleteResource(helper, res, true)
			} else {
				result = ok
			}
		}
		break
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return result
}

// CreateMedia 新建文件
func CreateMedia(helper dbhelper.DBHelper, name, url, desc, createDate string, creater int, catalogs []int) (model.Summary, bool) {
	media := model.Summary{Unit: model.Unit{Name: name}, Catalog: catalogs, CreateDate: createDate, Creater: creater}

	id := allocMediaID()
	result := false
	helper.BeginTransaction()

	for {
		// insert
		sql := fmt.Sprintf(`insert into content_media (id, name,url, description, createdate, creater) values (%d, '%s','%s','%s','%s',%d)`, id, name, url, desc, createDate, creater)
		_, result = helper.Execute(sql)
		if !result {
			break
		}

		media.ID = id
		res := resource.CreateSimpleRes(media.ID, model.MEDIA, media.Name, media.CreateDate, media.Creater)
		for _, c := range media.Catalog {
			ca, ok := resource.QueryResource(helper, c, model.CATALOG)
			if ok {
				res.AppendRelative(ca)
			} else {
				result = false
				break
			}
		}

		if result {
			result = resource.CreateResource(helper, res, true)
		}

		break
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return media, result
}

// SaveMedia 保存文件
func SaveMedia(helper dbhelper.DBHelper, media model.MediaDetail) (model.Summary, bool) {
	summary := model.Summary{Unit: model.Unit{ID: media.ID, Name: media.Name}, Catalog: media.Catalog, CreateDate: media.CreateDate, Creater: media.Creater}
	result := false
	helper.BeginTransaction()
	for {
		// modify
		sql := fmt.Sprintf(`update content_media set name='%s', url ='%s', description='%s', createdate='%s', creater=%d where id=%d`, media.Name, media.URL, media.Description, media.CreateDate, media.Creater, media.ID)
		_, result = helper.Execute(sql)
		if result {
			res, ok := resource.QueryResource(helper, media.ID, model.MEDIA)
			if !ok {
				result = false
				break
			}

			res.ResetRelative()
			for _, c := range media.Catalog {
				ca, ok := resource.QueryResource(helper, c, model.CATALOG)
				if ok {
					res.AppendRelative(ca)
				} else {
					result = false
					break
				}
			}

			if result {
				result = resource.SaveResource(helper, res, true)
			}
		}
		break
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return summary, result
}
