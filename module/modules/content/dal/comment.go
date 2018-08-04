package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/resource"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

func loadCommentID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from content_comment`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// QueryCommentByCatalog 查询指定分类下的Comment
func QueryCommentByCatalog(helper dbhelper.DBHelper, idValue int, typeValue string) []model.Summary {
	summaryList := []model.Summary{}

	resList := resource.QueryReferenceResource(helper, idValue, typeValue, model.COMMENT)
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

// QueryCommentByID 查询指定Comment
func QueryCommentByID(helper dbhelper.DBHelper, id int) (model.CommentDetail, bool) {
	comment := model.CommentDetail{}
	sql := fmt.Sprintf(`select id, subject, content, createdate, creater from content_comment where id =%d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&comment.ID, &comment.Subject, &comment.Content, &comment.CreateDate, &comment.Creater)
		result = true
	}
	helper.Finish()

	if result {
		ress := resource.QueryRelativeResource(helper, comment.ID, model.COMMENT)
		for _, r := range ress {
			comment.Catalog = append(comment.Catalog, r.RId())
		}

		// 如果Catalog没有父分类，则认为其父分类为BuildContentCatalog
		// unexpected
		//if len(comment.Catalog) == 0 {
		//comment.Catalog = append(comment.Catalog, common_const.SystemContentCatalog.ID)
		//}
	}

	return comment, result
}

// DisableCommentByID 禁止指定Comment
func DisableCommentByID(helper dbhelper.DBHelper, id int) bool {
	result := false

	sql := fmt.Sprintf(`update content_comment set flag=1 where id =%d`, id)
	_, result = helper.Execute(sql)

	return result
}

// DeleteCommentByID 删除指定Comment
func DeleteCommentByID(helper dbhelper.DBHelper, id int) bool {
	result := false
	helper.BeginTransaction()

	for {
		sql := fmt.Sprintf(`delete from content_comment where id =%d`, id)
		_, result = helper.Execute(sql)
		if result {
			res, ok := resource.QueryResourceByID(helper, id, model.COMMENT)
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

// CreateComment 新建Comment
func CreateComment(helper dbhelper.DBHelper, subject, content, createDate string, creater int, catalogs []int) (model.Summary, bool) {
	desc := util.ExtractSummary(content)
	cmt := model.Summary{Unit: model.Unit{Name: subject}, Description: desc, Type: model.COMMENT, Catalog: catalogs, CreateDate: createDate, Creater: creater}

	id := allocCommentID()
	result := false
	helper.BeginTransaction()

	for {
		// insert
		sql := fmt.Sprintf(`insert into content_comment (id, subject, content, createDate, creater) values (%d,'%s','%s','%s', %d)`, id, subject, content, createDate, creater)
		_, result = helper.Execute(sql)
		if !result {
			break
		}

		cmt.ID = id
		res := resource.CreateSimpleRes(cmt.ID, model.COMMENT, cmt.Name, cmt.Description, cmt.CreateDate, cmt.Creater)
		for _, c := range cmt.Catalog {
			if c != common_const.SystemContentCatalog.ID {
				ca, ok := resource.QueryResourceByID(helper, c, model.CATALOG)
				if ok {
					res.AppendRelative(ca)
				} else {
					result = false
					break
				}
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

	return cmt, result
}

// SaveComment 保存Comment
func SaveComment(helper dbhelper.DBHelper, cmt model.CommentDetail) (model.Summary, bool) {
	desc := util.ExtractSummary(cmt.Content)
	summary := model.Summary{Unit: model.Unit{ID: cmt.ID, Name: cmt.Subject}, Description: desc, Type: model.COMMENT, Catalog: cmt.Catalog, CreateDate: cmt.CreateDate, Creater: cmt.Creater}
	result := false
	helper.BeginTransaction()

	for {
		// modify
		sql := fmt.Sprintf(`update content_comment set subject ='%s', content ='%s', createdate='%s', creater=%d where id=%d`, cmt.Subject, cmt.Content, cmt.CreateDate, cmt.Creater, cmt.ID)
		_, result = helper.Execute(sql)

		if result {
			res, ok := resource.QueryResourceByID(helper, cmt.ID, model.COMMENT)
			if !ok {
				result = false
				break
			}

			res.UpdateName(cmt.Subject)
			res.UpdateDescription(desc)
			res.ResetRelative()
			for _, c := range cmt.Catalog {
				if c != common_const.SystemContentCatalog.ID {
					ca, ok := resource.QueryResourceByID(helper, c, model.CATALOG)
					if ok {
						res.AppendRelative(ca)
					} else {
						result = false
						break
					}
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
