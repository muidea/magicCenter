package session

import (
	"log"
	"net/http"
	"time"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/foundation/util"
)

var sessionCookieID = "session_id"

func init() {
	sessionCookieID = createUUID()

	initialize()
}

func initialize() {
	if sessionManager == nil {
		sessionManager = &sessionManagerImpl{}
		sessionManager.impl = make(sessionChanImpl)
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

// GetSession 获取Session对象
func GetSession(w http.ResponseWriter, r *http.Request) common.Session {
	var userSession common.Session

	cookie, err := r.Cookie(sessionCookieID)
	if err != nil {
		log.Printf("can't find cookie,create new session, err:" + err.Error())
		userSession = createSession()
	} else {
		cur, found := sessionManager.find(cookie.Value)
		if !found {
			log.Printf("invalid cookie,create new session, cookieValue:%s", cookie.Value)
			userSession = createSession()
		} else {
			log.Print("find exist ession from cookie")
			userSession = &cur
		}
	}

	// 存入cookie,使用cookie存储
	sessionCookie := http.Cookie{Name: sessionCookieID, Value: userSession.ID(), Path: "/"}
	http.SetCookie(w, &sessionCookie)

	return userSession
}

func updateSession(session common.Session) bool {
	cur, found := sessionManager.find(session.ID())
	if !found {
		return false
	}

	for _, key := range session.OptionKey() {
		cur.context[key], _ = session.GetOption(key)
	}

	return sessionManager.update(cur)
}

func createUUID() string {
	result := util.RandomAlphanumeric(64)
	return result
}

func createSession() common.Session {
	session := impl{}
	session.id = createUUID()
	session.context = make(map[string]interface{})

	session.refresh()

	sessionManager.insert(session)

	return &session
}

type sessionManagerImpl struct {
	impl sessionChanImpl
}

var sessionManager *sessionManagerImpl

func (sm *sessionManagerImpl) checkTimer() {
	timeOutTimer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timeOutTimer.C:
			sm.impl.checkTimeOut()
		}
	}
}

func (sm *sessionManagerImpl) insert(session impl) {
	sm.impl.insert(session)
}

func (sm *sessionManagerImpl) delete(id string) {
	sm.impl.remove(id)
}

func (sm *sessionManagerImpl) find(id string) (impl, bool) {
	return sm.impl.find(id)
}

func (sm *sessionManagerImpl) count() int {
	return sm.impl.count()
}

func (sm *sessionManagerImpl) update(session impl) bool {
	return sm.impl.update(session)
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

type sessionChanImpl chan commandData

func (right sessionChanImpl) insert(session impl) {
	log.Printf("insert session, id:%s", session.id)

	right <- commandData{action: insert, value: session}
}

func (right sessionChanImpl) remove(id string) {
	log.Printf("delete session, id:%s", id)
	right <- commandData{action: remove, value: id}
}

func (right sessionChanImpl) update(session impl) bool {
	log.Printf("update session, id:%s", session.id)

	reply := make(chan interface{})
	right <- commandData{action: update, value: session, result: reply}

	result := (<-reply).(bool)
	return result
}

func (right sessionChanImpl) find(id string) (impl, bool) {
	log.Printf("find session by id,id:%s", id)
	reply := make(chan interface{})
	right <- commandData{action: find, value: id, result: reply}

	result := (<-reply).(findResult)

	if result.found {
		return result.value.(impl), result.found
	}

	return impl{}, false
}

func (right sessionChanImpl) count() int {
	log.Print("count session")

	reply := make(chan interface{})
	right <- commandData{action: length, result: reply}

	result := (<-reply).(int)
	return result
}

func (right sessionChanImpl) run() {
	sessionInfo := make(map[string]interface{})
	for command := range right {
		switch command.action {
		case insert:
			session := command.value.(impl)
			sessionInfo[session.id] = &session
		case remove:
			id := command.value.(string)
			delete(sessionInfo, id)
		case update:
			session := command.value.(impl)
			_, found := sessionInfo[session.id]
			if found {
				sessionInfo[session.id] = &session
			}

			command.result <- found
		case find:
			id := command.value.(string)
			session := impl{}
			cur, found := sessionInfo[id]
			if found {
				cur.(*impl).refresh()
				session = *(cur.(*impl))
			}
			command.result <- findResult{session, found}
		case checkTimeOut:
			removeList := []string{}
			for k, v := range sessionInfo {
				session := v.(*impl)
				if session.timeOut() {
					removeList = append(removeList, k)
				}
			}

			for key := range removeList {
				delete(sessionInfo, removeList[key])
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

func (right sessionChanImpl) close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	right <- commandData{action: end, data: reply}
	return <-reply
}

func (right sessionChanImpl) checkTimeOut() {
	right <- commandData{action: checkTimeOut}
}
