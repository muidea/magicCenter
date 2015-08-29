package auth

import (
	"fmt"
	"log"
	"muidea.com/dao"
)

type User struct {
	Id int
	Account string
	password string
	NickName string
	Email string
	Group Group
}

func NewUser() User {
	user := User{}
	user.Id = -1
	user.Group = newGroup()
	
	return user
}

func GetAllUser(dao * dao.Dao) []User {
	userList := []User{}
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user")
	if !dao.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return userList
	}

	for dao.Next() {
		user := NewUser()
		dao.GetField(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group.Id)
		
		user.Group.query(dao)
		
		userList = append(userList, user)
	}
	
	return userList
}

func (this *User)IsAdmin() bool {
	return this.Group.IsAdminGroup()
}

func (this *User)queryById(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return false
	}

	result := false
	for dao.Next() {
		dao.GetField(&this.Id, &this.Account, &this.password, &this.NickName, &this.Email, &this.Group.Id)
		result = true
	}
	
	if result {
		result = this.Group.query(dao)		
	}
	
	return result
}

func (this *User)queryByAccount(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where account='%s'", this.Account)
	if !dao.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return false
	}
	
	result := false
	for dao.Next() {
		dao.GetField(&this.Id, &this.Account, &this.password, &this.NickName, &this.Email, &this.Group.Id)
		result = true
	}
	
	if result {
		result = this.Group.query(dao)
	}
	
	return result
}

func (this *User)Query(dao *dao.Dao) bool {
	if this.Id > 0 {
		return this.queryById(dao)
	} else {
		return this.queryByAccount(dao)
	}
}

func (this *User)insert(dao *dao.Dao) bool {
	if !this.Group.inert(dao) {
		return false
	}
	
	sql := fmt.Sprintf("insert into user value (%d, %s, %s, %s, %s, %d)", this.Id, this.Account, this.password, this.NickName, this.Email, this.Group.Id)
	if !dao.Execute(sql) {
		log.Printf("inser user failed, sql:%s", sql)
		return false
	}
		
	return true	
}

func (this *User)update(dao *dao.Dao) bool {
	sql := fmt.Sprintf("update user set account ='%s', password='%s', niciname='%s', email='%s', `group`=%d where id =%d", this.Account, this.password, this.NickName, this.Email, this.Group.Id, this.Id)
	if !dao.Execute(sql) {
		log.Printf("update user failed, sql:%s", sql)
		return false
	}
	
	return true
}

func (this *User)remove(dao *dao.Dao) {
	this.Group.remove(dao)
	
	sql := fmt.Sprintf("delete from user where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete user failed, sql:%s", sql)
		return
	}	
}



