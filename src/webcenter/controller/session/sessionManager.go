package session

import (
	"log"
)

type Session struct {
	id      string // session id
	context map[string]interface{}
}

type SessionManager struct {
	impl sessionManagerImpl
}

var sessionManager *SessionManager = nil

func Initialize() {
	if sessionManager == nil {
		sessionManager = &SessionManager{}
		sessionManager.impl = make(sessionManagerImpl)
		go sessionManager.impl.run()
	}
}

func Uninitialize() {
	if sessionManager != nil {
		sessionManager.impl.close()
		sessionManager = nil
	}
}

func SessionManger() *SessionManager {
	return sessionManager
}

func createUUID() string {
	return RandomAlphanumeric(32)
}

func CreateSession() *Session {
	session := new(Session)
	session.id = createUUID()
	session.context = make(map[string]interface{})
	
	sessionManager.Insert(session)
	
	return session
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

func (this *SessionManager) Insert(session *Session) {
	this.impl.insert(session)
}

func (this *SessionManager) Delete(id string) {
	this.impl.delete(id)
}

func (this *SessionManager) Find(id string) (*Session, bool) {
	return this.impl.find(id)
}

func (this *SessionManager) Count() int {
	return this.impl.count()
}

type commandData struct {
	action commandAction
	value  interface{}
	result chan<- interface{}
	data   chan<- map[string]interface{}
}

type commandAction int

const (
	insert commandAction = iota
	remove
	update
	find
	length
	end
)

type findResult struct {
	value interface{}
	found bool
}

type sessionManagerImpl chan commandData

func (right sessionManagerImpl) insert(session *Session) {
	log.Printf("insert session, id:%s", session.Id())
	
	right <- commandData{action: insert, value: session}
}

func (right sessionManagerImpl) delete(id string) {
	log.Printf("delete session, id:%s", id)
	right <- commandData{action: remove, value: id}
}

func (right sessionManagerImpl) update(session *Session) bool {
	log.Printf("update session, id:%s", session.Id())
	
	reply := make(chan interface{})
	right <- commandData{action: update, value: session, result: reply}

	result := (<-reply).(bool)
	return result
}

func (right sessionManagerImpl) find(id string) (*Session, bool) {
	log.Printf("find session by id,id:%s", id)
	reply := make(chan interface{})
	right <- commandData{action: find, value: id, result: reply}

	result := (<-reply).(findResult)

	if result.found {
		return result.value.(*Session), result.found
	} else {
		return nil, false
	}
}

func (right sessionManagerImpl) count() int {
	log.Print("count session")
	
	reply := make(chan interface{})
	right <- commandData{action: length, result: reply}

	result := (<-reply).(int)
	return result
}

func (right sessionManagerImpl) run() {
	sessionInfo := make(map[string]interface{})
	for command := range right {
		switch command.action {
		case insert:
			session := command.value.(*Session)
			sessionInfo[session.id] = session
		case remove:
			id := command.value.(string)
			delete(sessionInfo, id)
		case update:
			session := command.value.(*Session)
			_, found := sessionInfo[session.id]
			if found {
				sessionInfo[session.id] = session
			}
			command.result <- found
		case find:
			id := command.value.(string)
			session, found := sessionInfo[id]
			command.result <- findResult{session, found}
		case length:
			command.result <- len(sessionInfo)
		case end:
			close(right)
			command.data <- sessionInfo
		}
	}
	
	log.Print("session manager impl exit")
}

func (right sessionManagerImpl) close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	right <- commandData{action: end, data: reply}
	return <-reply
}
