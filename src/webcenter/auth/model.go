package auth

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

func (this *Model)FindUserByAccount(account string) (User, bool){
	user := NewUser()
	user.Id = 0
	user.Account = account
	found := user.Query(this.dao)
	return user, found
}


func (this *Model)GetAllUser() []User {
	return GetAllUser(this.dao)
}

func (this *Model)GetUser(Id int) (User, bool) {
	user := NewUser()
	user.Id = Id
	
	result := user.Query(this.dao)
	
	return user,result
}

func (this *Model)DeleteUser(Id int) {
	user := NewUser()
	user.Id = Id
	
	user.delete(this.dao)
}


func (this *Model)SaveUser(user User) bool {
	if !user.Group.query(this.dao) {
		log.Printf("group isn't exist, gid:%d", user.Group.Id)
		
		return false
	}
	
	return user.save(this.dao)
}

func (this *Model)GetAllGroup() []Group {
	return GetAllGroup(this.dao)
}


func (this *Model)GetGroup(id int) (Group,bool) {
	group := newGroup()
	group.Id = id
	
	result := group.query(this.dao)
	return group,result
}

func (this *Model)DeleteGroup(id int) {
	group := newGroup()
	group.Id = id
	
	group.delete(this.dao)
}

func (this *Model)SaveGroup(group Group) bool{
	if group.Catalog > 0 {
		pGroup := newGroup()
		pGroup.Id = group.Catalog
		if !pGroup.query(this.dao) {
			return false
		}		
	}
	
	return group.save(this.dao)
}

func (this *Model)QueryUserByGroup(id int) []User {
	return GetUserByGroup(id, this.dao)
}

func (this *Model)QuerySubGroup(id int) []Group {
	return GetAllSubGroup(id, this.dao)
}


