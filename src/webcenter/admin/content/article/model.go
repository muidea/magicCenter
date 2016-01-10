package article

import (
	"fmt"
	"html"
	"log"
	"webcenter/modelhelper"
	"webcenter/admin/common"
	"webcenter/admin/content/base"
)

type ArticleSummary struct {
	Id int
	Title string
	CreateDate string
	Catalog []string
	Author string
}

type ArticleDetail struct {
	Id int
	Title string
	Content string
	CreateDate string
	Catalog []string
	Author string
}

type Article interface {
	common.Resource
	Content() string
	CreateDate() string
	Author() int
	SetId(id int)
	SetName(name string)
	SetContent(content string)
	SetCreateDate(date string)
	SetAuthor(author int)
	SetCatalog(catalog []int)
}

type article struct {
	id int
	title string
	content string
	createDate string
	catalog []int
	author int
}


func (this *article) Id() int {
	return this.id
}

func (this *article) Name() string {
	return this.title
}

func (this *article) Type() int {
	return base.ARTICLE
}

func (this *article)Relative() []common.Resource {
	ress := []common.Resource{}
	
	for _, pid := range this.catalog {
		res := common.NewSimpleRes(pid,"", base.CATALOG)
		ress = append(ress, res)
	}
	
	return ress
}

func (this *article)Content() string {
	return this.content
}

func (this *article)CreateDate() string {
	return this.createDate
}

func (this *article)Author() int {
	return this.author
}

func (this *article)SetId(id int) {
	this.id = id
}

func (this *article)SetName(name string) {
	this.title = name
}

func (this *article)SetContent(content string) {
	this.content = content
}

func (this *article)SetCreateDate(date string) {
	this.createDate =date
}

func (this *article)SetAuthor(author int) {
	this.author = author
}

func (this *article)SetCatalog(catalog []int) {
	this.catalog = catalog
}

func NewArticle() Article {
	a := &article{}
	a.id = -1
	
	return a
}

func QueryAllArticle(model modelhelper.Model) []ArticleSummary {
	articleSummaryList := []ArticleSummary{}
	sql := fmt.Sprintf(`select a.id, a.title, u.nickname, a.createdate from article a, user u where a.author = u.id`)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		articleSummary := ArticleSummary{}
		model.GetValue(&articleSummary.Id, &articleSummary.Title, &articleSummary.Author, &articleSummary.CreateDate)
		
		articleSummaryList = append(articleSummaryList, articleSummary)
	}
	
	for index, summary := range articleSummaryList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, summary.Id, base.ARTICLE)
		name := "-"
		log.Print(sql)
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					articleSummaryList[index].Catalog = append(articleSummaryList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}
	}

	return articleSummaryList
}

func QueryArticleByCatalog(model modelhelper.Model, id int) []ArticleSummary {
	articleSummaryList := []ArticleSummary{}
	sql := fmt.Sprintf(`select a.id, a.title, u.nickname, a.createdate from article a, user u where a.author = u.id and a.id in (
		select src from resource_relative where dst = %d and dstType = %d and srcType = %d )`, id, base.CATALOG, base.ARTICLE)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		articleSummary := ArticleSummary{}
		model.GetValue(&articleSummary.Id, &articleSummary.Title, &articleSummary.Author, &articleSummary.CreateDate)
		
		articleSummaryList = append(articleSummaryList, articleSummary)
	}
	
	for index, summary := range articleSummaryList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, summary.Id, base.ARTICLE)
		name := "-"
		log.Print(sql)
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					articleSummaryList[index].Catalog = append(articleSummaryList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}				
	}

	return articleSummaryList
}

func QueryArticleByRang(model modelhelper.Model, begin int,offset int) []ArticleSummary {
	articleSummaryList := []ArticleSummary{}
	sql := fmt.Sprintf(`select a.id, a.title, u.nickname, a.createdate from article a, user u where a.author = u.id and a.id in (
		select src from resource_relative where dstType = %d and srcType = %d ) and a.id >= %d limit %d`, base.CATALOG, base.ARTICLE, begin, offset)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		articleSummary := ArticleSummary{}
		model.GetValue(&articleSummary.Id, &articleSummary.Title, &articleSummary.Author, &articleSummary.CreateDate)
		
		articleSummaryList = append(articleSummaryList, articleSummary)
	}
	
	for index, summary := range articleSummaryList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, summary.Id, base.ARTICLE)
		name := "-"
		log.Print(sql)
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					articleSummaryList[index].Catalog = append(articleSummaryList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}				
	}

	return articleSummaryList
}

