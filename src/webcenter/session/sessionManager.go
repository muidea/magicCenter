package session

import (
	"net/http"
	"log"
	"time"
	"muidea.com/util"	
)

var SESSION_COOKIE_ID string = "session_id"

func initialize() {
	if sessionManager == nil {
		sessionManager = &SessionManager{}
		sessionManager.impl = make(sessionManagerImpl)
		go sessionManager.impl.run()
				
		go sessionManager.checkTimer()
	}
}

func uninitialize() {
	if sessionManager != nil {
		sessionManager.impl.close()
		sessionManager = nil
	}
}

func GetSession(w http.ResponseWriter, r *http.Request) *Session {
	var userSession *Session
	
	cookie, err := r.Cookie(SESSION_COOKIE_ID)	
	if err != nil {
		log.Printf("can't find cookie,create new session, err:" + err.Error())
		userSession = createSession()
	} else {
		cur, found := sessionManager.Find(cookie.Value)
		if !found {
			log.Printf("invalid cookie,create new session, cookieValue:%s", cookie.Value)
			userSession = createSession()
		} else {
			log.Print("find exist ession from cookie")
			userSession = &cur
		}
	}
	
    // 存入cookie,使用cookie存储
    session_cookie := http.Cookie{Name: SESSION_COOKIE_ID, Value: userSession.Id(),Path:"/"}
    http.SetCookie(w, &session_cookie)
	
	return userSession
}

func updateSession(session *Session) bool {
	return sessionManager.Update(session)
}

func createUUID() string {
	result := util.RandomAlphanumeric(64)
	return result
}

func createSession() *Session {
	session := Session{}
	session.id = createUUID()
	session.context = make(map[string]interface{})
	
	session.refresh()
	
	sessionManager.Insert(session)
	
	return &session
}

type SessionManager struct {
	impl sessionManagerImpl
}

var sessionManager *SessionManager = nil

func (this *SessionManager)checkTimer() {
	timeOutTimer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timeOutTimer.C:
			this.impl.checkTimeOut()
		}
	}
}

func (this *SessionManager) Insert(session Session) {
	this.impl.insert(session)
}

func (this *SessionManager) Delete(id string) {
	this.impl.remove(id)
}

func (this *SessionManager) Find(id string) (Session, bool) {
	return this.impl.find(id)
}

func (this *SessionManager) Count() int {
	return this.impl.count()
}

func (this *SessionManager) Update(session *Session) bool {
	return this.impl.update(*session)
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
	checkTimeOut
	length
	end
)

type findResult struct {
	value interface{}
	found bool
}

type sessionManagerImpl chan commandData

func (right sessionManagerImpl) insert(session Session) {
	log.Printf("insert session, id:%s", session.Id())
	
	right <- commandData{action: insert, value: session}
}

func (right sessionManagerImpl) remove(id string) {
	log.Printf("delete session, id:%s", id)
	right <- commandData{action: remove, value: id}
}

func (right sessionManagerImpl) update(session Session) bool {
	log.Printf("update session, id:%s", session.Id())
	
	reply := make(chan interface{})
	right <- commandData{action: update, value: session, result: reply}

	result := (<-reply).(bool)
	return result
}

func (right sessionManagerImpl) find(id string) (Session, bool) {
	log.Printf("find session by id,id:%s", id)
	reply := make(chan interface{})
	right <- commandData{action: find, value: id, result: reply}

	result := (<-reply).(findResult)

	if result.found {
		return result.value.(Session), result.found
	} else {
		return Session{}, false
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
			session := command.value.(Session)
			sessionInfo[session.id] = session
		case remove:
			id := command.value.(string)
			delete(sessionInfo, id)
		case update:
			session := command.value.(Session)
			_, found := sessionInfo[session.id]
			if found {
				sessionInfo[session.id] = session
			}
			command.result <- found
		case find:
			id := command.value.(string)
			
			var session Session 
			value, found := sessionInfo[id]
			if found {
				session = value.(Session)
				session.refresh()
			}
			command.result <- findResult{session, found}
		case checkTimeOut:
			removeList := []string{}
			for k, v := range sessionInfo {
				session := v.(Session)
				if session.timeOut() {
					removeList = append(removeList, k)
					delete(sessionInfo,k)
				}
			}
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

func (right sessionManagerImpl) checkTimeOut() {
	right <- commandData{action: checkTimeOut}
}



