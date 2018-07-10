package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/resource"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

func loadArticleID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from content_article`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// QueryAllArticleSummary 查询所有文章摘要
func QueryAllArticleSummary(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}

	ress := resource.QueryResourceByType(helper, model.ARTICLE)
	for _, v := range ress {
		summary := model.Summary{Unit: model.Unit{ID: v.RId(), Name: v.RName()}, Description: v.RDescription(), Type: v.RType(), CreateDate: v.RCreateDate(), Creater: v.ROwner()}

		for _, r := range v.Relative() {
			summary.Catalog = append(summary.Catalog, r.RId())
		}

		summaryList = append(summaryList, summary)
	}

	return summaryList
}

// QueryArticles 查询指定文章
func QueryArticles(helper dbhelper.DBHelper, ids []int) []model.Article {
	articleList := []model.Article{}

	if len(ids) == 0 {
		return articleList
	}

	sql := fmt.Sprintf(`select id, title from content_article where id in(%s)`, util.IntArray2Str(ids))
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		summary := model.Article{}
		helper.GetValue(&summary.ID, &summary.Name)

		articleList = append(articleList, summary)
	}

	return articleList
}

// QueryArticleByID 查询指定文章
func QueryArticleByID(helper dbhelper.DBHelper, id int) (model.ArticleDetail, bool) {
	ar := model.ArticleDetail{}

	sql := fmt.Sprintf(`select id, title, content, creater, createdate from content_article where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&ar.ID, &ar.Name, &ar.Content, &ar.Creater, &ar.CreateDate)
		result = true
	}
	helper.Finish()

	if result {
		ress := resource.QueryRelativeResource(helper, ar.ID, model.ARTICLE)
		for _, r := range ress {
			ar.Catalog = append(ar.Catalog, r.RId())
		}
	}

	return ar, result
}

// QueryArticleSummaryByCatalog 查询指定分类下的所有文章摘要
func QueryArticleSummaryByCatalog(helper dbhelper.DBHelper, catalog int) []model.Summary {
	summaryList := []model.Summary{}
	resList := resource.QueryReferenceResource(helper, catalog, model.CATALOG, model.ARTICLE)
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

// CreateArticle 保存文章
func CreateArticle(helper dbhelper.DBHelper, title, content string, catalogs []int, creater int, createDate string) (model.Summary, bool) {
	article := model.Summary{Unit: model.Unit{Name: title}, Type: model.ARTICLE, Catalog: catalogs, CreateDate: createDate, Creater: creater}

	id := allocArticleID()
	result := false
	helper.BeginTransaction()
	for {
		// insert
		sql := fmt.Sprintf(`insert into content_article (id, title,content,creater,createdate) values (%d, '%s','%s',%d,'%s')`, id, title, content, creater, createDate)
		_, result = helper.Execute(sql)
		if !result {
			break
		}

		desc := util.ExtractSummary(content)
		article.ID = id
		res := resource.CreateSimpleRes(article.ID, model.ARTICLE, article.Name, desc, article.CreateDate, article.Creater)
		for _, c := range article.Catalog {
			ca, ok := resource.QueryResourceByID(helper, c, model.CATALOG)
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

	return article, result
}

// SaveArticle 保存文章
func SaveArticle(helper dbhelper.DBHelper, article model.ArticleDetail) (model.Summary, bool) {
	summary := model.Summary{Unit: model.Unit{ID: article.ID, Name: article.Name}, Type: model.ARTICLE, Catalog: article.Catalog, CreateDate: article.CreateDate, Creater: article.Creater}
	result := false

	helper.BeginTransaction()
	for {
		// modify
		sql := fmt.Sprintf(`update content_article set title ='%s', content ='%s', creater =%d, createdate ='%s' where id=%d`, article.Name, article.Content, article.Creater, article.CreateDate, article.ID)
		_, result = helper.Execute(sql)

		if result {
			res, ok := resource.QueryResourceByID(helper, article.ID, model.ARTICLE)
			if !ok {
				result = false
				break
			}

			res.UpdateName(article.Name)
			desc := util.ExtractSummary(article.Content)
			res.UpdateDescription(desc)

			res.ResetRelative()
			for _, c := range article.Catalog {
				ca, ok := resource.QueryResourceByID(helper, c, model.CATALOG)
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

			break
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

// DeleteArticle 删除文章
func DeleteArticle(helper dbhelper.DBHelper, id int) bool {
	result := false
	helper.BeginTransaction()

	for {
		sql := fmt.Sprintf(`delete from content_article where id=%d`, id)

		_, result = helper.Execute(sql)
		if result {
			res, ok := resource.QueryResourceByID(helper, id, model.ARTICLE)
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
