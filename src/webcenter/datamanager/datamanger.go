package datamanager

import (

)

func init() {
	initDataManager()
}

var userManager *UserManager
var groupManager *GroupManager
var routelineManager *RoutelineManager
var checkpointManager *CheckpointManager

func initDataManager() {
	userManager = &UserManager{}
	groupManager = &GroupManager{}
	routelineManager = &RoutelineManager{}
	checkpointManager = &CheckpointManager{}
	
	userManager.Load()
	groupManager.Load()
	routelineManager.Load()	
	checkpointManager.Load()
}

func uninitDataManager() {
	userManager.Unload()
	groupManager.Unload()
	routelineManager.Unload()
	checkpointManager.Unload()
	
	userManager = nil
	groupManager = nil
	routelineManager = nil
	checkpointManager = nil
}

func GetUserManager() (*UserManager) {
	return userManager
}

func GetGroupManager() (*GroupManager) {
	return groupManager
}

func GetRoutelineManager() (*RoutelineManager) {
	return routelineManager
}


func GetCheckpointManager() (*CheckpointManager) {
	return checkpointManager
}
