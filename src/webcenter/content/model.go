package content


import (
	"muidea.com/dao"
)

type Model struct {
	dao *dao.Dao
}

func NewModel()(Model, error) {
	model := Model{}
	
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		return model, err
	}
	
	model.dao = dao
	
	return model, err
}

func (this *Model)Release() {
	this.dao.Release()
} 

func (this *Model)GetAllArticleInfo() []ArticleInfo {
	return GetAllArticleInfo(this.dao)
}

func (this *Model)GetArticle(Id int) (Article, bool) {
	article := newArticle()
	article.Id = Id
	
	result := article.Query(this.dao)
	
	return article,result
}

func (this *Model)DeleteArticle(Id int) {
	article := newArticle()
	article.Id = Id
	
	article.delete(this.dao)
}


func (this *Model)SaveArticle(article Article) bool {
	if !article.Author.Query(this.dao) {
		return false
	}
	
	if !article.Catalog.Query(this.dao) {
		return false
	}
	
	return article.save(this.dao)
}

func (this *Model)GetAllCatalog() []Catalog {
	return GetAllCatalog(this.dao)
}


func (this *Model)GetCatalog(Id int) (Catalog,bool) {
	catalog := newCatalog()
	catalog.Id = Id
	
	result := catalog.Query(this.dao)
	return catalog,result
}

