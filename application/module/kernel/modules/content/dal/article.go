package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllArticleSummary 查询所有文章摘要
func QueryAllArticleSummary(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}
	sql := fmt.Sprintf(`select id, title,createdate,creater from article`)
	helper.Query(sql)

	for helper.Next() {
		summary := model.Summary{}
		helper.GetValue(&summary.ID, &summary.Name, &summary.CreateDate, &summary.Creater)

		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryArticleByID 查询指定文章
func QueryArticleByID(helper dbhelper.DBHelper, id int) (model.ArticleDetail, bool) {
	ar := model.ArticleDetail{}

	sql := fmt.Sprintf(`select id, title, content, creater, createdate from article where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&ar.ID, &ar.Name, &ar.Content, &ar.Creater, &ar.CreateDate)
		result = true
	}

	if result {
		ress := resource.QueryRelativeResource(helper, ar.ID, model.ARTICLE)
		for _, r := range ress {
			ar.Catalog = append(ar.Catalog, r.RId())
		}
	}

	return ar, result
}

// QueryArticleSummaryByCatalog 查询指定分类下的所有文章摘要
func QueryArticleSummaryByCatalog(helper dbhelper.DBHelper, id int) []model.Summary {
	summaryList := []model.Summary{}
	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.ARTICLE)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, title, createdate,creater from article where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			summary := model.Summary{}
			helper.GetValue(&summary.ID, &summary.Name, &summary.CreateDate, &summary.Creater)

			summaryList = append(summaryList, summary)
		}
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// CreateArticle 保存文章
func CreateArticle(helper dbhelper.DBHelper, title, content string, catalogs []int, creater int, createDate string) (model.Summary, bool) {
	article := model.Summary{}
	article.Name = title
	article.Catalog = catalogs

	// insert
	sql := fmt.Sprintf(`insert into article (title,content,creater,createdate) values ('%s','%s',%d,'%s')`, title, content, creater, createDate)
	num, result := helper.Execute(sql)
	if num != 1 || !result {
		return article, false
	}

	sql = fmt.Sprintf(`select id from article where title='%s' and creater =%d and createdate='%s'`, title, creater, createDate)

	helper.Query(sql)
	result = false
	if helper.Next() {
		helper.GetValue(&article.ID)
		result = true
	}

	if result {
		res := resource.CreateSimpleRes(article.ID, model.ARTICLE, article.Name)
		for _, c := range article.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return article, result
}

// SaveArticle 保存文章
func SaveArticle(helper dbhelper.DBHelper, article model.ArticleDetail) (model.Summary, bool) {
	// modify
	sql := fmt.Sprintf(`update article set title ='%s', content ='%s', creater =%d, createdate ='%s' where id=%d`, article.Name, article.Content, article.Creater, article.CreateDate, article.ID)
	num, result := helper.Execute(sql)

	if num == 1 && result {
		res := resource.CreateSimpleRes(article.ID, model.ARTICLE, article.Name)
		for _, c := range article.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	} else {
		result = false
	}

	return model.Summary{ID: article.ID, Name: article.Name, Catalog: article.Catalog, CreateDate: article.CreateDate, Creater: article.Creater}, result
}

// DeleteArticle 删除文章
func DeleteArticle(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from article where id=%d`, id)

	num, result := helper.Execute(sql)
	if num >= 1 && result {
		// 删除资源时，名称时不用关注的，所以这里填“”好了
		res := resource.CreateSimpleRes(id, model.ARTICLE, "")
		result = resource.DeleteResource(helper, res)
	}

	return result
}
