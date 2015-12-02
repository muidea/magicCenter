package account

import (
	"fmt"
	"log"
	"webcenter/modelhelper"
	"webcenter/auth/group"
)

type User struct {
	Id int
	Account string
	password string
	NickName string
	Email string
	Group int
}

func newUser() User {
	user := User{}
	user.Id = -1
	user.Group = -1
	
	return user
}

func (user User)VerifyPassword(password string) bool {
	return user.password == password
}

func IsAdmin(model modelhelper.Model, user User) bool {
	userGroup, found := group.GetGroupById(model, user.Group)
	if !found {
		return false;
	}
	
	return userGroup.AdminGroup()
}

func GetAllUser(model modelhelper.Model) []User {
	userList := []User{}
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user")
	if !model.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return userList
	}

	for model.Next() {
		user := newUser()
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group)
		
		userList = append(userList, user)
	}

	return userList
}

func QueryUserByAccount(model modelhelper.Model, account string) (User,bool) {
	user := newUser()
	
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where account='%s'", account)
	if !model.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return user, false
	}
	
	result := false
	for model.Next() {
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group)
		result = true
		break
	}
	
	return user, result
}

func QueryUserById(model modelhelper.Model, id int) (User,bool) {
	user := newUser()
	
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where id='%s'", id)
	if !model.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return user, false
	}
	
	result := false
	for model.Next() {
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group)
		result = true
	}
	
	return user, result
}

func DeleteUser(model modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from user where id =%d", id)
	if !model.Execute(sql) {
		log.Printf("delete user failed, sql:%s", sql)
		return false
	}
	
	return true
}

func SaveUser(model modelhelper.Model, user User) bool {
	sql := fmt.Sprintf("select id from user where id=%d", user.Id)
	if !model.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into user(account,password,nickname,email,`group`) values ('%s', '%s', '%s', '%s', %d)", user.Account, user.password, user.NickName, user.Email, user.Group)
	} else {
		// modify
		sql = fmt.Sprintf("update user set account ='%s', password='%s', nickname='%s', email='%s', `group`=%d where id =%d", user.Account, user.password, user.NickName, user.Email, user.Group, user.Id)
	}
	
	result = model.Execute(sql)
	
	return result	
}

func GetUserByGroup(model modelhelper.Model, id int) []User {
	userList := []User{}
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where `group`=%d", id)
	if !model.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return userList
	}

	for model.Next() {
		user := newUser()
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group)
		
		userList = append(userList, user)
	}
		
	return userList
}

func QueryDefaultUser(model modelhelper.Model) (User, bool) {
	user := newUser()
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group` from user where `group` in (select id from `group` where catalog = 0) order by id limit 1")
	if !model.Query(sql) {
		log.Printf("query user failed, sql:%s", sql)
		return user, false
	}
	
	result := false
	for model.Next() {
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group)
		result = true
	}
	
	return user, result
}






