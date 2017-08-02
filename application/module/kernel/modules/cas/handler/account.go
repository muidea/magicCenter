package handler

import (
	"log"
	"time"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
)

type onlineAccountInfo struct {
	id         int    // User ID
	loginTime  int64  // 登陆时间
	updateTime int64  // 更新时间
	address    string // 访问IP
}

type accountManager struct {
	accountHandler common.AccountHandler
	onlineUser     map[int]onlineAccountInfo
}

func createAccountManager(moduleHub common.ModuleHub) (accountManager, bool) {
	s := accountManager{accountHandler: nil, onlineUser: make(map[int]onlineAccountInfo)}
	mod, ok := moduleHub.FindModule(common.AccountModuleID)
	if !ok {
		return s, false
	}

	entryPoint := mod.EntryPoint()
	switch entryPoint.(type) {
	case common.AccountHandler:
		s.accountHandler = entryPoint.(common.AccountHandler)
	}

	return s, s.accountHandler != nil
}

// model.UserDetail 登陆用户
// bool 是否登陆成功
func (s *accountManager) userLogin(account, password, remoteAddr string) (model.UserDetail, bool) {
	user, ok := s.accountHandler.FindUserByAccount(account, password)
	if !ok {
		return user, ok
	}

	info, ok := s.onlineUser[user.ID]
	if ok {
		if info.address != remoteAddr {
			log.Printf("duplicate user[%d] logining,pre address:%s, cur address:%s", info.id, info.address, remoteAddr)
		}
	}

	info = onlineAccountInfo{}
	info.id = user.ID
	info.address = remoteAddr
	info.loginTime = time.Now().Unix()
	info.updateTime = info.loginTime
	s.onlineUser[user.ID] = info
	return user, true
}

func (s *accountManager) userRefresh(id int, remoteAddr string) {
	info, ok := s.onlineUser[id]
	if ok {
		if info.address != remoteAddr {
			log.Printf("illegal user[%d] refresh, pre address:%s, cur address:%s", id, info.address, remoteAddr)
		}
		info.updateTime = time.Now().Unix()
		s.onlineUser[id] = info
	} else {
		log.Printf("illegal user[%d] refresh, not login, address:%s", id, remoteAddr)
	}
}

func (s *accountManager) userLogout(id int, remoteAddr string) {
	delete(s.onlineUser, id)
}

func (s accountManager) userList() []int {
	userList := []int{}
	for k := range s.onlineUser {
		userList = append(userList, k)
	}

	return userList
}
