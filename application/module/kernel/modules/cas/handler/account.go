package handler

import (
	"log"
	"time"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCommon/model"
)

type accountManager struct {
	accountHandler common.AccountHandler
	onlineUser     map[int]model.AccountOnlineView
}

func createAccountManager(moduleHub common.ModuleHub) (accountManager, bool) {
	s := accountManager{accountHandler: nil, onlineUser: make(map[int]model.AccountOnlineView)}
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
func (s *accountManager) userLogin(account, password, remoteAddr, authToken string) (model.AccountOnlineView, bool) {
	info := model.AccountOnlineView{}

	user, ok := s.accountHandler.FindUserByAccount(account, password)
	if !ok {
		return info, ok
	}

	info, ok = s.onlineUser[user.ID]
	if ok {
		if info.Address != remoteAddr {
			log.Printf("duplicate user[%d] logining,pre address:%s, cur address:%s", info.ID, info.Address, remoteAddr)
		}

		return info, ok
	}

	info.ID = user.ID
	info.Name = user.Name
	info.Address = remoteAddr
	info.LoginTime = time.Now().Unix()
	info.UpdateTime = info.LoginTime
	info.AuthToken = authToken
	s.onlineUser[user.ID] = info
	return info, true
}

func (s *accountManager) userRefresh(id int, remoteAddr string) {
	info, ok := s.onlineUser[id]
	if ok {
		if info.Address != remoteAddr {
			log.Printf("illegal user[%d] refresh, pre address:%s, cur address:%s", id, info.Address, remoteAddr)
		}
		info.UpdateTime = time.Now().Unix()
		s.onlineUser[id] = info
	} else {
		log.Printf("illegal user[%d] refresh, not login, address:%s", id, remoteAddr)
	}
}

func (s accountManager) userVerify(id int) (model.AccountOnlineView, bool) {
	info, ok := s.onlineUser[id]
	return info, ok
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
