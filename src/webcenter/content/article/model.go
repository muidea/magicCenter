package article

import (
	"fmt"
	"log"
	"html"
	"webcenter/modelhelper"
)

type ArticleInfo struct {
	Id int
	Title string
	CreateDate string
	Catalog int
	Author int
}

type Article struct {
	Id int
	Title string
	Content string
	CreateDate string
	Catalog int
	Author int
}


func newArticleInfo() ArticleInfo {
	articleInfo := ArticleInfo{}
	articleInfo.Id = -1
	articleInfo.Catalog = -1
	articleInfo.Author = -1
	
	return articleInfo
}

func newArticle() Article {
	article := Article{}
	article.Id = -1
	article.Catalog = -1
	article.Author = -1
	
	return article
}

func GetAllArticleInfo(model modelhelper.Model) []ArticleInfo {
	articleInfoList := []ArticleInfo{}
	sql := fmt.Sprintf("select id, title, author, createdate, catalog from article")
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return articleInfoList
	}

	for model.Next() {
		articleInfo := newArticleInfo()
		model.GetValue(&articleInfo.Id, &articleInfo.Title, &articleInfo.Author, &articleInfo.CreateDate, &articleInfo.Catalog)
		
		articleInfoList = append(articleInfoList, articleInfo)
	}
	
	return articleInfoList
}

func GetArticleByCatalog(model modelhelper.Model, id int) []ArticleInfo {
	articleInfoList := []ArticleInfo{}
	sql := fmt.Sprintf("select id, title, author, createdate, catalog from article where catalog=%d", id)
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return articleInfoList
	}

	for model.Next() {
		articleInfo := newArticleInfo()
		model.GetValue(&articleInfo.Id, &articleInfo.Title, &articleInfo.Author, &articleInfo.CreateDate, &articleInfo.Catalog)
		
		articleInfoList = append(articleInfoList, articleInfo)
	}
		
	return articleInfoList	
}

func QueryArticleByRang(model modelhelper.Model, begin int,offset int) []Article {
	articleList := []Article{}
	sql := fmt.Sprintf("select * from (select id, title, content, author, createdate, catalog from article order by id) c where id > %d limit %d", begin, offset)
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return articleList
	}

	for model.Next() {
		article := newArticle()
		model.GetValue(&article.Id, &article.Title, &article.Content, &article.Author, &article.CreateDate, &article.Catalog)
		
		article.Content = html.UnescapeString(article.Content)
		articleList = append(articleList, article)
	}
		
	return articleList	
}


func QueryArticleById(model modelhelper.Model, id int) (Article, bool) {
	article := Article{}
	sql := fmt.Sprintf("select id, title, content, author, createdate, catalog from article where id = %d", id)
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return article, false
	}

	result := false
	for model.Next() {
		result = model.GetValue(&article.Id, &article.Title, &article.Content, &article.Author, &article.CreateDate, &article.Catalog)
		if result {
			article.Content = html.UnescapeString(article.Content)
		}
	}
	
	return article, result	
}

func DeleteArticleById(model modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from article where id=%d", id)
	
	result := model.Execute(sql)
	
	return result	
}

func SaveArticle(model modelhelper.Model, article Article) bool {
	sql := fmt.Sprintf("select id from article where id=%d", article.Id)
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into article (title,content,author,createdate,catalog) values ('%s','%s',%d,'%s',%d)", article.Title, html.EscapeString(article.Content), article.Author, article.CreateDate, article.Catalog)
	} else {
		// modify
		sql = fmt.Sprintf("update article set title ='%s', content ='%s', author =%d, createdate ='%s', catalog =%d where id=%d", article.Title, html.EscapeString(article.Content), article.Author, article.CreateDate, article.Catalog, article.Id)
	}
	
	result = model.Execute(sql)
	
	return result
}


func (this *Article)Query(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id, title, content, author, createdate, catalog from article where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		result = model.GetValue(&this.Id, &this.Title, &this.Content, &this.Author, &this.CreateDate, &this.Catalog)
		
		if result {		
			this.Content = html.UnescapeString(this.Content)
		}		
	}
	
	return result		
}

func (this *Article)delete(model modelhelper.Model) bool {
	sql := fmt.Sprintf("delete from article where id=%d", this.Id)
	
	result := model.Execute(sql)
	
	return result
}

func (this *Article)save(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id from article where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into article (title,content,author,createdate,catalog) values ('%s','%s',%d,'%s',%d)", this.Title, html.EscapeString(this.Content), this.Author, this.CreateDate, this.Catalog)
	} else {
		// modify
		sql = fmt.Sprintf("update article set title ='%s', content ='%s', author =%d, createdate ='%s', catalog =%d where id=%d", this.Title, html.EscapeString(this.Content), this.Author, this.CreateDate, this.Catalog, this.Id)
	}
	
	result = model.Execute(sql)
	
	return result
}


