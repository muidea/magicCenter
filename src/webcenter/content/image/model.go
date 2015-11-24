package image


import (
	"fmt"
	"log"
	"webcenter/modelhelper"
)

type Image struct {
	Id int
	Url string
	Desc string
	Creater int
}

func newImage() Image {
	image := Image{}
	image.Id = -1
	image.Creater = -1

	return image
}

func GetAllImage(model modelhelper.Model) []Image {
	imageList := []Image{}
	sql := fmt.Sprintf("select id, url, description, creater from image")
	if !model.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return imageList
	}

	for model.Next() {
		image := newImage()
		model.GetValue(&image.Id, &image.Url, &image.Desc, &image.Creater)
		
		imageList = append(imageList, image)
	}
		
	return imageList
}


func DeleteImage(model modelhelper.Model, id int) {
	sql := fmt.Sprintf("delete from image where id =%d", id)
	if !model.Execute(sql) {
		log.Printf("delete image failed, sql:%s", sql)
	}
}


func SaveImage(model modelhelper.Model, image Image) bool {
	sql := fmt.Sprintf("select id from image where id=%d", image.Id)
	if !model.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `image` (url,description,creater) values ('%s','%s',%d)", image.Url, image.Desc, image.Creater)
	} else {
		// modify
		sql = fmt.Sprintf("update `image` set url ='%s', description='%s', creater=%d where id=%d", image.Url, image.Desc, image.Creater, image.Id)
	}
	
	log.Print(sql)
	
	result = model.Execute(sql)
	
	return result
}



func (this *Image)Query(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id, url, description, creater from image where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return false
	}

	result := false
	for model.Next() {
		model.GetValue(&this.Id, &this.Url, &this.Desc, &this.Creater)
		result = true
	}
		
	return result	
}


func (this *Image)delete(model modelhelper.Model) {
	sql := fmt.Sprintf("delete from image where id =%d", this.Id)
	if !model.Execute(sql) {
		log.Printf("delete image failed, sql:%s", sql)
	}
}

func (this *Image)save(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id from image where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `image` (url,description,creater) values ('%s','%s',%d)", this.Url, this.Desc, this.Creater)
	} else {
		// modify
		sql = fmt.Sprintf("update `image` set url ='%s', description='%s', creater=%d where id=%d", this.Url, this.Desc, this.Creater, this.Id)
	}
	
	log.Print(sql)
	
	result = model.Execute(sql)
	
	return result
}


