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
	log.Print("sessionRegistry init....")
	sessionCookieID = createUUID()
}

func createUUID() string {
	return util.RandomAlphanumeric(64)
}

type sessionRegistryImpl struct {
	commandChan commandChanImpl
}

// CreateSessionRegistry 创建Session仓库
func CreateSessionRegistry() common.SessionRegistry {
	impl := sessionRegistryImpl{}
	impl.commandChan = make(commandChanImpl)
	go impl.commandChan.run()
	go impl.checkTimer()

	return &impl
}

// GetSession 获取Session对象
func (sm *sessionRegistryImpl) GetSession(w http.ResponseWriter, r *http.Request) common.Session {
	var userSession common.Session

	cookie, err := r.Cookie(sessionCookieID)
	if err != nil {
		log.Printf("can't find cookie,create new session, err:" + err.Error())
		id := createUUID()
		userSession = sm.CreateSession(id)
	} else {
		cur, found := sm.FindSession(cookie.Value)
		if !found {
			log.Printf("invalid cookie,create new session, cookieValue:%s", cookie.Value)
			id := createUUID()
			userSession = sm.CreateSession(id)
		} else {
			log.Print("find exist ession from cookie")
			userSession = cur
		}
	}

	// 存入cookie,使用cookie存储
	sessionCookie := http.Cookie{Name: sessionCookieID, Value: userSession.ID(), Path: "magicCenter"}
	http.SetCookie(w, &sessionCookie)

	return userSession
}

// CreateSession 新建Session
func (sm *sessionRegistryImpl) CreateSession(sessionID string) common.Session {
	session := sessionImpl{id: sessionID, context: make(map[string]interface{}), registry: sm}

	session.refresh()

	sm.commandChan.insert(session)

	return &session
}

func (sm *sessionRegistryImpl) FindSession(sessionID string) (common.Session, bool) {
	session, found := sm.commandChan.find(sessionID)
	return &session, found
}

// UpdateSession 更新Session
func (sm *sessionRegistryImpl) UpdateSession(session common.Session) bool {
	cur, found := sm.commandChan.find(session.ID())
	if !found {
		return false
	}

	for _, key := range session.OptionKey() {
		cur.context[key], _ = session.GetOption(key)
	}

	return sm.commandChan.update(cur)
}

func (sm *sessionRegistryImpl) checkTimer() {
	timeOutTimer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timeOutTimer.C:
			sm.commandChan.checkTimeOut()
		}
	}
}

func (sm *sessionRegistryImpl) insert(session sessionImpl) {
	sm.commandChan.insert(session)
}

func (sm *sessionRegistryImpl) delete(id string) {
	sm.commandChan.remove(id)
}

func (sm *sessionRegistryImpl) find(id string) (sessionImpl, bool) {
	return sm.commandChan.find(id)
}

func (sm *sessionRegistryImpl) count() int {
	return sm.commandChan.count()
}

func (sm *sessionRegistryImpl) update(session sessionImpl) bool {
	return sm.commandChan.update(session)
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

type commandChanImpl chan commandData

func (right commandChanImpl) insert(session sessionImpl) {
	log.Printf("insert session, id:%s", session.id)

	right <- commandData{action: insert, value: session}
}

func (right commandChanImpl) remove(id string) {
	log.Printf("delete session, id:%s", id)
	right <- commandData{action: remove, value: id}
}

func (right commandChanImpl) update(session sessionImpl) bool {
	log.Printf("update session, id:%s", session.id)

	reply := make(chan interface{})
	right <- commandData{action: update, value: session, result: reply}

	result := (<-reply).(bool)
	return result
}

func (right commandChanImpl) find(id string) (sessionImpl, bool) {
	log.Printf("find session by id,id:%s", id)
	reply := make(chan interface{})
	right <- commandData{action: find, value: id, result: reply}

	result := (<-reply).(findResult)

	if result.found {
		return result.value.(sessionImpl), result.found
	}

	return sessionImpl{}, false
}

func (right commandChanImpl) count() int {
	log.Print("count session")

	reply := make(chan interface{})
	right <- commandData{action: length, result: reply}

	result := (<-reply).(int)
	return result
}

func (right commandChanImpl) run() {
	sessionInfo := make(map[string]interface{})
	for command := range right {
		switch command.action {
		case insert:
			session := command.value.(sessionImpl)
			sessionInfo[session.id] = &session
		case remove:
			id := command.value.(string)
			delete(sessionInfo, id)
		case update:
			session := command.value.(sessionImpl)
			_, found := sessionInfo[session.id]
			if found {
				sessionInfo[session.id] = &session
			}

			command.result <- found
		case find:
			id := command.value.(string)
			session := sessionImpl{}
			cur, found := sessionInfo[id]
			if found {
				cur.(*sessionImpl).refresh()
				session = *(cur.(*sessionImpl))
			}
			command.result <- findResult{session, found}
		case checkTimeOut:
			removeList := []string{}
			for k, v := range sessionInfo {
				session := v.(*sessionImpl)
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

	log.Print("session manager sessionImpl exit")
}

func (right commandChanImpl) close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	right <- commandData{action: end, data: reply}
	return <-reply
}

func (right commandChanImpl) checkTimeOut() {
	right <- commandData{action: checkTimeOut}
}
