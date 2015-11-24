package ui


import (
	"html/template"
	"webcenter/modelhelper"
	"webcenter/content/article"
	"webcenter/content/catalog"
	"webcenter/content/image"
	"webcenter/content/link"
)

type ArticleSummary struct {
	Id int
	Title string
	Content template.HTML
	CreateDate string
	Catalog string
	Author string	
}

type ArticleContent struct {
	Id int
	Title string
	Content template.HTML
	CreateDate string
	Catalog string
	Author string	
}

type ArticleCatalog struct {
	Id int
	Name string
}

type SiteLink struct {
	Name string
	Url string
	Logo string
	Style int	
} 

func GetArticleSummary(model modelhelper.Model, begin int, end int) []ArticleSummary {
	articleSummaryList := []ArticleSummary{}
	
	articleLis := article.QueryArticleByRang(model,0, 4)
	for ii := 0; ii < len(articleLis); ii++ {
		article := articleLis[ii]
		summary := ArticleSummary{}
				
		summary.Id = article.Id
		summary.Title = article.Title
		summary.Content = template.HTML(article.Content)
		summary.CreateDate = article.CreateDate
		catalog, found := catalog.QueryCatalogById(model, article.Catalog)
		if found {
			summary.Catalog = catalog.Name
		}
		
		
		summary.Author = article.Author.NickName
		
		articleSummaryList = append(articleSummaryList,summary)
	}
	
	return articleSummaryList
}

func (this *Model)GetArticleCatalog() []ArticleCatalog {
	articleCatalogList := []ArticleCatalog{}
	
	catalogLis := content.GetAllCatalog(this.dao)
	for ii := 0; ii < len(catalogLis); ii++ {
		catalog := catalogLis[ii]
		cnt := ArticleCatalog{}

		cnt.Id = catalog.Id
		cnt.Name = catalog.Name
				
		articleCatalogList = append(articleCatalogList, cnt)
	}
	
	return articleCatalogList
}

func (this *Model)GetSiteLink() []SiteLink {
	siteLinkList := []SiteLink{}
	
	links := content.GetAllLink(this.dao)
	for ii := 0; ii < len(links); ii++ {
		link := links[ii]
		cnt := SiteLink{}

		cnt.Name = link.Name
		cnt.Url = link.Url
		cnt.Logo = link.Logo
		cnt.Style = link.Style
				
		siteLinkList = append(siteLinkList, cnt)
	}
	
	return siteLinkList
}

func (this *Model)GetArticleView(id int) (ArticleContent, bool) {
	cnt := ArticleContent{}
	
	article, found := content.QueryArticleById(id, this.dao)
	if found {
		cnt.Id = article.Id
		cnt.Title = article.Title
		cnt.Content = template.HTML(article.Content)
		cnt.CreateDate = article.CreateDate
		cnt.Catalog = article.Catalog.Name
		cnt.Author = article.Author.NickName
	}
	
	return cnt, found
}

