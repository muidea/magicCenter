package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
)

// CreateAuthorityHandler 新建AuthorityHandler
func CreateAuthorityHandler(modHub common.ModuleHub) common.AuthorityHandler {
	i := impl{accountHandler: accountActionHandler{moduleHub: modHub}}

	return &i
}

type impl struct {
	accountHandler accountActionHandler
}

func (i *impl) LoginAccount(account, password string) (model.UserDetail, string, bool) {
	return i.accountHandler.LoginAccount(account, password)
}

func (i *impl) LogoutAccount(authID string) bool {
	return i.accountHandler.LogoutAccount(authID)
}
