package dal

import (
	"fmt"
	"magiccenter/common/model"
	"magiccenter/util/dbhelper"
	"strconv"
	"strings"
)

type tempPair struct {
	user   model.UserDetail
	groups string
}

//QueryAllUserList 查询全部用户列表
func QueryAllUserList(helper dbhelper.DBHelper) []model.User {
	userList := []model.User{}
	sql := fmt.Sprintf("select id,nickname from user")
	helper.Query(sql)

	for helper.Next() {
		user := model.User{}
		helper.GetValue(&user.ID, &user.Name)

		userList = append(userList, user)
	}

	return userList
}

//QueryAllUser 查询全部用户信息
func QueryAllUser(helper dbhelper.DBHelper) []model.UserDetail {
	userList := []model.UserDetail{}
	sql := fmt.Sprintf("select id,account,nickname,email,`group`, status from user")
	helper.Query(sql)

	tmpPairList := []tempPair{}
	for helper.Next() {
		groups := ""
		user := model.UserDetail{}
		helper.GetValue(&user.ID, &user.Account, &user.Name, &user.Email, &groups, &user.Status)

		tmp := tempPair{}
		tmp.user = user
		tmp.groups = groups

		tmpPairList = append(tmpPairList, tmp)
	}

	for _, tmp := range tmpPairList {
		groupArray := strings.Split(tmp.groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				tmp.user.Groups = append(tmp.user.Groups, gid)
			}
		}

		userList = append(userList, tmp.user)
	}

	return userList
}

// QueryUserByAccount 根据账号查询用户信息
func QueryUserByAccount(helper dbhelper.DBHelper, account string) (model.UserDetail, bool) {
	user := model.UserDetail{}

	sql := fmt.Sprintf("select id,account,nickname,email, `group`, status from user where account='%s'", account)
	helper.Query(sql)

	groups := ""
	result := false
	if helper.Next() {
		helper.GetValue(&user.ID, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		result = true
	}

	if result {
		groupArray := strings.Split(groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				user.Groups = append(user.Groups, gid)
			}
		}
	}

	return user, result
}

// VerifyUserByAccount 校验账号信息，如果账号信息正确，返回用户信息
func VerifyUserByAccount(helper dbhelper.DBHelper, account, password string) (model.UserDetail, bool) {
	user := model.UserDetail{}

	sql := fmt.Sprintf("select id,account,nickname,email, `group`, status from user where account='%s' and password='%s'", account, password)
	helper.Query(sql)

	groups := ""
	result := false
	if helper.Next() {
		helper.GetValue(&user.ID, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		result = true
	}

	if result {
		groupArray := strings.Split(groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				user.Groups = append(user.Groups, gid)
			}
		}
	}

	return user, result
}

// QueryUserByID 根据用户ID查询用户信息
func QueryUserByID(helper dbhelper.DBHelper, id int) (model.UserDetail, bool) {
	user := model.UserDetail{}

	sql := fmt.Sprintf("select id,account,nickname,email,`group`, status from user where id=%d", id)
	helper.Query(sql)

	groups := ""
	result := false
	if helper.Next() {
		helper.GetValue(&user.ID, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		result = true
	}

	if result {
		groupArray := strings.Split(groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				user.Groups = append(user.Groups, gid)
			}
		}
	}

	return user, result
}

// DeleteUser 删除用户，根据用户ID
func DeleteUser(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from user where id =%d", id)
	_, ret := helper.Execute(sql)
	return ret
}

// DeleteUserByAccount 删除用户，根据用户账号&密码
func DeleteUserByAccount(helper dbhelper.DBHelper, account, password string) bool {
	sql := fmt.Sprintf("delete from user where account ='%s' and password='%s'", account, password)
	_, ret := helper.Execute(sql)
	return ret
}

// CreateUser 创建新用户，根据用户信息和密码
func CreateUser(helper dbhelper.DBHelper, user model.UserDetail, password string) bool {
	groups := ""
	for _, g := range user.Groups {
		groups = fmt.Sprintf("%s%d,", groups, g)
	}
	groups = groups[0 : len(groups)-1]

	// insert
	sql := fmt.Sprintf("insert into user(account,password,nickname,email,`group`,status) values ('%s', '%s', '%s', '%s', '%s', %d)", user.Account, password, user.Name, user.Email, groups, user.Status)
	_, result := helper.Execute(sql)

	return result
}

// SaveUser 保存用户信息
func SaveUser(helper dbhelper.DBHelper, user model.UserDetail) bool {
	groups := ""
	for _, g := range user.Groups {
		groups = fmt.Sprintf("%s%d,", groups, g)
	}
	groups = groups[0 : len(groups)-1]

	// modify
	sql := fmt.Sprintf("update user set nickname='%s', email='%s', `group`='%s', status=%d where id =%d", user.Name, user.Email, groups, user.Status, user.ID)
	_, result := helper.Execute(sql)

	return result
}

// QueryUserByGroup 查询指定分组下的用户信息
func QueryUserByGroup(helper dbhelper.DBHelper, id int) []model.UserDetail {
	userList := []model.UserDetail{}
	sql := fmt.Sprintf("select id,account,nickname,email,`group`, status from user where `group` like '%d' union select id,account,nickname,email,`group`, status from user where `group` like '%%,%d' union select id,account,nickname,email,`group`, status from user where `group` like '%d,%%' union select id,account,nickname,email,`group`, status from user where `group` like '%%,%d,%%'", id, id, id, id)
	helper.Query(sql)

	tmpPairList := []tempPair{}
	for helper.Next() {
		groups := ""
		user := model.UserDetail{}
		helper.GetValue(&user.ID, &user.Account, &user.Name, &user.Email, &groups, &user.Status)

		tmp := tempPair{}
		tmp.user = user
		tmp.groups = groups

		tmpPairList = append(tmpPairList, tmp)
	}

	for _, tmp := range tmpPairList {
		groupArray := strings.Split(tmp.groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				tmp.user.Groups = append(tmp.user.Groups, gid)
			}
		}

		userList = append(userList, tmp.user)
	}

	return userList
}
