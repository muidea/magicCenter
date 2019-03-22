package dbhelper

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/muidea/magicCommon/foundation/dao"
)

// DBHelper 数据访问助手
type DBHelper interface {
	BeginTransaction()

	Commit()

	Rollback()

	Query(string)

	Next() bool

	Finish()

	GetValue(...interface{})

	Execute(string) (int64, bool)

	Release()
}

type helper struct {
	dao dao.Dao
}

type databaseConfigInfo struct {
	Server   string
	Name     string
	Account  string
	Password string
}

const (
	putInHelper = iota
	fetchOutHelper
	releaseHelper
	timerCheck
)

const maxIdleSize = 10

type replyResult struct {
	dao     dao.Dao
	errInfo error
}

type helperAction struct {
	actionCode int
	dao        dao.Dao
	reply      chan replyResult
}

type helperRegistry struct {
	databaseInfo *databaseConfigInfo

	idleDaoList   []dao.Dao
	actionChannel chan *helperAction
}

var databaseInfo *databaseConfigInfo

var dbHelperRegistry *helperRegistry

// InitDB 初始化数据库
func InitDB(server, name, account, password string) {
	databaseInfo = &databaseConfigInfo{Server: server, Name: name, Account: account, Password: password}

	dbHelperRegistry = &helperRegistry{databaseInfo: databaseInfo, idleDaoList: []dao.Dao{}, actionChannel: make(chan *helperAction)}

	go dbHelperRegistry.run()
	go dbHelperRegistry.checkTimer()
}

// ParseError 解析错误信息
func ParseError(errString string) (int, error) {
	items := strings.Split(strings.ToUpper(errString), ":")
	if len(items) > 0 {
		errorReg := regexp.MustCompile("ERROR [0-9]+")
		val := errorReg.FindString(items[0])
		if len(val) > 0 {
			numReg := regexp.MustCompile("[0-9]+")
			val = numReg.FindString(val)
			errCode, _ := strconv.Atoi(val)
			return errCode, nil
		}

		msg := fmt.Sprintf("illegal errString, [%s]", errString)
		return -1, errors.New(msg)
	}

	return 0, nil
}

// NewHelper 创建数据助手
func NewHelper() (DBHelper, error) {
	if databaseInfo == nil {
		return nil, errors.New("illegal database config info")
	}

	m := &helper{dao: nil}
	dao, err := dbHelperRegistry.Fetch()
	if err == nil {
		m.dao = dao
	}

	return m, err
}

func (db *helper) BeginTransaction() {
	db.dao.BeginTransaction()
}

func (db *helper) Commit() {
	db.dao.Commit()
}

func (db *helper) Rollback() {
	db.dao.Rollback()
}

func (db *helper) Query(sql string) {
	db.dao.Query(sql)
}

func (db *helper) Next() bool {
	return db.dao.Next()
}

func (db *helper) Finish() {
	db.dao.Finish()
}

func (db *helper) GetValue(val ...interface{}) {
	db.dao.GetField(val...)
}

func (db *helper) Execute(sql string) (int64, bool) {
	return db.dao.Execute(sql)
}

func (db *helper) Release() {
	if db.dao != nil {
		dbHelperRegistry.Put(db.dao)
	}
}

func (s *helperRegistry) Fetch() (dao.Dao, error) {
	reply := make(chan replyResult)
	defer close(reply)

	action := &helperAction{actionCode: fetchOutHelper, reply: reply}

	s.actionChannel <- action

	ret := <-reply
	return ret.dao, ret.errInfo
}

func (s *helperRegistry) Put(dao dao.Dao) {
	action := &helperAction{actionCode: putInHelper, dao: dao}

	s.actionChannel <- action
}

func (s *helperRegistry) run() {
	for {
		action := <-s.actionChannel
		switch action.actionCode {
		case putInHelper:
			s.idleDaoList = append(s.idleDaoList, action.dao)
		case fetchOutHelper:
			if len(s.idleDaoList) == 0 {
				dao, err := dao.Fetch(databaseInfo.Account, databaseInfo.Password, databaseInfo.Server, databaseInfo.Name)
				if err != nil {
					log.Print("fetch database failed, err:" + err.Error())
				}
				action.reply <- replyResult{dao: dao, errInfo: err}
			} else {
				dao := s.idleDaoList[0]
				s.idleDaoList = s.idleDaoList[1:]
				action.reply <- replyResult{dao: dao, errInfo: nil}
			}
		case timerCheck:
			if len(s.idleDaoList) > maxIdleSize {
				releaseList := s.idleDaoList[maxIdleSize:]
				s.idleDaoList = s.idleDaoList[:maxIdleSize-1]
				for _, val := range releaseList {
					val.Release()
				}
			}
		}
	}
}

func (s *helperRegistry) checkTimer() {
	timeOutTimer := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-timeOutTimer.C:
			action := &helperAction{actionCode: timerCheck}
			s.actionChannel <- action
		}
	}
}
