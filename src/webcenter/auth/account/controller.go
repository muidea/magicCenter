package account

import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type GetAllUserParam struct {
	accessCode string	
}

type GetAllUserResult struct {
	common.Result
	User []User
}

type GetUserParam struct {
	accessCode string
	id int
}

type GetUserResult struct {
	common.Result
	User User
}

type DeleteUserParam struct {
	accessCode string
	id int
}

type DeleteUserResult struct {
	common.Result
}

type SubmitUserParam struct {
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

type accountController struct {
}


func (this *accountController)getAllUserAction(param GetAllUserParam) GetAllUserResult {
	result := GetAllUserResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	result.User = GetAllUser(model)
	result.ErrCode = 0

	return result
}

func (this *accountController)getUserAction(param GetUserParam) GetUserResult {
	result := GetUserResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	user, found := QueryUserById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.User = user
	}
	
	return result
}

func (this *accountController)deleteUserAction(param DeleteUserParam) DeleteUserResult {
	result := DeleteUserResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	DeleteUser(model, param.id)
	result.ErrCode = 0
	result.Reason = "删除用户成功"
	
	model.Release()
	
	return result
}

func (this *accountController)submitUserAction(param SubmitUserParam) SubmitUserResult {
	result := SubmitUserResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	user := newUser()
	user.Id = param.id
	user.Account = param.account
	user.password = param.password
	user.NickName = param.nickname
	user.Email = param.email
	user.Group = param.group
	
	if !SaveUser(model, user) {
		result.ErrCode = 1
		result.Reason = "保存用户信息失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存用户信息成功"
	}

	return result
}








