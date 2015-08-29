package content

import (
	"fmt"
	"log"
	"muidea.com/dao"
	"webcenter/auth"
)

type ArticleInfo struct {
	Id int
	Title string
	CreateDate string
	Catalog Catalog
	Author auth.User
}


func newArticleInfo() ArticleInfo {
	articleInfo := ArticleInfo{}
	articleInfo.Id = -1
	articleInfo.Catalog = newCatalog()
	articleInfo.Author = auth.NewUser()
	
	return articleInfo
}

func GetAllArticleInfo(dao * dao.Dao) []ArticleInfo {
	articleInfoList := []ArticleInfo{}
	sql := fmt.Sprintf("select id, title, author, createdate, catalog from article")
	if !dao.Query(sql) {
		log.Printf("query article failed, sql:%s", sql)
		return articleInfoList
	}

	for dao.Next() {
		articleInfo := newArticleInfo()
		dao.GetField(&articleInfo.Id, &articleInfo.Title, &articleInfo.Author.Id, &articleInfo.CreateDate, &articleInfo.Catalog.Id)
		
		articleInfo.Catalog.Query(dao)
		articleInfo.Author.Query(dao)
		
		log.Printf("%d,%s", articleInfo.Id, articleInfo.Title)
						
		articleInfoList = append(articleInfoList, articleInfo)
	}
	
	return articleInfoList
}

