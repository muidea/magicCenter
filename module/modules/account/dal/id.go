package dal

import (
	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/common/initializer"
)

type idHolder struct {
	userID  int
	groupID int
}

func (s *idHolder) Handle() {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
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
