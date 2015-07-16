package model

import (
	"webcenter/model/datamanager"
)

type Routeline struct {
	datamanager.Routeline
	Creater string
}

func Initialize() {
	datamanager.InitDataManager()	
	
}

func Uninitialized() {
	datamanager.UninitDataManager()	
}

// 获取全部路径
func GetRouteLine() []Routeline {
	routeLine := []Routeline{}
	
	userManager := datamanager.GetUserManager()
	routelineManager := datamanager.GetRoutelineManager()
	allrouteline := routelineManager.GetAll()
	for _, ii := range allrouteline {
		route := Routeline{}
		route.Id = ii.Id
		route.Name = ii.Name
		route.Description = ii.Description
		route.CreateDate = ii.CreateDate
		user, found := userManager.FindUserById(ii.Creater)
		if found {
			route.Creater = user.Name
		} else {
			route.Creater = "admin"
		}
		
		routeLine = append(routeLine,route)
	}
		
	return routeLine
}