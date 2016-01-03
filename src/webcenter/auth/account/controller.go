package account

import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
	"webcenter/auth/group"
)

type QueryManageInfo struct {
	UserInfo []UserInfo
	GroupInfo []group.GroupInfo
}

type CheckAccountParam struct {
	accessCode string
	account string	
}

type CheckAccountResult struct {
	common.Result
}

type QueryAllUserParam struct {
	accessCode string	
}

type QueryAllUserResult struct {
	common.Result
	User []UserInfo
}

type QueryUserParam struct {
	accessCode string
	id int
}

type QueryUserResult struct {
	common.Result
	User User
}

type QueryUserByAccountParam struct {
	account string
}

type QueryUserByAccountResult struct {
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
	email string
	group string
}

type SubmitUserResult struct {
	common.Result
}

type SubmitVerifyInfoParam struct {
	accessCode string
	id int
	account string
	nickname string
	password string
}

type SubmitVerifyInfoResult struct {
	common.Result
}

type accountController struct {
}

func (this *accountController)queryManageInfoAction() QueryManageInfo {
	info := QueryManageInfo{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	info.UserInfo = QueryAllUser(model)
	info.GroupInfo = group.QueryAllGroup(model)

	return info
}

func (this *accountController)checkAccountAction(param CheckAccountParam) CheckAccountResult {
	result := CheckAccountResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	_ ,found := QueryUserByAccount(model, param.account)
	if !found {
		result.ErrCode = 0
		result.Reason ="该账号可用"
	} else {
		result.ErrCode = 1
		result.Reason ="该账号已经存在"
	}

	return result
}

func (this *accountController)queryAllUserAction(param QueryAllUserParam) QueryAllUserResult {
	result := QueryAllUserResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	result.User = QueryAllUser(model)
	result.ErrCode = 0

	return result
}

func (this *accountController)queryUserAction(param QueryUserParam) QueryUserResult {
	result := QueryUserResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
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

func (this *accountController)queryUserByAccountAction(param QueryUserByAccountParam) QueryUserByAccountResult {
	result := QueryUserByAccountResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
		
	user, found := QueryUserByAccount(model, param.account)
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
		panic("construct model failed")
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
		panic("construct model failed")
	}
	defer model.Release()

	if param.id == -1 {
		if createNewUser(model, param.account, param.email, param.group) {
			result.ErrCode = 0
			result.Reason = "新建用户成功，请到用户邮箱确认创建信息"
		} else {
			result.ErrCode = 1
			result.Reason = "新建用户失败，请稍后重试！"			
		}
	} else {
		usr, found := QueryUserById(model,param.id)
		if !found {
			result.ErrCode = 1
			result.Reason = "修改用户失败，指定用户不存在"
		} else {
			modMail := true
			modGroup := true
			
			model.BeginTransaction()
			
			if usr.Email != param.email {
				modMail = modifyUserMail(model, param.id, param.email)
			}
			
			modGroup = modifyUserGroup(model, param.id, param.group)
			
			if modMail && modGroup {
				result.ErrCode = 0
				result.Reason = "更新用户信息成功"
				model.Commit()
			} else {
				result.ErrCode = 1
				result.Reason = "更新用户信息失败"
				model.Rollback()
			}
		}
	}

	return result
}

func (this *accountController)submitVerifyInfoAction(param SubmitVerifyInfoParam) SubmitVerifyInfoResult {
	result := SubmitVerifyInfoResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	if updateUserInfo(model, param.id, param.nickname,param.password) {
		result.ErrCode = 0
		result.Reason = "校验用户成功"		
	} else {
		result.ErrCode = 1
		result.Reason = "校验用户失败"
	}

	return result
}








