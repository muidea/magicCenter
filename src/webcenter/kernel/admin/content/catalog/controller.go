package catalog


import (
	"webcenter/util/modelhelper"
	"webcenter/kernel/admin/common"
	"webcenter/kernel/admin/content/base"
)

type QueryManageInfo struct {
	Catalog []CatalogInfo
}

type QueryAllCatalogInfoParam struct {
	accessCode string	
}

type QueryAllCatalogInfoResult struct {
	common.Result
	Catalog []CatalogInfo
}

type QueryCatalogInfoParam struct {
	accessCode string
	id int
}

type QueryCatalogInfoResult struct {
	common.Result
	Catalog CatalogInfo
}

type QueryAvalibleParentCatalogInfoParam struct {
	accessCode string
	id int
}

type QueryAvalibleParentCatalogInfoResult struct {
	common.Result
	Catalog []CatalogInfo
}


type QuerySubCatalogInfoParam struct {
	accessCode string
	id int
}

type QuerySubCatalogInfoResult struct {
	common.Result
	Catalog []CatalogInfo
}

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

type SubmitCatalogParam struct {
	id int
	name string
	pid []int
	creater int
}

type SubmitCatalogResult struct {
	common.Result
}

type EditCatalogParam struct {
	accessCode string
	id int	
}

type EditCatalogResult struct {
	common.Result
	Catalog CatalogInfo
	AvalibleParent []CatalogInfo
}

type catalogController struct {
}

func (this *catalogController)queryManageInfoAction() QueryManageInfo {
	info := QueryManageInfo{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	info.Catalog = QueryAllCatalogInfo(model)

	return info
}


func (this *catalogController)queryAllCatalogInfoAction(param QueryAllCatalogInfoParam) QueryAllCatalogInfoResult {
	result := QueryAllCatalogInfoResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	result.Catalog = QueryAllCatalogInfo(model)
	result.ErrCode = 0

	return result
}

func (this *catalogController)queryCatalogInfoAction(param QueryCatalogInfoParam) QueryCatalogInfoResult {
	result := QueryCatalogInfoResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	catalog, found := QueryCatalogInfoById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Catalog = catalog
	}
	
	return result
}

func (this *catalogController)queryAvalibleParentCatalogInfoAction(param QueryAvalibleParentCatalogInfoParam) QueryAvalibleParentCatalogInfoResult {
	result := QueryAvalibleParentCatalogInfoResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	result.Catalog = QueryAvalibleParentCatalogInfo(model, param.id)
	
	return result
}

func (this *catalogController)querySubCatalogInfoAction(param QuerySubCatalogInfoParam) QuerySubCatalogInfoResult {
	result := QuerySubCatalogInfoResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	result.Catalog = QuerySubCatalogInfo(model, param.id)
	
	return result
}

func (this *catalogController)queryCatalogAction(param QueryCatalogParam) QueryCatalogResult {
	result := QueryCatalogResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	catalog, found := QueryCatalogById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Catalog = catalog
	}

	return result
}

func (this *catalogController)deleteCatalogAction(param DeleteCatalogParam) DeleteCatalogResult {
	result := DeleteCatalogResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	subCatalogList := common.QueryReferenceResource(model,param.id, base.CATALOG, false)
	if len(subCatalogList) >0 {
		result.ErrCode = 100;
		result.Reason = "该分类被引用，无法直接删除";
		return result
	}

	DeleteCatalogById(model, param.id)
	result.ErrCode = 0
	
	return result
}
 

func (this *catalogController)submitCatalogAction(param SubmitCatalogParam) SubmitCatalogResult {
	result := SubmitCatalogResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	catalog := NewCatalog()
	catalog.SetId(param.id)
	catalog.SetName(param.name)
	catalog.SetParent(param.pid)
	catalog.SetCreater(param.creater)
		
	if !SaveCatalog(model, catalog) {
		result.ErrCode = 1
		result.Reason = "保存分类失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存分类成功"
	}
	
	return result
}


func (this *catalogController)editCatalogAction(param EditCatalogParam) EditCatalogResult {
	result := EditCatalogResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	catalog, found := QueryCatalogInfoById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		
		result.AvalibleParent = QueryAvalibleParentCatalogInfo(model, param.id)
		
		result.ErrCode = 0
		result.Catalog = catalog
	}
	
	return result
}


