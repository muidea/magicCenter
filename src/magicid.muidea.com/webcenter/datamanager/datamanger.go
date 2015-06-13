package datamanager

import (

)

var userManager *UserManager
var groupManager *GroupManager

func InitDataManager() {
	userManager = &UserManager{}
	groupManager = &GroupManager{}
	
	userManager.Load()
	groupManager.Load()
}

func UninitDataManager() {
	userManager.Unload()
	groupManager.Unload()
	
	userManager = nil
	groupManager = nil
}

func GetUserManager() (*UserManager) {
	return userManager
}

func GetGroupManager() (*GroupManager) {
	return groupManager
}
