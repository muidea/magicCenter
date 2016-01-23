package admin

import (
	"log"
	"webcenter/util/modelhelper"
	"webcenter/util/session"
	"webcenter/kernel/admin/auth/account"
	"webcenter/kernel/admin/common"
)

type ManageParam struct {
	session *session.Session
}

type ManageResult struct {
	common.Result
	user account.User
}

type manageController struct {
}

func (this *manageController) ManageAction(param *ManageParam) ManageResult {
	result := ManageResult{}

	for true {
		accountId, found := param.session.GetAccountId()
		if !found {
			result.ErrCode = -1
			result.Reason = "当前未登陆"
			break
		}

		model, err := modelhelper.NewHelper()
		if err != nil {
			panic("new model failed")
		}

		user, found := account.QueryUserById(model, accountId)
		if !found || !account.IsAdmin(model, user) {
			log.Printf("found:%d", found)
			result.ErrCode = -1
			result.Reason = "非法账号"
			break
		}

		result.user = user
		result.ErrCode = 0
		break
	}

	return result
}
