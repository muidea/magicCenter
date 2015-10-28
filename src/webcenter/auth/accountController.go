package auth

import (
	"log"
    "webcenter/session"
	"webcenter/common"
)


type GetAllAccountInfoParam struct {
	session *session.Session
	accessCode string	
}

type GetAllAccountInfoResult struct {
	common.Result
	User []User
	Group []Group
}

type GetAllUserParam struct {
	session *session.Session
	accessCode string	
}

type GetAllUserResult struct {
	common.Result
	User []User
}

type GetUserParam struct {
	session *session.Session
	accessCode string
	id int
}

type GetUserResult struct {
	common.Result
	User User
}

type DeleteUserParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteUserResult struct {
	common.Result
}

type GetGroupParam struct {
	session *session.Session
	accessCode string
	id int
}

type GetGroupResult struct {
	common.Result
	Group Group
}


type DeleteGroupParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteGroupResult struct {
	common.Result
}

type GetAllGroupParam struct {
	session *session.Session
	accessCode string	
}

type GetAllGroupResult struct {
	common.Result
	Group []Group
}


type SubmitUserParam struct {
	session *session.Session
	accessCode string
	id int
	account string
	password string
	nickname string
	email string
	group int
	submitDate string	
}

type SubmitUserResult struct {
	common.Result
}

type SubmitGroupParam struct {
	session *session.Session
	accessCode string
	id int
	name string
	parent int
	submitDate string	
}

type SubmitGroupResult struct {
	common.Result
}

type accountController struct {
}

func (this *accountController)getAllAccountInfoAction(param GetAllAccountInfoParam) GetAllAccountInfoResult {
	result := GetAllAccountInfoResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	result.User = model.GetAllUser()
	result.Group = model.GetAllGroup()
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *accountController)getAllUserAction(param GetAllUserParam) GetAllUserResult {
	result := GetAllUserResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	result.User = model.GetAllUser()
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *accountController)getUserAction(param GetUserParam) GetUserResult {
	result := GetUserResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, ok := session.GetAccount()
	if !ok {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	user, found := model.GetUser(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.User = user
	}
	
	model.Release()
	
	return result
}

func (this *accountController)deleteUserAction(param DeleteUserParam) DeleteUserResult {
	result := DeleteUserResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	if user.Id == param.id {
		result.ErrCode = 1
		result.Reason = "不允许删除当前用户"
		return result
	}

	model.DeleteUser(param.id)
	result.ErrCode = 0
	result.Reason = "删除用户成功"
	
	model.Release()
	
	return result
}
 
func (this *accountController)getAllGroupAction(param GetAllGroupParam) GetAllGroupResult {
	result := GetAllGroupResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	result.Group = model.GetAllGroup()
	result.ErrCode = 0
	model.Release()
	
	return result
}

 
func (this *accountController)getGroupAction(param GetGroupParam) GetGroupResult {
	result := GetGroupResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	catalog, found := model.GetGroup(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Group = catalog
	}

	model.Release()

	return result
}

func (this *accountController)deleteGroupAction(param DeleteGroupParam) DeleteGroupResult {
	result := DeleteGroupResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	userInfoList := model.QueryUserByGroup(param.id)
	subGroupList := model.QuerySubGroup(param.id)
	if (len(userInfoList) >0 || len(subGroupList) >0) {
		result.ErrCode = 1
		result.Reason = "该分组被引用，无法立即删除"
		return result
	}
	
	model.DeleteGroup(param.id)
	result.ErrCode = 0
	result.Reason = "删除分组成功"
	
	model.Release()
	
	return result
}


func (this *accountController)submitUserAction(param SubmitUserParam) SubmitUserResult {
	result := SubmitUserResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	user = NewUser()
	user.Id = param.id
	user.Account = param.account
	user.password = param.password
	user.NickName = param.nickname
	user.Email = param.email
	user.Group.Id = param.group
	
	if !model.SaveUser(user) {
		result.ErrCode = 1
		result.Reason = "保存用户信息失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存用户信息成功"
	}
	
	model.Release()

	return result
}

func (this *accountController)submitGroupAction(param SubmitGroupParam) SubmitGroupResult {
	result := SubmitGroupResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := model.FindUserByAccount(account)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	group := newGroup()
	group.Id = param.id
	group.Name = param.name
	group.Catalog = param.parent
	
	if !model.SaveGroup(group) {
		result.ErrCode = 1
		result.Reason = "保存分组失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存分组成功"
	}
	
	model.Release()

	return result
}


