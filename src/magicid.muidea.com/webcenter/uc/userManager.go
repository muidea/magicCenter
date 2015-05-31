package uc

import (
	"fmt"
	"log"
	"muidea.com/dao"
)

type userManager struct {
	userInfo map[int]User
}

func (this * userManager) Load() bool {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		log.Printf("fetch dao failed, err:%s", err.Error())
		return false
	}
	defer dao.Release()
	
	sql := fmt.Sprintf("%s","select id,name,password,email,group from magicid_db.user")
	if !dao.Query(sql) {
		return false
	}
	
	this.userInfo = make(map[int]User)
	for dao.Next() {
		user := User{}
		dao.GetField(&user.id, &user.name, &user.password, &user.email, &user.group)
		this.userInfo[user.id] = user
	}
	
	return true
}

func (this * userManager) AddUser(user User) bool {
	
	return true
}

func (this * userManager) ModUser(user User) bool {
	return true
}

func (this * userManager) DelUser(id int) {
	
}

func (this * userManager) FindUserById(id int) User {
	user := User{}
	return user
}

func (this * userManager) FindUserByName(name string) User {
	user := User{}
	return user	
}
