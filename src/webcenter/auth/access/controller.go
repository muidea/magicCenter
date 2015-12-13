package access

import (
	"log"
    "webcenter/session"
    "webcenter/common"
    "webcenter/modelhelper"
    "webcenter/auth/account"
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


type accessController struct {
	
}


func (this *accessController)VerifyAction(param *VerifyParam) VerifyResult {
	result := VerifyResult{}
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = err.Error()
		return result
	}
	defer model.Release()
	
	user, found := account.QueryUserByAccount(model,param.account)
	if !found {
		result.ErrCode = 1
		result.Reason = "用户不存在"
		return result
	}
	
	session := param.session	
	if user.VerifyPassword(param.password) {
		result.ErrCode = 0
		result.Reason = "登陆成功"
		result.RedirectUrl = "/admin/"
		
		session.SetAccountId(user.Id)
		session.Save()
	} else {
		result.ErrCode = 1
		result.Reason = "账号或密码错误"
	}

	return result
}

