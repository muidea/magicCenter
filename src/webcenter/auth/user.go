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
		
		userList = append(userList, user)
	}
	
	for i:=0; i < len(userList); i++ {
		user := &userList[i]
		
		user.Group.query(dao)
	}
	
	return userList
}


func GetUserByGroup(id int, dao* dao.Dao) []User {
	userList := []User{}
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where `group`=%d", id)
	if !dao.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return userList
	}

	for dao.Next() {
		user := NewUser()
		dao.GetField(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group.Id)
		
		userList = append(userList, user)
	}
	
	for i:=0; i < len(userList); i++ {
		user := &userList[i]
		
		user.Group.query(dao)
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


func (this *User)delete(dao *dao.Dao) {
	this.Group.delete(dao)
	
	sql := fmt.Sprintf("delete from user where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete user failed, sql:%s", sql)
		return
	}	
}

func (this *User)save(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id from user where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return false
	}

	result := false;
	for dao.Next() {
		var id = 0
		result = dao.GetField(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into user(account,password,nickname,email,`group`) values ('%s', '%s', '%s', '%s', %d)", this.Account, this.password, this.NickName, this.Email, this.Group.Id)
	} else {
		// modify
		sql = fmt.Sprintf("update user set account ='%s', password='%s', nickname='%s', email='%s', `group`=%d where id =%d", this.Account, this.password, this.NickName, this.Email, this.Group.Id, this.Id)
	}
	
	result = dao.Execute(sql)
	
	return result	
}



