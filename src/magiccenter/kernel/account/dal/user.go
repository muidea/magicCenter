package dal

import (
	"fmt"
	"strings"
	"strconv"	
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/account/model"
)

type tempPair struct {
	user model.UserDetailView
	groups string
}

func QueryAllUser(helper modelhelper.Model) []model.UserDetailView {
	userList := []model.UserDetailView{}
	sql := fmt.Sprintf("select id,account,nickname,email,`group`, status from user")
	helper.Query(sql)

	tmpPairList := []tempPair{}
	for helper.Next() {
		groups := ""
		user := model.UserDetailView{}
		helper.GetValue(&user.Id, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		
		tmp := tempPair{}
		tmp.user = user
		tmp.groups = groups
		
		tmpPairList = append(tmpPairList, tmp)
	}
	
	for i, _ := range tmpPairList {
		tmp := &tmpPairList[i]
		groupArray := strings.Split(tmp.groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				group, found := QueryGroupById(helper, gid)
				if found {					
					tmp.user.Groups = append(tmp.user.Groups, group)
				}
			}
		}		
		
		userList = append(userList, tmp.user)
	}
	
	return userList
}

func QueryUserByAccount(helper modelhelper.Model, account string) (model.UserDetailView,bool) {
	user := model.UserDetailView{}
	
	sql := fmt.Sprintf("select id,account,nickname,email, `group`, status from user where account='%s'", account)
	helper.Query(sql)
	
	groups := ""
	result := false
	if helper.Next() {
		helper.GetValue(&user.Id, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		result = true
	}
	
	if result {
		groupArray := strings.Split(groups,",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				group, found := QueryGroupById(helper, gid)
				if found {					
					user.Groups = append(user.Groups, group)
				}
			}
		}		
	}
		
	return user, result
}

func VerifyUserByAccount(helper modelhelper.Model, account, password string) (model.UserDetailView,bool) {
	user := model.UserDetailView{}
	
	sql := fmt.Sprintf("select id,account,nickname,email, `group`, status from user where account='%s' and password='%s'", account, password)
	helper.Query(sql)
	
	groups := ""
	result := false
	if helper.Next() {
		helper.GetValue(&user.Id, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		result = true
	}
	
	if result {
		groupArray := strings.Split(groups,",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				group, found := QueryGroupById(helper, gid)
				if found {					
					user.Groups = append(user.Groups, group)
				}
			}
		}		
	}
		
	return user, result
}

func QueryUserById(helper modelhelper.Model, id int) (model.UserDetailView, bool) {
	user := model.UserDetailView{}
	
	sql := fmt.Sprintf("select id,account,nickname,email,`group`, status from user where id=%d", id)
	helper.Query(sql)
	
	groups := ""
	result := false
	if helper.Next() {
		helper.GetValue(&user.Id, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		result = true
	}
	
	if result {
		groupArray := strings.Split(groups,",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				group, found := QueryGroupById(helper, gid)
				if found {					
					user.Groups = append(user.Groups, group)
				}
			}
		}		
	}
		
	return user, result
}

func DeleteUser(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from user where id =%d", id)
	_ ,ret := helper.Execute(sql)
	return ret
}

func DeleteUserByAccount(helper modelhelper.Model, account, password string) bool {
	sql := fmt.Sprintf("delete from user where account ='%s' and password='%s'", account, password)
	_, ret := helper.Execute(sql) 
	return ret
}


func SaveUser(helper modelhelper.Model, user model.UserDetail) bool {
	sql := fmt.Sprintf("select id from user where id=%d", user.Id)
	helper.Query(sql)

	result := false;
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}
	
	groups := ""
	for _, g := range user.Groups {
		groups = fmt.Sprintf("%s%d,", groups, g)
	}
	groups = groups[0:len(groups)-1]

	if !result {
		// insert
		sql = fmt.Sprintf("insert into user(account,password,nickname,email,`group`,status) values ('%s', '%s', '%s', '%s', '%s', %d)", user.Account, user.Password, user.Name, user.Email, groups, user.Status)
	} else {
		// modify
		sql = fmt.Sprintf("update user set password='%s' nickname='%s', email='%s', `group`='%s', status=%d where id =%d", user.Password, user.Name, user.Email, groups, user.Status, user.Id)
	}
	
	_, result = helper.Execute(sql)
	
	return result
}

func QueryUserByGroup(helper modelhelper.Model, id int) []model.UserDetailView {
	userList := []model.UserDetailView{}
	sql := fmt.Sprintf("select id,account,nickname,email,`group`, status from user where `group` like '%d' union select id,account,nickname,email,`group`, status from user where `group` like '%%,%d' union select id,account,nickname,email,`group`, status from user where `group` like '%d,%%' union select id,account,nickname,email,`group`, status from user where `group` like '%%,%d,%%'", id,id,id,id)
	helper.Query(sql)

	tmpPairList := []tempPair{}
	for helper.Next() {
		groups := ""
		user := model.UserDetailView{}
		helper.GetValue(&user.Id, &user.Account, &user.Name, &user.Email, &groups, &user.Status)
		
		tmp := tempPair{}
		tmp.user = user
		tmp.groups = groups
		
		tmpPairList = append(tmpPairList, tmp)
	}
	
	for i, _ := range tmpPairList {
		tmp := &tmpPairList[i]
		groupArray := strings.Split(tmp.groups, ",")
		for _, g := range groupArray {
			gid, err := strconv.Atoi(g)
			if err == nil {
				group, found := QueryGroupById(helper, gid)
				if found {					
					tmp.user.Groups = append(tmp.user.Groups, group)
				}
			}
		}		
		
		userList = append(userList, tmp.user)
	}

	return userList
}




