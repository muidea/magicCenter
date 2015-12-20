package account

import (
	"fmt"
	"webcenter/modelhelper"
	"webcenter/auth/group"
)

type UserInfo struct {
	Id int
	Account string
	NickName string
	Email string
	Group string
	Status int
}

type User struct {
	Id int
	Account string
	password string
	NickName string
	Email string
	Group int
	Status int
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
	userGroup, found := group.QueryGroupById(model, user.Group)
	if !found {
		return false;
	}
	
	return userGroup.AdminGroup()
}

func QueryAllUser(model modelhelper.Model) []UserInfo {
	userInfoList := []UserInfo{}
	sql := fmt.Sprintf("select u.id, u.account, u.nickname, u.email, g.name, u.status from user u, `group` g where u.group = g.id")
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		user := UserInfo{}
		model.GetValue(&user.Id, &user.Account, &user.NickName, &user.Email, &user.Group, &user.Status)
		
		userInfoList = append(userInfoList, user)
	}

	return userInfoList
}

func QueryUserByAccount(model modelhelper.Model, account string) (User,bool) {
	user := newUser()
	
	sql := fmt.Sprintf("select id,account,password,nickname,email, `group`, status from user where account='%s'", account)
	if !model.Query(sql) {
		panic("query failed")
	}
	
	result := false
	for model.Next() {
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group, &user.Status)
		result = true
		break
	}
	
	return user, result
}

func QueryUserById(model modelhelper.Model, id int) (User,bool) {
	user := newUser()
	
	sql := fmt.Sprintf("select id,account,password,nickname,email,`group`, status from user where id=%d", id)
	if !model.Query(sql) {
		panic("query failed")
	}
	
	result := false
	for model.Next() {
		model.GetValue(&user.Id, &user.Account, &user.password, &user.NickName, &user.Email, &user.Group, &user.Status)
		result = true
	}
	
	return user, result
}

func DeleteUser(model modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from user where id =%d", id)
	if !model.Execute(sql) {
		panic("execute failed")
	}
	
	return true
}

func SaveUser(model modelhelper.Model, user User) bool {
	sql := fmt.Sprintf("select id from user where id=%d", user.Id)
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into user(account,password,nickname,email,`group`,status) values ('%s', '%s', '%s', '%s', %d, %d)", user.Account, user.password, user.NickName, user.Email, user.Group, user.Status)
	} else {
		// modify
		sql = fmt.Sprintf("update user set account ='%s', password='%s', nickname='%s', email='%s', `group`=%d, status=%d where id =%d", user.Account, user.password, user.NickName, user.Email, user.Group, user.Status, user.Id)
	}
	
	result = model.Execute(sql)
	
	return result	
}

func QueryUserByGroup(model modelhelper.Model, id int) []UserInfo {
	userList := []UserInfo{}
	sql := fmt.Sprintf("select u.id, u.account, u.nickname, u.email, g.name, u.status from user u, `group` g where u.group = g.id and u.group=%d", id)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		user := UserInfo{}
		model.GetValue(&user.Id, &user.Account, &user.NickName, &user.Email, &user.Group, &user.Status)
		
		userList = append(userList, user)
	}
		
	return userList
}






