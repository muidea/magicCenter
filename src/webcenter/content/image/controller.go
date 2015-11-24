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
	Image []Image
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
	
	result.Image = GetAllImage(model)
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

	image := newImage()
	image.Id = param.id
	image.Url = param.url
	image.Desc = param.desc
	image.Creater = param.creater
	
	if !SaveImage(model, image) {
		result.ErrCode = 1
		result.Reason = "保存链接失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存链接成功"
	}
	
	return result
}