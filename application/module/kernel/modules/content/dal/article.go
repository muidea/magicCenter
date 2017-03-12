package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllArticleSummary 查询所有文章摘要
func QueryAllArticleSummary(helper dbhelper.DBHelper) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	sql := fmt.Sprintf(`select id, title, author, createdate from article`)
	helper.Query(sql)

	for helper.Next() {
		summary := model.ArticleSummary{}
		helper.GetValue(&summary.ID, &summary.Title, &summary.Author, &summary.CreateDate)

		articleSummaryList = append(articleSummaryList, summary)
	}

	for index, value := range articleSummaryList {
		summary := &articleSummaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return articleSummaryList
}

// QueryArticleByID 查询指定文章
func QueryArticleByID(helper dbhelper.DBHelper, id int) (model.Article, bool) {
	ar := model.Article{}

	sql := fmt.Sprintf(`select id, title, content, author, createdate from article where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&ar.ID, &ar.Title, &ar.Content, &ar.Author, &ar.CreateDate)
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
func QueryArticleSummaryByCatalog(helper dbhelper.DBHelper, id int) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.ARTICLE)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, title, author, createdate from article where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			summary := model.ArticleSummary{}
			helper.GetValue(&summary.ID, &summary.Title, &summary.Author, &summary.CreateDate)

			articleSummaryList = append(articleSummaryList, summary)
		}
	}

	for index, value := range articleSummaryList {
		summary := &articleSummaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return articleSummaryList
}

// CreateArticle 保存文章
func CreateArticle(helper dbhelper.DBHelper, title, content string, catalogs []int, author int, createDate string) (model.ArticleSummary, bool) {
	article := model.ArticleSummary{}
	article.Title = title
	article.Author = author
	article.Catalog = catalogs
	article.CreateDate = createDate

	// insert
	sql := fmt.Sprintf(`insert into article (title,content,author,createdate) values ('%s','%s',%d,'%s')`, title, content, author, createDate)
	num, result := helper.Execute(sql)
	if num != 1 || !result {
		return article, false
	}

	sql = fmt.Sprintf(`select id from article where title='%s' and author =%d and createdate='%s'`, title, author, createDate)

	helper.Query(sql)
	result = false
	if helper.Next() {
		helper.GetValue(&article.ID)
		result = true
	}

	if result {
		res := resource.CreateSimpleRes(article.ID, model.ARTICLE, article.Title)
		for _, c := range article.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return article, result
}

// SaveArticle 保存文章
func SaveArticle(helper dbhelper.DBHelper, article model.Article) (model.ArticleSummary, bool) {
	sql := fmt.Sprintf(`select id from article where id=%d`, article.ID)
	helper.Query(sql)

	summary := model.ArticleSummary{ID: article.ID, Title: article.Title, CreateDate: article.CreateDate, Catalog: article.Catalog, Author: article.Author}
	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		return summary, false
	}

	// modify
	sql = fmt.Sprintf(`update article set title ='%s', content ='%s', author =%d, createdate ='%s' where id=%d`, article.Title, article.Content, article.Author, article.CreateDate, article.ID)
	_, result = helper.Execute(sql)

	if result {
		res := resource.CreateSimpleRes(article.ID, model.ARTICLE, article.Title)
		for _, c := range article.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return summary, result
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
