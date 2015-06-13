package datamanager

import (
	"fmt"
	"log"
	"muidea.com/dao"
)

type UserManager struct {
	userInfo map[int]User
	dao *dao.Dao
}

func (this *UserManager) Load() bool {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		log.Printf("fetch dao failed, err:%s", err.Error())
		return false
	}
	
	this.userInfo = make(map[int]User)
	this.dao = dao	
	return true
}

func (this *UserManager) Unload() {
	this.dao.Release()
	this.dao = nil
	this.userInfo = nil
}

func (this * UserManager) AddUser(user User) bool {
	sql := fmt.Sprintf("insert into magicid_db.user value (%d, %s, %s, %s, %d)", user.id, user.name, user.password, user.email, user.group)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}
		
	return true
}

func (this * UserManager) ModUser(user User) bool {
	sql := fmt.Sprintf("update magicid_db.user set name ='%s', password='%s', group=%d where id =%d", user.name, user.password, user.group, user.id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}
	
	this.userInfo[user.id] = user	
	return true
}

func (this * UserManager) DelUser(id int) {	
	delete(this.userInfo, id)
	
	sql := fmt.Sprintf("delete from magicid_db.user where id =%d", id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return 
	}
}

func (this * UserManager) DelUserByGroup(group int) {
	for id, user := range this.userInfo {
		if user.group == group {
			delete(this.userInfo, id)
		}
	}
	
	sql := fmt.Sprintf("delete from magicid_db.user where group =%d", group)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return 
	}
}


func (this * UserManager) FindUserById(id int) (User, bool) {
	user, found := this.userInfo[id]
	if !found {
		sql := fmt.Sprintf("select * from magicid_db.user where id=%d", id)
		if !this.dao.Query(sql) {
			log.Printf("query failed, sql:%s", sql)
			return user, false
		}
	
		for this.dao.Next() {
			user := User{}
			this.dao.GetField(&user.id, &user.name, &user.password, &user.email, &user.group)
			this.userInfo[user.id] = user
		}
	
	}
	user, found = this.userInfo[id]
	
	return user, found
}

func (this * UserManager) FindUserByEMail(email string) (User, bool) {
	user := User{}
	sql := fmt.Sprintf("select * from magicid_db.user where email='%s'", email)
	if !this.dao.Query(sql) {
		log.Printf("query failed, sql:%s", sql)
		return user, false
	}
	
	found := false
	for this.dao.Next() {
		this.dao.GetField(&user.id, &user.name, &user.password, &user.email, &user.group)
		this.userInfo[user.id] = user
		found = true
	}
	
	return user, found	
}
