package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/content/model"
	"magiccenter/kernel/account/dal"
)

func QueryAllArticleSummary(helper modelhelper.Model) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	sql := fmt.Sprintf(`select id, title, author, createdate from article`)
	helper.Query(sql)

	for helper.Next() {
		summary := model.ArticleSummary{}
		helper.GetValue(&summary.Id, &summary.Title, &summary.Author.Id, &summary.CreateDate)
		
		articleSummaryList = append(articleSummaryList, summary)
	}
	
	for i, _ := range articleSummaryList {
		summary := &articleSummaryList[i]
		
		user, found := dal.QueryUserById(helper, summary.Author.Id)
		if found {
			summary.Author.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, summary.Id, model.ARTICLE)
		for _, r := range ress {
			ca := model.Catalog{}
			ca.Id = r.RId()
			ca.Name = r.RName()
			
			summary.Catalog = append(summary.Catalog, ca)
		}
	}

	return articleSummaryList
}

func QueryArticleByCatalog(helper modelhelper.Model, id int) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	resList := QueryReferenceResource(helper, id, model.CATALOG, model.ARTICLE)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, title, author, createdate from article where id =%d`, r.RId())
		helper.Query(sql)
		
		if helper.Next() {
			summary := model.ArticleSummary{}
			helper.GetValue(&summary.Id, &summary.Title, &summary.Author.Id, &summary.CreateDate)
			
			articleSummaryList = append(articleSummaryList, summary)
		}
	}
	
	for i, _ := range articleSummaryList {
		summary := &articleSummaryList[i]
		
		user, found := dal.QueryUserById(helper, summary.Author.Id)
		if found {
			summary.Author.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, summary.Id, model.ARTICLE)
		for _, r := range ress {
			ca := model.Catalog{}
			ca.Id = r.RId()
			ca.Name = r.RName()
			
			summary.Catalog = append(summary.Catalog, ca)
		}
	}	

	return articleSummaryList
}

func QueryArticleByRang(helper modelhelper.Model, begin int,offset int) []model.ArticleSummary {
	articleSummaryList := []model.ArticleSummary{}
	sql := fmt.Sprintf(`select id, title, author, createdate from article order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		summary := model.ArticleSummary{}
		helper.GetValue(&summary.Id, &summary.Title, &summary.Author.Id, &summary.CreateDate)

		articleSummaryList = append(articleSummaryList, summary)
	}
	
	for i, _ := range articleSummaryList {
		summary := &articleSummaryList[i]
		
		user, found := dal.QueryUserById(helper, summary.Author.Id)
		if found {
			summary.Author.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, summary.Id, model.ARTICLE)
		for _, r := range ress {
			ca := model.Catalog{}
			ca.Id = r.RId()
			ca.Name = r.RName()
			
			summary.Catalog = append(summary.Catalog, ca)
		}		
	}	

	return articleSummaryList
}

func QueryArticleById(helper modelhelper.Model, id int) (model.Article, bool) {
	ar := model.Article{}
	
	sql := fmt.Sprintf(`select id, title, content, author, createdate from article where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&ar.Id, &ar.Title, &ar.Content, &ar.Author.Id, &ar.CreateDate)
		result = true
	}
	
	if result {
		user, found := dal.QueryUserById(helper, ar.Author.Id)
		if found {
			ar.Author.Name = user.Name
		}
				
		ress := QueryRelativeResource(helper, ar.Id, model.ARTICLE)
		for _, r := range ress {
			ca := model.Catalog{}
			ca.Id = r.RId()
			ca.Name = r.RName()
			
			ar.Catalog = append(ar.Catalog, ca)
		}
	}

	return ar, result	
}

func DeleteArticle(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf(`delete from article where id=%d`, id)
	
	_, result := helper.Execute(sql)
	if result {
		ar := model.Article{}
		ar.Id = id
		result  = DeleteResource(helper, &ar)
	}
	
	return result	
}

func SaveArticle(helper modelhelper.Model, article model.Article) bool {
	sql := fmt.Sprintf(`select id from article where id=%d`, article.Id)
	helper.Query(sql)

	result := false;
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into article (title,content,author,createdate) values ('%s','%s',%d,'%s')`, article.Title, article.Content, article.Author.Id, article.CreateDate)
		_, result = helper.Execute(sql)
		sql = fmt.Sprintf(`select id from article where title='%s' and author =%d and createdate='%s'`, article.Title, article.Author.Id, article.CreateDate)
		
		helper.Query(sql)
		result = false
		if helper.Next() {
			helper.GetValue(&article.Id)
			result = true
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update article set title ='%s', content ='%s', author =%d, createdate ='%s' where id=%d`, article.Title, article.Content, article.Author.Id, article.CreateDate, article.Id)
		_, result = helper.Execute(sql)
	}
	
	if result {
		result = SaveResource(helper, &article)
	}
	
	return result
}


