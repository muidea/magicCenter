package group

import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)


type GetGroupParam struct {
	accessCode string
	id int
}

type GetGroupResult struct {
	common.Result
	Group Group
}

type GetAllSubGroupParam struct {
	accessCode string
	id int
}

type GetAllSubGroupResult struct {
	common.Result
	Group []Group
}

type DeleteGroupParam struct {
	accessCode string
	id int
}

type DeleteGroupResult struct {
	common.Result
}

type GetAllGroupParam struct {
	accessCode string	
}

type GetAllGroupResult struct {
	common.Result
	Group []Group
}


type SubmitGroupParam struct {
	accessCode string
	id int
	name string
	parent int
	submitDate string	
}

type SubmitGroupResult struct {
	common.Result
}

type accountController struct {
}

 
func (this *accountController)getAllGroupAction(param GetAllGroupParam) GetAllGroupResult {
	result := GetAllGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	result.Group = GetAllGroup(model)
	result.ErrCode = 0
	
	return result
}

func (this *accountController)getAllSubGroupAction(param GetAllSubGroupParam) GetAllSubGroupResult {
	result := GetAllSubGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	subGroups := GetAllSubGroup(model, param.id)
	result.ErrCode = 0
	result.Group = subGroups

	return result
}

func (this *accountController)getGroupAction(param GetGroupParam) GetGroupResult {
	result := GetGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	catalog, found := GetGroupById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Group = catalog
	}

	return result
}

func (this *accountController)deleteGroupAction(param DeleteGroupParam) DeleteGroupResult {
	result := DeleteGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	DeleteGroup(model, param.id)
	result.ErrCode = 0
	result.Reason = "删除分组成功"
	
	return result
}


func (this *accountController)submitGroupAction(param SubmitGroupParam) SubmitGroupResult {
	result := SubmitGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	group := newGroup()
	group.Id = param.id
	group.Name = param.name
	group.Catalog = param.parent

	SaveGroup(model, group)
		
	result.ErrCode = 0
	result.Reason = "保存分组成功"
	
	return result
}



