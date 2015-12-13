package image


import (
	"fmt"
	"path"
	"webcenter/modelhelper"
	"webcenter/common"
	"webcenter/content/base"
)


type ImageInfo struct {
	Id int
	Url string
	Desc string
	Creater string
	Catalog []string
}


type Image interface {
	common.Resource
	Url() string
	Desc() string
	Creater() int
	SetId(id int)
	SetUrl(url string)
	SetDesc(desc string)
	SetCreater(user int)
	SetCatalog(catalog []int)
}

type image struct {
	id int
	url string
	desc string
	creater int
	catalog []int
}

func (this *image) Id() int {
	return this.id
}

func (this *image) Name() string {
	_, name := path.Split(this.url)
	
	return name
}

func (this *image) Type() int {
	return base.IMAGE
}

func (this *image)Relative() []common.Resource {
	ress := []common.Resource{}
	
	for _, catalog := range this.catalog {
		res := common.NewSimpleRes(catalog,"", base.CATALOG)
		ress = append(ress, res)
	}
	
	return ress
}

func (this *image)Url() string {
	return this.url
}

func (this *image)Desc() string {
	return this.desc
}

func (this *image)Creater() int {
	return this.creater
}

func (this *image)SetId(id int) {
	this.id = id
}

func (this *image)SetUrl(url string) {
	this.url = url
}

func (this *image)SetDesc(desc string) {
	this.desc = desc
}

func (this *image)SetCreater(user int) {
	this.creater = user
}

func (this *image)SetCatalog(catalog []int) {
	this.catalog = catalog
}

func NewImage() Image {
	image := &image{}

	return image
}

func QueryAllImage(model modelhelper.Model) []ImageInfo {
	imageInfoList := []ImageInfo{}
	sql := fmt.Sprintf(`select i.id, i.url, i.description, u.nickname from image i, user u where i.creater = u.id`)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		image := ImageInfo{}
		model.GetValue(&image.Id, &image.Url, &image.Desc, &image.Creater)
		
		imageInfoList = append(imageInfoList, image)
	}
		
	for index, info := range imageInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.IMAGE)
		name := "-"
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					imageInfoList[index].Catalog = append(imageInfoList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}				
	}
	
	return imageInfoList
}

func QueryImageByCatalog(model modelhelper.Model, id int) []ImageInfo {
	imageInfoList := []ImageInfo{}
	sql := fmt.Sprintf(`select i.id, i.url, i.description, u.nickname from image i, user u where i.creater = u.id and i.id in (
		select src from resource_relative where dst = %d and dstType = %d and srcType = %d )`, id, base.CATALOG, base.IMAGE)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		image := ImageInfo{}
		model.GetValue(&image.Id, &image.Url, &image.Desc, &image.Creater)
		
		imageInfoList = append(imageInfoList, image)
	}
	
	for index, info := range imageInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.IMAGE)
		name := "-"
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					imageInfoList[index].Catalog = append(imageInfoList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}				
	}
	
	return imageInfoList
}

func QueryImageByRang(model modelhelper.Model, begin int,offset int) []ImageInfo {
	imageInfoList := []ImageInfo{}
	sql := fmt.Sprintf(`select i.id, i.url, i.description, u.nickname from image i, user u where i.creater = u.id and i.id in (
		select src from resource_relative where dstType = %d and srcType = %d ) and i.id >= %d limit %d`, base.CATALOG, base.IMAGE, begin, offset)
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		image := ImageInfo{}
		if model.GetValue(&image.Id, &image.Url, &image.Desc, &image.Creater) {
			imageInfoList = append(imageInfoList, image)			
		}		
	}
	
	for index, info := range imageInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.IMAGE)
		name := "-"
		if model.Query(sql) {
			for model.Next() {
				if model.GetValue(&name) {
					imageInfoList[index].Catalog = append(imageInfoList[index].Catalog, name)
				}
			}
		} else {
			panic("query failed")
		}				
	}
	
	return imageInfoList
}

func QueryImageById(model modelhelper.Model, id int) (ImageInfo, bool) {
	image := ImageInfo{}
	
	sql := fmt.Sprintf(`select i.id, i.url, i.description, u.nickname from image i, user u where i.creater = u.id and i.id = %d`, id)
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false
	for model.Next() {
		result = model.GetValue(&image.Id, &image.Url, &image.Desc, &image.Creater)
		break
	}
	if !result {
		return image, result
	}

	sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type == rr.dstType and rr.src = %d and rr.srcType=%d`, image.Id, base.IMAGE)
	name := "-"
	if model.Query(sql) {
		for model.Next() {
			if model.GetValue(&name) {
				image.Catalog = append(image.Catalog, name)
			}
		}
	} else {
		panic("query failed")
	}
	
	return image, result	
}

func DeleteImage(model modelhelper.Model, id int) bool {
	if !model.BeginTransaction() {
		return false
	}
	
	sql := fmt.Sprintf(`delete from image where id =%d`, id)	
	result := model.Execute(sql)
	if result {
		img := image{}
		img.id = id
		result  = common.DeleteResource(model, &img)
	}
		
	if !result {
		model.Rollback()
	} else {
		model.Commit()
	}
		
	return result	
}


func SaveImage(model modelhelper.Model, image Image) bool {
	sql := fmt.Sprintf(`select id from image where id=%d`, image.Id())
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
	}

	if !model.BeginTransaction() {
		return false
	}
	
	if !result {
		// insert
		sql = fmt.Sprintf(`insert into image (url,description,creater) values ('%s','%s',%d)`, image.Url(), image.Desc(), image.Creater())
		result = model.Execute(sql)
		sql = fmt.Sprintf(`select id from image where url='%s' and description ='%s' and creater=%d`, image.Url(), image.Desc(), image.Creater())
		
		id := -1
		result = model.Query(sql)
		if result {
			result = false
			for model.Next() {
				result = model.GetValue(&id)
			}
			
			image.SetId(id)
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update image set url ='%s', description='%s', creater=%d where id=%d`, image.Url(), image.Desc(), image.Creater(), image.Id())
		result = model.Execute(sql)
	}
	
	if result {
		result = common.SaveResource(model, image)
	}
	
	if result {
		model.BeginTransaction()
	} else {
		model.Rollback()
	}
	
	return result	
}


