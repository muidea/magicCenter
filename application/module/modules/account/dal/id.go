package dal

import "muidea.com/magicCenter/application/common/dbhelper"

var userID int
var groupID int

func init() {
	dbhelper, _ := dbhelper.NewHelper()
	defer dbhelper.Release()

	userID = loadUserID(dbhelper)
	groupID = loadGroupID(dbhelper)
}

func allocUserID() int {
	userID = userID + 1
	return userID
}

func allocGroupID() int {
	groupID = groupID + 1
	return groupID
}
