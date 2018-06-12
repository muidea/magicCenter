package dal

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/initializer"
)

type idHolder struct {
	userID  int
	groupID int
}

func (s *idHolder) Handle() {
	dbhelper, _ := dbhelper.NewHelper()
	defer dbhelper.Release()

	s.userID = loadUserID(dbhelper)
	s.groupID = loadGroupID(dbhelper)
}

func (s *idHolder) AllocUserID() int {
	s.userID++
	return s.userID
}

func (s *idHolder) AllocGroupID() int {
	s.groupID++
	return s.groupID
}

var holder = &idHolder{}

func init() {
	initializer.RegisterHandler(holder)
}

func allocUserID() int {
	return holder.AllocUserID()
}

func allocGroupID() int {
	return holder.AllocGroupID()
}
