package bll

import (
	"log"
	"time"
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/content/dal"
    "magiccenter/kernel/content/model"
)

func QueryAllArticleSummary() []model.ArticleSummary {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	
	
	return dal.QueryAllArticleSummary(helper)
}

func QueryArticleById(id int) (model.Article, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	
	
	ar, result := dal.QueryArticleById(helper, id)
		
	return ar, result		
}

func DeleteArticle(id int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.DeleteArticle(helper, id)	
}

func SaveArticle(id int, title, content string, uId int, catalogs []int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	article := model.Article{}
	article.Id = id
	article.Title = title
	article.Content = content
	article.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	article.Author.Id = uId
	
	for _, ca := range catalogs {
		catalog, found := dal.QueryCatalogById(helper, ca)
		if found {
			c := model.Catalog{}
			c.Id = catalog.Id
			c.Name = catalog.Name
			article.Catalog = append(article.Catalog, c)
		} else {
			log.Printf("illegal catalog id, id:%d", ca)
		}
	}
		
	return dal.SaveArticle(helper, article)
}

func QueryArticleByCatalog(id int) []model.ArticleSummary {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryArticleByCatalog(helper, id)
}

func QueryArticleByRang(begin int,offset int) []model.ArticleSummary {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()	

	return dal.QueryArticleByRang(helper, begin, offset)
}


