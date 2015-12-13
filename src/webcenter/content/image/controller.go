package image


import (
	"log"
	"webcenter/common"
	"webcenter/modelhelper"
)

type QueryAllImageParam struct {
	accessCode string	
}

type QueryAllImageResult struct {
	common.Result
	Image []ImageInfo
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
	url string
	desc string
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
	Image ImageInfo
}



type imageController struct {
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
	
	result.Image = QueryAllImage(model)
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
	image.SetUrl(param.url)
	image.SetDesc(param.desc)
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

	imageInfo, found := QueryImageById(model, param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Image = imageInfo
	}
	
	return result
}



