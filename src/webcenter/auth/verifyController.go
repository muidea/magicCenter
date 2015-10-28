package auth

import (
	"log"
	"webcenter/common"
    "webcenter/session"    
)

type VerifyParam struct {
	account string
	password string
	accesscode string
	session *session.Session
}

type VerifyResult struct {
	common.Result
	RedirectUrl string
}

type verifyController struct {
}
 
func (this *verifyController)Action(param *VerifyParam) VerifyResult {
	result := VerifyResult{}
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = err.Error()
		return result
	}
	defer model.Release()
	
	user, found := model.FindUserByAccount(param.account)
	if !found {
		result.ErrCode = 1
		result.Reason = "用户不存在"
		return result
	}
	
	session := param.session
	if user.password == param.password {
		result.ErrCode = 0
		result.Reason = "登陆成功"
		result.RedirectUrl = "/admin/"
		
		session.SetAccount(user.Account)
		session.Save()
	} else {
		result.ErrCode = 1
		result.Reason = "账号或密码错误"
	}
		
	return result
}
