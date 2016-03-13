package image


import (
	"fmt"
	"webcenter/util/modelhelper"
	"webcenter/kernel/admin/common"
	"webcenter/kernel/admin/content/base"
)

type ImageInfo struct {
	Id int
	Name string
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
	SetName(name string)
	SetUrl(url string)
	SetDesc(desc string)
	SetCreater(user int)
	SetCatalog(catalog []int)
}

type image struct {
	id int
	name string
	url string
	desc string
	creater int
	catalog []int
}

func (this *image) Id() int {
	return this.id
}

func (this *image) Name() string {
	return this.name
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

func (this *image)SetName(name string) {
	this.name = name
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
	sql := fmt.Sprintf(`select i.id, i.name, i.url, i.description, u.nickname from image i, user u where i.creater = u.id`)
	model.Query(sql)

	for model.Next() {
		image := ImageInfo{}
		model.GetValue(&image.Id, &image.Name, &image.Url, &image.Desc, &image.Creater)
		
		imageInfoList = append(imageInfoList, image)
	}
		
	for index, info := range imageInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.IMAGE)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			imageInfoList[index].Catalog = append(imageInfoList[index].Catalog, name)
		}
	}
	
	return imageInfoList
}

func QueryImageByCatalog(model modelhelper.Model, id int) []ImageInfo {
	imageInfoList := []ImageInfo{}
	sql := fmt.Sprintf(`select i.id, id.name, i.url, i.description, u.nickname from image i, user u where i.creater = u.id and i.id in (
		select src from resource_relative where dst = %d and dstType = %d and srcType = %d )`, id, base.CATALOG, base.IMAGE)
	model.Query(sql)
	
	for model.Next() {
		image := ImageInfo{}
		model.GetValue(&image.Id, &image.Name, &image.Url, &image.Desc, &image.Creater)
		
		imageInfoList = append(imageInfoList, image)
	}
	
	for index, info := range imageInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.IMAGE)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			imageInfoList[index].Catalog = append(imageInfoList[index].Catalog, name)
		}
	}
	
	return imageInfoList
}

func QueryImageByRang(model modelhelper.Model, begin int,offset int) []ImageInfo {
	imageInfoList := []ImageInfo{}
	sql := fmt.Sprintf(`select i.id, id.name, i.url, i.description, u.nickname from image i, user u where i.creater = u.id and i.id in (
		select src from resource_relative where dstType = %d and srcType = %d ) and i.id >= %d limit %d`, base.CATALOG, base.IMAGE, begin, offset)
	model.Query(sql)

	for model.Next() {
		image := ImageInfo{}
		model.GetValue(&image.Id, &image.Url, &image.Desc, &image.Creater)
		imageInfoList = append(imageInfoList, image)			
	}
	
	for index, info := range imageInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.IMAGE)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			imageInfoList[index].Catalog = append(imageInfoList[index].Catalog, name)
		}
	}
	
	return imageInfoList
}

func QueryImageById(model modelhelper.Model, id int) (Image, bool) {
	image := &image{}
	
	sql := fmt.Sprintf(`select id, name, url, description, creater from image where id = %d`, id)
	model.Query(sql)

	result := false
	if model.Next() {
		model.GetValue(&image.id, &image.name, &image.url, &image.desc, &image.creater)
		result = true
	}

	if result {
		sql = fmt.Sprintf(`select dst from resource_relative where src = %d and srcType = %d and dstType =%d`, image.id, base.IMAGE, base.CATALOG)
		pid := -1
		model.Query(sql)
		for model.Next() {
			model.GetValue(&pid)
			image.catalog = append(image.catalog, pid)
		}
	}
	
	return image, result	
}

func DeleteImage(model modelhelper.Model, id int) bool {
	model.BeginTransaction()
	
	sql := fmt.Sprintf(`delete from image where id =%d`, id)	
	_, result := model.Execute(sql)
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
	model.Query(sql)

	result := false;
	if model.Next() {
		var id = 0
		model.GetValue(&id)
		result = true
	}

	model.BeginTransaction()
	
	if !result {
		// insert
		sql = fmt.Sprintf(`insert into image (name,url,description,creater) values ('%s','%s','%s',%d)`, image.Name(), image.Url(), image.Desc(), image.Creater())
		_, result = model.Execute(sql)
		sql = fmt.Sprintf(`select id from image where url='%s' and description ='%s' and creater=%d`, image.Url(), image.Desc(), image.Creater())
		
		id := -1
		model.Query(sql)
		result = false
		if model.Next() {
			model.GetValue(&id)
			image.SetId(id)
			result = true
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update image set name='%s', url ='%s', description='%s', creater=%d where id=%d`, image.Name(), image.Url(), image.Desc(), image.Creater(), image.Id())
		_, result = model.Execute(sql)
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


