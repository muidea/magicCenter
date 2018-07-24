package handler

import (
	"time"

	"muidea.com/magicCenter/common"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

type endpointManager struct {
	endpointHandler common.EndpointHandler
}

func createEndpointManager(moduleHub common.ModuleHub) (endpointManager, bool) {
	s := endpointManager{endpointHandler: nil}
	mod, ok := moduleHub.FindModule(common.EndpointModuleID)
	if !ok {
		return s, false
	}

	entryPoint := mod.EntryPoint()
	switch entryPoint.(type) {
	case common.EndpointHandler:
		s.endpointHandler = entryPoint.(common.EndpointHandler)
	}

	return s, s.endpointHandler != nil
}

// bool 是否登陆成功
func (s *endpointManager) endpointLogin(endpointID, authToken, remoteAddr string) (model.OnlineEntryView, bool) {
	info := model.OnlineEntryView{}

	endpoint, ok := s.endpointHandler.QueryEndpointByID(endpointID)
	if !ok {
		return info, ok
	}

	if endpoint.AuthToken != authToken {
		return info, false
	}

	info.ID = common_const.SystemAccountUser.ID
	info.Name = common_const.SystemAccountUser.Name
	info.Address = remoteAddr
	info.LoginTime = time.Now().Unix()
	info.UpdateTime = info.LoginTime
	info.IdentifyID = endpointID

	return info, true
}
