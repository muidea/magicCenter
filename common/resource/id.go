package resource

import (
	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/common/initializer"
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
