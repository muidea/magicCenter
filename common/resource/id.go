package resource

import "muidea.com/magicCenter/common/dbhelper"

var resourceOID int

func init() {
	dbhelper, _ := dbhelper.NewHelper()
	defer dbhelper.Release()

	resourceOID = loadResourceOID(dbhelper)
}

func allocResourceOID() int {
	resourceOID = resourceOID + 1
	return resourceOID
}
