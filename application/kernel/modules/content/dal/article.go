package dal

import (
	"fmt"
	"magiccenter/common"
	"magiccenter/common/model"
	resdal "magiccenter/resource/dal"
)

// QueryAllArticleSummary 查询所有文章摘要
func QueryAllArticleSummary(helper common.DBHelper) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	sql := fmt.Sprintf(`select id, title, author, createdate from article`)
	helper.Query(sql)

	for helper.Next() {
		summary := model.ArticleSummary{}
		helper.GetValue(&summary.ID, &summary.Title, &summary.Author, &summary.CreateDate)

		articleSummaryList = append(articleSummaryList, summary)
	}

	for index, _ := range articleSummaryList {
		summary := &articleSummaryList[index]
		ress := resdal.QueryRelativeResource(helper, summary.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return articleSummaryList
}

// QueryArticleByCatalog 查询指定分类下的所有文章摘要
func QueryArticleByCatalog(helper common.DBHelper, id int) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	resList := resdal.QueryReferenceResource(helper, id, model.CATALOG, model.ARTICLE)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, title, author, createdate from article where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			summary := model.ArticleSummary{}
			helper.GetValue(&summary.ID, &summary.Title, &summary.Author, &summary.CreateDate)

			articleSummaryList = append(articleSummaryList, summary)
		}
	}

	for _, summary := range articleSummaryList {
		ress := resdal.QueryRelativeResource(helper, summary.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return articleSummaryList
}

// QueryArticleByRang 查询指定范围的文章摘要
func QueryArticleByRang(helper common.DBHelper, begin int, offset int) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	sql := fmt.Sprintf(`select id, title, author, createdate from article order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		summary := model.ArticleSummary{}
		helper.GetValue(&summary.ID, &summary.Title, &summary.Author, &summary.CreateDate)

		articleSummaryList = append(articleSummaryList, summary)
	}

	for _, summary := range articleSummaryList {
		ress := resdal.QueryRelativeResource(helper, summary.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return articleSummaryList
}

// QueryArticleByID 查询指定文章
func QueryArticleByID(helper common.DBHelper, id int) (model.Article, bool) {
	ar := model.Article{}

	sql := fmt.Sprintf(`select id, title, content, author, createdate from article where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&ar.ID, &ar.Title, &ar.Content, &ar.Author, &ar.CreateDate)
		result = true
	}

	if result {
		ress := resdal.QueryRelativeResource(helper, ar.ID, model.ARTICLE)
		for _, r := range ress {
			ar.Catalog = append(ar.Catalog, r.RId())
		}
	}

	return ar, result
}

// DeleteArticle 删除文章
func DeleteArticle(helper common.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from article where id=%d`, id)

	num, result := helper.Execute(sql)
	if num >= 1 && result {
		// 删除资源时，名称时不用关注的，所以这里填“”好了
		res := resdal.CreateSimpleRes(id, model.ARTICLE, "")
		result = resdal.DeleteResource(helper, res)
	}

	return result
}

// CreateArticle 保存文章
func CreateArticle(helper common.DBHelper, title, content string, catalogs []int, author int, createDate string) (model.ArticleSummary, bool) {
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
		res := resdal.CreateSimpleRes(article.ID, model.ARTICLE, article.Title)
		for _, c := range article.Catalog {
			ca := resdal.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resdal.SaveResource(helper, res)
	}

	return article, result
}

// SaveArticle 保存文章
func SaveArticle(helper common.DBHelper, article model.Article) (model.ArticleSummary, bool) {
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
		res := resdal.CreateSimpleRes(article.ID, model.ARTICLE, article.Title)
		for _, c := range article.Catalog {
			ca := resdal.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resdal.SaveResource(helper, res)
	}

	return summary, result
}
