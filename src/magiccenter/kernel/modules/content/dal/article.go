package dal

import (
	"fmt"
	"magiccenter/kernel/modules/content/model"
	resdal "magiccenter/resource/dal"
	"magiccenter/util/dbhelper"
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

	for _, summary := range articleSummaryList {
		ress := resdal.QueryRelativeResource(helper, summary.ID, model.ARTICLE)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return articleSummaryList
}

// QueryArticleByCatalog 查询指定分类下的所有文章摘要
func QueryArticleByCatalog(helper dbhelper.DBHelper, id int) []model.ArticleSummary {
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
func QueryArticleByRang(helper dbhelper.DBHelper, begin int, offset int) []model.ArticleSummary {
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
		ress := resdal.QueryRelativeResource(helper, ar.ID, model.ARTICLE)
		for _, r := range ress {
			ar.Catalog = append(ar.Catalog, r.RId())
		}
	}

	return ar, result
}

// DeleteArticle 删除文章
func DeleteArticle(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from article where id=%d`, id)

	num, result := helper.Execute(sql)
	if num >= 1 && result {
		// 删除资源时，名称时不用关注的，所以这里填“”好了
		res := resdal.CreateSimpleRes(id, model.ARTICLE, "")
		result = resdal.DeleteResource(helper, res)
	}

	return result
}

// SaveArticle 保存文章
func SaveArticle(helper dbhelper.DBHelper, article model.Article) bool {
	sql := fmt.Sprintf(`select id from article where id=%d`, article.ID)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into article (title,content,author,createdate) values ('%s','%s',%d,'%s')`, article.Title, article.Content, article.Author, article.CreateDate)
		_, result = helper.Execute(sql)
		sql = fmt.Sprintf(`select id from article where title='%s' and author =%d and createdate='%s'`, article.Title, article.Author, article.CreateDate)

		helper.Query(sql)
		result = false
		if helper.Next() {
			helper.GetValue(&article.ID)
			result = true
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update article set title ='%s', content ='%s', author =%d, createdate ='%s' where id=%d`, article.Title, article.Content, article.Author, article.CreateDate, article.ID)
		_, result = helper.Execute(sql)
	}

	if result {
		res := resdal.CreateSimpleRes(article.ID, model.ARTICLE, article.Title)
		for _, c := range article.Catalog {
			ca := resdal.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resdal.SaveResource(helper, res)
	}

	return result
}
