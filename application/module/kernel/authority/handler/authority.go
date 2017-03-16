package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateAuthorityHandler 新建AuthorityHandler
func CreateAuthorityHandler(modHub common.ModuleHub, cache cache.Cache) common.AuthorityHandler {
	i := impl{authData: cache, accountHandler: accountActionHandler{moduleHub: modHub, authCache: cache}}

	return &i
}

type impl struct {
	authData       cache.Cache
	accountHandler accountActionHandler
}

func (i *impl) LoginAccount(account, password string) (model.UserDetail, string, bool) {
	return i.accountHandler.LoginAccount(account, password)
}

func (i *impl) LogoutAccount(authID string) bool {
	return i.accountHandler.LogoutAccount(authID)
}
