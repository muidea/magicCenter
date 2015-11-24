package link


import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)


type QueryLinkParam struct {
	accessCode string
	id int
}

type QueryLinkResult struct {
	common.Result
	Link Link
}


type EditLinkParam struct {
	accessCode string
	id int
}

type EditLinkResult struct {
	common.Result
	Link Link
}

type DeleteLinkParam struct {
	accessCode string
	id int
}

type DeleteLinkResult struct {
	common.Result
}

type QueryAllLinkParam struct {
	accessCode string	
}

type QueryAllLinkResult struct {
	common.Result
	Link []Link
}

type SubmitLinkParam struct {
	accessCode string
	id int
	name string
	url string
	logo string
	style int
	creater int
}

type SubmitLinkResult struct {
	common.Result
}

type linkController struct {
}


func (this *linkController)queryAllLinkAction(param QueryAllLinkParam) QueryAllLinkResult {
	result := QueryAllLinkResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	result.Link = QueryAllLink(model)
	result.ErrCode = 0
	
	return result
}
 
func (this *linkController)queryLinkAction(param QueryLinkParam) QueryLinkResult {
	result := QueryLinkResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	link, found := QueryLink(model,param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Link = link
	}

	return result
}

func (this *linkController)editLinkAction(param EditLinkParam) EditLinkResult {
	result := EditLinkResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	link, found := QueryLink(model,param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Link = link
	}

	return result
}

func (this *linkController)deleteLinkAction(param DeleteLinkParam) DeleteLinkResult {
	result := DeleteLinkResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	DeleteLink(model, param.id)
	
	result.ErrCode = 0
	
	return result
}


func (this *linkController)submitLinkAction(param SubmitLinkParam) SubmitLinkResult {
	result := SubmitLinkResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	link := newLink()
	link.Id = param.id
	link.Name = param.name
	link.Url = param.url
	link.Logo = param.logo
	link.Style = param.style
	link.Creater = param.creater
	
	if !SaveLink(model, link) {
		result.ErrCode = 1
		result.Reason = "保存链接失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存链接成功"
	}
	
	return result
}