func QueryArticleDetailByRang(model modelhelper.Model, begin int,offset int) []ArticleDetail {
	articleDetailList := []ArticleDetail{}
	sql := fmt.Sprintf(`select a.id, a.title, a.content, u.nickname, a.createdate from article a, user u where a.author = u.id and a.id in (
		select src from resource_relative where dstType = %d and srcType = %d ) and a.id >= %d limit %d`, base.CATALOG, base.ARTICLE, begin, offset)
	
	log.Print(sql)
		
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		articleDetail := ArticleDetail{}
		result := model.GetValue(&articleDetail.Id, &articleDetail.Title, &articleDetail.Content, &articleDetail.Author, &articleDetail.CreateDate)
		if result {
			articleDetail.Content = html.UnescapeString(articleDetail.Content)
		}
				
		articleDetailList = append(articleDetailList, articleDetail)
	}
	
	for index, detail := range articleDetailList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, detail.Id, base.ARTICLE)
		name := "-"
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					articleDetailList[index].Catalog = append(articleDetailList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}
				
	}

	return articleDetailList
}

func QueryArticleDetailById(model modelhelper.Model, id int) (ArticleDetail, bool) {
	article := ArticleDetail{}
	
	sql := fmt.Sprintf(`select a.id, a.title, a.content, u.nickname, a.createdate from article a, user u where a.author = u.id and a.id = %d`, id)
	if !model.Query(sql) {          
		panic("query failed")
	}

	result := false
	for model.Next() {
		result = model.GetValue(&article.Id, &article.Title, &article.Content, &article.Author, &article.CreateDate)
		if result {
			article.Content = html.UnescapeString(article.Content)
		}
		break
	}
	if !result {
		return article, result
	}

	sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, article.Id, base.ARTICLE)
	name := "-"
	if model.Query(sql) {
		for model.Next() {
			if model.GetValue(&name) {
				article.Catalog = append(article.Catalog, name)
			}
		}
	} else {
		panic("query failed")
	}
	
	return article, result	
}

func QueryArticleById(model modelhelper.Model, id int) (Article, bool) {
	article := &article{}
	
	sql := fmt.Sprintf(`select id, title, content,author, createdate from article where id = %d`, id)
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false
	for model.Next() {
		result = model.GetValue(&article.id, &article.title, &article.content, &article.author, &article.createDate)
		if result {
			article.content = html.UnescapeString(article.content)
		}
		break
	}
	if !result {
		return article, result
	}

	sql = fmt.Sprintf(`select dst from resource_relative where src = %d and srcType = %d and dstType =%d`, article.id, base.ARTICLE, base.CATALOG)
	pid := -1
	if model.Query(sql) {
		for model.Next() {
			if model.GetValue(&pid) {
				article.catalog = append(article.catalog, pid)
			}
		}
	} else {
		panic("query failed")
	}
	
	return article, result	
}

func DeleteArticleById(model modelhelper.Model, id int) bool {
	if !model.BeginTransaction() {
		return false
	}
	
	sql := fmt.Sprintf(`delete from article where id=%d`, id)
	
	result := model.Execute(sql)
	if result {
		ar := article{}
		ar.id = id
		result  = common.DeleteResource(model, &ar)
	}
		
	if !result {
		model.Rollback()
	} else {
		model.Commit()
	}
		
	return result	
}

func SaveArticle(model modelhelper.Model, article Article) bool {
	sql := fmt.Sprintf(`select id from article where id=%d`, article.Id())
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
	}

	if !model.BeginTransaction() {
		return false
	}
	
	if !result {
		// insert
		sql = fmt.Sprintf(`insert into article (title,content,author,createdate) values ('%s','%s',%d,'%s')`, article.Name(), html.EscapeString(article.Content()), article.Author(), article.CreateDate())
		result = model.Execute(sql)
		sql = fmt.Sprintf(`select id from article where title='%s' and author =%d and createdate='%s'`, article.Name(), article.Author(), article.CreateDate())
		
		id := -1
		result = model.Query(sql)
		if result {
			result = false
			for model.Next() {
				result = model.GetValue(&id)
			}
			
			article.SetId(id)
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update article set title ='%s', content ='%s', author =%d, createdate ='%s' where id=%d`, article.Name(), html.EscapeString(article.Content()), article.Author(), article.CreateDate(), article.Id())
		result = model.Execute(sql)
	}
	
	if result {
		result = common.SaveResource(model, article)
	}
	
	if result {
		model.BeginTransaction()
	} else {
		model.Rollback()
	}
	
	return result
}


