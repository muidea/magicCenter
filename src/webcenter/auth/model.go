package auth

import (
	"muidea.com/dao"
)

type Model struct {
	dao *dao.Dao
}

var AccountSessionKey string = "account"

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

func (this *Model)FindUserByAccount(account string) (User, bool){
	user := User{}
	user.Id = 0
	user.Account = account
	found := user.Query(this.dao)
	return user, found
}