package content


import (
	"log"
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
		log.Printf("illegal author ,author:%d", article.Author.Id)
		return false
	}
	
	if !article.Catalog.Query(this.dao) {
		log.Printf("illegal catalog ,catalog:%d", article.Catalog.Id)
		return false
	}
	
	return article.save(this.dao)
}

func (this *Model)GetAllCatalog() []Catalog {
	return GetAllCatalog(this.dao)
}


func (this *Model)GetCatalog(id int) (Catalog,bool) {
	catalog := newCatalog()
	catalog.Id = id
	
	result := catalog.Query(this.dao)
	return catalog,result
}

func (this *Model)DeleteCatalog(id int) {
	catalog := newCatalog()
	catalog.Id = id
	
	catalog.delete(this.dao)
}

func (this *Model)SaveCatalog(catalog Catalog) bool {
	if !catalog.Creater.Query(this.dao) {
		return false
	}
		
	return catalog.save(this.dao)
}

func (this *Model)QueryArticleByCatalog(id int) []ArticleInfo {
	return GetArticleByCatalog(id, this.dao)
}

