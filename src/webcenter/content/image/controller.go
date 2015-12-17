package image


import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type QueryManageInfo struct {
	ImageInfo []ImageInfo
}

type QueryAllImageParam struct {
	accessCode string	
}

type QueryAllImageResult struct {
	common.Result
	ImageInfo []ImageInfo
}

type DeleteImageParam struct {
	accessCode string
	id int
}

type DeleteImageResult struct {
	common.Result
}

type SubmitImageParam struct {
	accessCode string
	id int
	name string
	url string
	desc string
	catalog []int
	creater int
}

type SubmitImageResult struct {
	common.Result
}

type EditImageParam struct {
	accessCode string
	id int
}

type EditImageResult struct {
	common.Result
	Image Image
}

type imageController struct {
}


func (this *imageController)queryManageInfoAction() QueryManageInfo {
	info := QueryManageInfo{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
			
	info.ImageInfo = QueryAllImage(model)

	return info
}

func (this *imageController)queryAllImageAction(param QueryAllImageParam) QueryAllImageResult {
	result := QueryAllImageResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
	
	result.ImageInfo = QueryAllImage(model)
	result.ErrCode = 0
	
	return result
}

func (this *imageController)deleteImageAction(param DeleteImageParam) DeleteImageResult {
	result := DeleteImageResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()
		
	DeleteImage(model, param.id)
	result.ErrCode = 0
	
	return result
}


func (this *imageController)submitImageAction(param SubmitImageParam) SubmitImageResult {
	result := SubmitImageResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	image := NewImage()
	image.SetId(param.id)
	image.SetName(param.name)
	image.SetUrl(param.url)
	image.SetDesc(param.desc)
	image.SetCatalog(param.catalog)
	image.SetCreater(param.creater)
	
	if !SaveImage(model, image) {
		result.ErrCode = 1
		result.Reason = "保存图片失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存图片成功"
	}
	
	return result
}


func (this *imageController)editImageAction(param EditImageParam) EditImageResult {
	result := EditImageResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	image, found := QueryImageById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Image = image
	}
	
	return result
}



