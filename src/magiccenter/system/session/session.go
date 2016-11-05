package session

import (
	"magiccenter/common/model"
	"time"
)

const (
	maxTimeOut = 10
)

type impl struct {
	id      string // session id
	context map[string]interface{}
}

func (s *impl) ID() string {
	return s.id
}

func (s *impl) SetOption(key string, value interface{}) {
	s.context[key] = value

	s.save()
}

func (s *impl) GetOption(key string) (interface{}, bool) {
	value, found := s.context[key]

	return value, found
}

func (s *impl) RemoveOption(key string) {
	delete(s.context, key)

	s.save()
}

func (s *impl) GetAccount() (model.UserDetail, bool) {
	account := model.UserDetail{}
	user, found := s.context["$$userAccount"]
	if found {
		account = user.(model.UserDetail)
	}

	return account, found
}

func (s *impl) SetAccount(user model.UserDetail) {
	s.context["$$userAccount"] = user

	s.save()
}

func (s *impl) ClearAccount() {
	delete(s.context, "$$userAccount")

	s.save()
}

func (s *impl) OptionKey() []string {
	keys := []string{}
	for key := range s.context {
		keys = append(keys, key)
	}

	return keys
}

func (s *impl) refresh() {
	s.context["$$refreshTime"] = time.Now()
}

func (s *impl) timeOut() bool {
	preTime, found := s.context["$$refreshTime"]
	if !found {
		return true
	}

	nowTime := time.Now()
	elapse := nowTime.Sub(preTime.(time.Time)).Minutes()

	return elapse > maxTimeOut
}

func (s *impl) save() {
	updateSession(s)
}
