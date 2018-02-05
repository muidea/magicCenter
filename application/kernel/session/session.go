package session

import (
	"time"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
)

const (
	maxTimeOut = 10
)

type sessionImpl struct {
	id       string // session id
	context  map[string]interface{}
	registry common.SessionRegistry
}

func (s *sessionImpl) ID() string {
	return s.id
}

func (s *sessionImpl) SetOption(key string, value interface{}) {
	s.context[key] = value

	s.save()
}

func (s *sessionImpl) GetOption(key string) (interface{}, bool) {
	value, found := s.context[key]

	return value, found
}

func (s *sessionImpl) RemoveOption(key string) {
	delete(s.context, key)

	s.save()
}

func (s *sessionImpl) GetAccount() (model.User, bool) {
	account := model.User{}
	user, found := s.context["$$userAccount"]
	if found {
		account = user.(model.User)
	}

	return account, found
}

func (s *sessionImpl) SetAccount(user model.User) {
	s.context["$$userAccount"] = user

	s.save()
}

func (s *sessionImpl) ClearAccount() {
	delete(s.context, "$$userAccount")

	s.save()
}

func (s *sessionImpl) OptionKey() []string {
	keys := []string{}
	for key := range s.context {
		keys = append(keys, key)
	}

	return keys
}

func (s *sessionImpl) refresh() {
	s.context["$$refreshTime"] = time.Now()
}

func (s *sessionImpl) timeOut() bool {
	preTime, found := s.context["$$refreshTime"]
	if !found {
		return true
	}

	nowTime := time.Now()
	elapse := nowTime.Sub(preTime.(time.Time)).Minutes()

	return elapse > maxTimeOut
}

func (s *sessionImpl) save() {
	s.registry.UpdateSession(s)
}
