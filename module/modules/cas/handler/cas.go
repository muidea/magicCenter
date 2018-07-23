package handler

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

const userPrefix = "user"
const endpointPrefix = "endpoint"

// CreateCASHandler 新建CASHandler
func CreateCASHandler(moduleHub common.ModuleHub) common.CASHandler {
	accountManager, _ := createAccountManager(moduleHub)
	endpointManager, _ := createEndpointManager(moduleHub)

	i := impl{
		accountManager:  accountManager,
		endpointManager: endpointManager,
		onlineUser:      make(map[string]model.AccountOnlineView)}

	return &i
}

type impl struct {
	accountManager  accountManager
	endpointManager endpointManager
	onlineUser      map[string]model.AccountOnlineView
}

func (i *impl) getTokenPrefix(authToken string) (string, error) {
	items := strings.Split(authToken, ".")
	if len(items) != 2 {
		msg := fmt.Sprintf("illegal authToken,[%s]", authToken)
		return "", errors.New(msg)
	}
	if items[0] != userPrefix && items[0] != endpointPrefix {
		msg := fmt.Sprintf("illegal authToken,[%s]", authToken)
		return "", errors.New(msg)
	}

	return items[0], nil
}

func (i *impl) allocAuthToken(prefix string) string {
	return fmt.Sprintf("%s.%s", prefix, strings.ToLower(util.RandomAlphanumeric(32)))
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

func (i *impl) LoginAccount(account, password, remoteAddr string) (model.AccountOnlineView, bool) {
	token := i.allocAuthToken(userPrefix)
	onlineUser, ok := i.accountManager.userLogin(account, password, i.getIP(remoteAddr), token)
	if ok {
		i.onlineUser[onlineUser.AuthToken] = onlineUser
	}

	return onlineUser, ok
}

func (i *impl) LoginEndpoint(identifyID, authToken, remoteAddr string) (model.AccountOnlineView, bool) {
	token := i.allocAuthToken(endpointPrefix)
	onlineUser, ok := i.endpointManager.endpointLogin(identifyID, authToken, i.getIP(remoteAddr), token)
	if ok {
		i.onlineUser[onlineUser.AuthToken] = onlineUser
	}

	return onlineUser, ok
}

func (i *impl) Logout(authToken, remoteAddr string) bool {
	delete(i.onlineUser, authToken)

	return true
}

func (i *impl) VerifyToken(authToken string) (model.AccountOnlineView, bool) {
	info, ok := i.onlineUser[authToken]
	return info, ok
}

func (i *impl) RefreshToken(authToken, remoteAddr string) bool {
	info, ok := i.onlineUser[authToken]
	if ok {
		info.UpdateTime = time.Now().Unix()
		i.onlineUser[authToken] = info
	} else {
		log.Printf("illegal authToken[%s] refresh, not login, address:%s", authToken, remoteAddr)
	}

	return ok
}

func (i *impl) GetSummary() model.CasSummary {
	result := model.CasSummary{}
	onlineItem := model.UnitSummary{Name: "在线", Type: "online", Count: len(i.onlineUser)}
	result = append(result, onlineItem)

	return result
}
