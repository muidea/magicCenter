package resource

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/initializer"
)

type idHolder struct {
	resourceOID int
}

func (s *idHolder) Handle() {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	s.resourceOID = loadResourceOID(dbhelper)
}

func (s *idHolder) AllocResourceOID() int {
	s.resourceOID++
	return s.resourceOID
}

var holder = &idHolder{}

func init() {
	initializer.RegisterHandler(holder)
}

func allocResourceOID() int {
	return holder.AllocResourceOID()
}
