package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/resource"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

func loadMediaID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from content_media`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// QueryAllMedia 查询所有图像
func QueryAllMedia(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}

	ress := resource.QueryResourceByType(helper, model.MEDIA)
	for _, v := range ress {
		summary := model.Summary{Unit: model.Unit{ID: v.RId(), Name: v.RName()}, Description: v.RDescription(), Type: v.RType(), CreateDate: v.RCreateDate(), Creater: v.ROwner()}

		for _, r := range v.Relative() {
			summary.Catalog = append(summary.Catalog, r.RId())
		}

		summaryList = append(summaryList, summary)
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
	defer helper.Finish()

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
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, value.Type)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryMediaByID 查询指定的图像
func QueryMediaByID(helper dbhelper.DBHelper, id int) (model.MediaDetail, bool) {
	media := model.MediaDetail{}

	sql := fmt.Sprintf(`select id, name, description, url, createdate, creater from content_media where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&media.ID, &media.Name, &media.Description, &media.URL, &media.CreateDate, &media.Creater)
		result = true
	}
	helper.Finish()

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
func CreateMedia(helper dbhelper.DBHelper, name, description, url, createDate string, expiration, creater int, catalogs []int) (model.Summary, bool) {
	media := model.Summary{Unit: model.Unit{Name: name}, Description: description, Type: model.MEDIA, Catalog: catalogs, CreateDate: createDate, Creater: creater}

	id := allocMediaID()
	result := false
	helper.BeginTransaction()

	for {
		// insert
		sql := fmt.Sprintf(`insert into content_media (id, name, description, url, createdate, creater, expiration) values (%d, '%s','%s','%s','%s',%d,%d)`, id, name, description, url, createDate, creater, expiration)
		_, result = helper.Execute(sql)
		if !result {
			break
		}

		media.ID = id
		res := resource.CreateSimpleRes(media.ID, model.MEDIA, media.Name, media.Description, media.CreateDate, media.Creater)
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
	summary := model.Summary{Unit: model.Unit{ID: media.ID, Name: media.Name}, Type: model.MEDIA, Catalog: media.Catalog, CreateDate: media.CreateDate, Creater: media.Creater}
	result := false
	helper.BeginTransaction()
	for {
		// modify
		sql := fmt.Sprintf(`update content_media set name='%s', description='%s', url ='%s', createdate='%s', creater=%d where id=%d`, media.Name, media.Description, media.URL, media.CreateDate, media.Creater, media.ID)
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
