package handler

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCommon/foundation/util"
	"github.com/muidea/magicCommon/model"
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
		onlineEntry:     make(map[string]model.OnlineEntryView)}

	return &i
}

type impl struct {
	accountManager  accountManager
	endpointManager endpointManager
	onlineEntry     map[string]model.OnlineEntryView
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

func (i *impl) LoginAccount(account, password, remoteAddr string) (model.OnlineEntryView, string, bool) {
	token := i.allocAuthToken(userPrefix)
	onlineEntry, ok := i.accountManager.userLogin(account, password, i.getIP(remoteAddr))
	if ok {
		i.onlineEntry[token] = onlineEntry
	}

	return onlineEntry, token, ok
}

func (i *impl) ChangeAccountPassword(accountID int, oldPassword, newPassword string) bool {
	return i.accountManager.userChangePassword(accountID, oldPassword, newPassword)
}

func (i *impl) LoginEndpoint(identifyID, authToken, remoteAddr string) (model.OnlineEntryView, string, bool) {
	token := i.allocAuthToken(endpointPrefix)
	onlineEntry, ok := i.endpointManager.endpointLogin(identifyID, authToken, i.getIP(remoteAddr))
	if ok {
		i.onlineEntry[token] = onlineEntry
	}

	return onlineEntry, token, ok
}

func (i *impl) Logout(authToken, remoteAddr string) bool {
	delete(i.onlineEntry, authToken)

	return true
}

func (i *impl) VerifyToken(authToken string) (model.OnlineEntryView, string, bool) {
	info, ok := i.onlineEntry[authToken]
	return info, authToken, ok
}

func (i *impl) RefreshToken(authToken, remoteAddr string) bool {
	info, ok := i.onlineEntry[authToken]
	if ok {
		info.UpdateTime = time.Now().Unix()
		i.onlineEntry[authToken] = info
	} else {
		log.Printf("illegal authToken[%s] refresh, not login, address:%s", authToken, remoteAddr)
	}

	return ok
}

func (i *impl) GetSummary() model.CasSummary {
	result := model.CasSummary{}
	onlineItem := model.UnitSummary{Name: "在线", Type: "online", Count: len(i.onlineEntry)}
	result = append(result, onlineItem)

	return result
}
