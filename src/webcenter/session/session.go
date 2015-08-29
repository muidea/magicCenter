package session

import (
)

func init() {
	initialize()
}

type Session struct {
	id      string // session id
	context map[string]interface{}
}

func (this *Session) Id() string {
	return this.id
}

func (this *Session) Account() (string, bool) {
	account, found := this.context["account"]
	if found {
		return account.(string), found
	}

	return "", found
}

func (this *Session) AccessToken() string {
	token := createUUID()

	this.context["access_token"] = token
	
	return token
}

func (this *Session) ValidToken(token string) bool {
	cur, found := this.context["access_token"]
	if !found {
		return false
	}

	return cur.(string) == token
}

func (this *Session) ReleaseAccessToken() {
	delete(this.context, "access_token")
}

func (this *Session) SetOption(key string, value interface{}) {
	this.context[key] = value
}

func (this *Session) GetOption(key string) (interface{}, bool) {
	value, found := this.context[key]
	
	return value, found	
}

func (this *Session) RemoveOption(key string) {
	delete(this.context, key)
}

func (this *Session) Save() {
	updateSession(this)
}



