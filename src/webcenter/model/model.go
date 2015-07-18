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
func GetAllRouteLine() []Routeline {
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

func GetRouline(id int) (Routeline, bool) {
	userManager := datamanager.GetUserManager()
	routelineManager := datamanager.GetRoutelineManager()
	
	routeline := Routeline{}
	line, found := routelineManager.FindById(id)
	if found {
		routeline.Id = id
		routeline.Name = line.Name
		routeline.Description = line.Description
		routeline.CreateDate = line.CreateDate
		user, found := userManager.FindUserById(line.Creater)
		if found {
			routeline.Creater = user.Name
		} else {
			routeline.Creater = "admin"
		}
	}
	
	return routeline, found
}


