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

func (this *Model)GetAllCatalog() []Catalog {
	return GetAllCatalog(this.dao)
}

