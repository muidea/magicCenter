package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/authority/dal"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateAuthorityHandler 新建AuthorityHandler
func CreateAuthorityHandler(modHub common.ModuleHub) common.AuthorityHandler {
	i := impl{accountManager: accountManager{moduleHub: modHub}}

	return &i
}

type impl struct {
	accountManager accountManager
	cacheData      cache.Cache
}

func (i *impl) LoginAccount(account, password string) (model.UserDetail, string, bool) {
	user, ok := i.accountManager.FindUser(account, password)
	if !ok {
		return user, "", ok
	}

	token := i.cacheData.PutIn(user, cache.MaxAgeValue)
	return user, token, ok
}

func (i *impl) LogoutAccount(authToken string) bool {
	_, ok := i.cacheData.FetchOut(authToken)
	if !ok {
		return false
	}

	i.cacheData.Remove(authToken)
	return true
}

func (i *impl) VerifyAuth(res http.ResponseWriter, req *http.Request) bool {
	return true
}

func (i *impl) AdjustUserAuthGroup(userID int, authGroup []int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		return false
	}
	return dal.UpateUserAuthorityGroup(dbhelper, userID, authGroup)
}
