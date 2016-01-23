package link


import (
	"log"
	"webcenter/util/modelhelper"
	"webcenter/kernel/admin/common"
)


type QueryManageInfo struct {
	LinkInfo []LinkInfo
}

type QueryAllLinkParam struct {
	accessCode string	
}

type QueryAllLinkResult struct {
	common.Result
	Link []LinkInfo
}

type QueryLinkParam struct {
	accessCode string
	id int
}

type QueryLinkResult struct {
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

type SubmitLinkParam struct {
	accessCode string
	id int
	name string
	url string
	logo string
	style int
	catalog []int
	creater int
}

type SubmitLinkResult struct {
	common.Result
}

type EditLinkParam struct {
	accessCode string
	id int
}

type EditLinkResult struct {
	common.Result
	Link Link
}

type linkController struct {
}

func (this *linkController)queryManageInfoAction() QueryManageInfo {
	info := QueryManageInfo{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	info.LinkInfo = QueryAllLink(model)

	return info
}

func (this *linkController)queryAllLinkAction(param QueryAllLinkParam) QueryAllLinkResult {
	result := QueryAllLinkResult{}
	
	model, err := modelhelper.NewHelper()
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
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	link, found := QueryLinkById(model,param.id)
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
	
	model, err := modelhelper.NewHelper()
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
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	link := NewLink()
	link.SetId(param.id)
	link.SetName(param.name)
	link.SetUrl(param.url)
	link.SetLogo(param.logo)
	link.SetStyle(param.style)
	link.SetCatalog(param.catalog)
	link.SetCreater(param.creater)
	
	if !SaveLink(model, link) {
		result.ErrCode = 1
		result.Reason = "保存链接失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存链接成功"
	}
	
	return result
}

func (this *linkController)editLinkAction(param EditLinkParam) EditLinkResult {
	result := EditLinkResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	link, found := QueryLinkById(model,param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Link = link
	}

	return result
}



