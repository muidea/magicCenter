package group

import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type QueryManageInfo struct {
	GroupInfo []GroupInfo
}

type QueryAllGroupParam struct {
	accessCode string	
}

type QueryAllGroupResult struct {
	common.Result
	Group []GroupInfo
}

type QueryGroupParam struct {
	accessCode string
	id int
}

type QueryGroupResult struct {
	common.Result
	Group Group
}

type DeleteGroupParam struct {
	accessCode string
	id int
}

type DeleteGroupResult struct {
	common.Result
}

type SubmitGroupParam struct {
	accessCode string
	id int
	name string
}

type SubmitGroupResult struct {
	common.Result
}

type accountController struct {
}

func (this *accountController)queryManageInfoAction() QueryManageInfo {
	info := QueryManageInfo{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	info.GroupInfo = QueryAllGroup(model)

	return info
}
 
func (this *accountController)queryAllGroupAction(param QueryAllGroupParam) QueryAllGroupResult {
	result := QueryAllGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
		
	result.Group = QueryAllGroup(model)
	result.ErrCode = 0
	
	return result
}

func (this *accountController)queryGroupAction(param QueryGroupParam) QueryGroupResult {
	result := QueryGroupResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
		
	catalog, found := QueryGroupById(model, param.id)
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

	SaveGroup(model, group)
		
	result.ErrCode = 0
	result.Reason = "保存分组成功"
	
	return result
}



