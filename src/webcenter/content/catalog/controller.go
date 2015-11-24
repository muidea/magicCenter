package catalog


import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type QueryCatalogParam struct {
	accessCode string
	id int
}

type QueryCatalogResult struct {
	common.Result
	Catalog Catalog
}

type DeleteCatalogParam struct {
	accessCode string
	id int
}

type DeleteCatalogResult struct {
	common.Result
}

type QueryAllCatalogParam struct {
	accessCode string	
}

type QueryAllCatalogResult struct {
	common.Result
	Catalog []Catalog
}

type SubmitCatalogParam struct {
	id int
	name string
	pid int
	submitDate string
	creater int
}

type SubmitCatalogResult struct {
	common.Result
}

type catalogController struct {
}


func (this *catalogController)queryAllCatalogAction(param QueryAllCatalogParam) QueryAllCatalogResult {
	result := QueryAllCatalogResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
			
	result.Catalog = GetAllCatalog(model)
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *catalogController)queryCatalogAction(param QueryCatalogParam) QueryCatalogResult {
	result := QueryCatalogResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	article, found := QueryCatalogById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Catalog = article
	}
	
	return result
}

func (this *catalogController)deleteCatalogAction(param DeleteCatalogParam) DeleteCatalogResult {
	result := DeleteCatalogResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	DeleteCatalogById(model, param.id)
	result.ErrCode = 0
	
	return result
}
 

func (this *catalogController)submitCatalogAction(param SubmitCatalogParam) SubmitCatalogResult {
	result := SubmitCatalogResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	catalog := newCatalog()
	catalog.Id = param.id
	catalog.Name = param.name
	catalog.Pid = param.pid
	catalog.Creater = param.creater
	
	if !SaveCatalog(model, catalog) {
		result.ErrCode = 1
		result.Reason = "保存分类失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存分类成功"
	}
	
	model.Release()

	return result
}


