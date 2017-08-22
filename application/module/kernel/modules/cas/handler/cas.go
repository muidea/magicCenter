package handler

import (
	"net"
	"strings"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/util"
)

// CreateCASHandler 新建CASHandler
func CreateCASHandler(moduleHub common.ModuleHub) common.CASHandler {
	manager, _ := createAccountManager(moduleHub)
	i := impl{
		accountManager: manager,
		token2IDMap:    make(map[string]int)}

	return &i
}

type impl struct {
	accountManager accountManager
	token2IDMap    map[string]int
}

func (i *impl) allocAuthToken() string {
	return strings.ToLower(util.RandomAlphanumeric(32))
}

func (i *impl) getIP(remoteAddr string) string {
	ip := "127.0.0.1"

	addr, err := net.ResolveIPAddr("ip4", remoteAddr)
	if err != nil {
		return ip
	}
	ip = addr.IP.String()

	return ip
}

func (i *impl) LoginAccount(account, password, remoteAddr string) (model.UserDetail, string, bool) {
	user, ok := i.accountManager.userLogin(account, password, i.getIP(remoteAddr))
	if !ok {
		return user, "", ok
	}

	token := i.allocAuthToken()
	i.token2IDMap[token] = user.ID

	return user, token, ok
}

func (i *impl) LoginToken(authToken, remoteAddr string) (string, bool) {
	return authToken, true
}

func (i *impl) Logout(authToken, remoteAddr string) bool {
	id, ok := i.token2IDMap[authToken]
	if ok {
		i.accountManager.userLogout(id, i.getIP(remoteAddr))
	}

	return ok
}

func (i *impl) RefreshToken(authToken, remoteAddr string) bool {
	id, ok := i.token2IDMap[authToken]
	if ok {
		i.accountManager.userRefresh(id, i.getIP(remoteAddr))
	}

	return ok
}

func (i *impl) VerifyToken(authToken string) (model.OnlineAccountInfo, bool) {
	var info model.OnlineAccountInfo
	id, ok := i.token2IDMap[authToken]
	if ok {
		info, ok = i.accountManager.userVerify(id)
		if !ok {
			delete(i.token2IDMap, authToken)
		}
	}
	return info, ok
}

func (i *impl) AllocStaticToken(id string, expiration int64) (string, bool) {
	return "", true
}
