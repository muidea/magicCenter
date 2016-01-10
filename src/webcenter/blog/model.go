package blog


import (
	"log"
	"html/template"
	"webcenter/modelhelper"
	"webcenter/admin/content/article"
	"webcenter/admin/content/catalog"
	"webcenter/admin/content/link"
)

type ArticleContent struct {
	Id int
	Title string
	Content template.HTML
	CreateDate string
	Catalog []string
	Author string	
}

type CatalogContent struct {
	Id int
	Name string
}

type LinkContent struct {
	Name string
	Url string
	Logo string
	Style int	
} 

func GetArticleSummary(model modelhelper.Model, begin int, offSet int) []ArticleContent {
	articleSummaryList := []ArticleContent{}
	
	articleLis := article.QueryArticleDetailByRang(model, begin, offSet)
	for ii := 0; ii < len(articleLis); ii++ {
		article := articleLis[ii]
		summary := ArticleContent{}
				
		summary.Id = article.Id
		summary.Title = article.Title
		summary.Content = template.HTML(article.Content)
		summary.CreateDate = article.CreateDate
		summary.Catalog = article.Catalog
		summary.Author = article.Author
		
		articleSummaryList = append(articleSummaryList,summary)
	}
	
	log.Printf("QueryArticleDetailByRang, begin:%d, offSet:%d, result Size:%d", begin, offSet, len(articleSummaryList))
		
	return articleSummaryList
}

func GetCatalog(model modelhelper.Model) []CatalogContent {
	articleCatalogList := []CatalogContent{}
	
	catalogLis := catalog.QueryAllCatalogInfo(model)
	for ii := 0; ii < len(catalogLis); ii++ {
		catalog := catalogLis[ii]
		cnt := CatalogContent{}

		cnt.Id = catalog.Id
		cnt.Name = catalog.Name
		
		articleCatalogList = append(articleCatalogList, cnt)
	}
	
	return articleCatalogList
}

func GetLink(model modelhelper.Model) []LinkContent {
	siteLinkList := []LinkContent{}
	
	links := link.QueryAllLink(model)
	for ii := 0; ii < len(links); ii++ {
		link := links[ii]
		cnt := LinkContent{}

		cnt.Name = link.Name
		cnt.Url = link.Url
		cnt.Logo = link.Logo
		cnt.Style = link.Style
				
		siteLinkList = append(siteLinkList, cnt)
	}
	
	return siteLinkList
}

func GetArticleContent(model modelhelper.Model, id int) (ArticleContent, bool) {
	cnt := ArticleContent{}
	
	ar, found := article.QueryArticleDetailById(model, id)
	if found {
		cnt.Id = ar.Id
		cnt.Title = ar.Title
		cnt.Content = template.HTML(ar.Content)
		cnt.CreateDate = ar.CreateDate		
		cnt.Catalog = ar.Catalog
		cnt.Author = ar.Author
	}
	
	return cnt, found
}
