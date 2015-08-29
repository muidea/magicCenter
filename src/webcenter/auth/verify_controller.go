package auth

import (
    "webcenter/common"
    "webcenter/session"    
)

type VerifyParam struct {
	account string
	password string
	accesscode string
}

type VerifyResult struct {
	common.Result
	RedirectUrl string
}

type verifyController struct {
}
 
func (this *verifyController)Action(param *VerifyParam, session *session.Session) VerifyResult {
	result := VerifyResult{}
	model, err := NewModel()
	if err != nil {
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
	
	if user.password == param.password {
		result.ErrCode = 0
		result.RedirectUrl = "/admin/"
	}
	
	session.SetOption(AccountSessionKey, user.Account)
	session.Save()
	
	return result
}
