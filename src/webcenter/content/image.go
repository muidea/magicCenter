package content

import (
	"fmt"
	"log"
	"muidea.com/dao"
	"webcenter/auth"
)

type Image struct {
	Id int
	Url string
	Desc string
	Creater auth.User
}

func newImage() Image {
	image := Image{}
	image.Id = -1
	image.Creater = auth.NewUser()

	return image
}

func GetAllImage(dao * dao.Dao) []Image {
	imageList := []Image{}
	sql := fmt.Sprintf("select id, url, description, creater from image")
	if !dao.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return imageList
	}

	for dao.Next() {
		image := newImage()
		dao.GetField(&image.Id, &image.Url, &image.Desc, &image.Creater.Id)
		
		imageList = append(imageList, image)
	}
	
	for i :=0; i<len(imageList); i++ {
		image := &imageList[i]
		if !image.Creater.Query(dao) {
			image.Creater, _ = auth.QueryDefaultUser(dao)
		}
	}
	
	return imageList
}

func (this *Image)Query(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id, url, description, creater from image where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return false
	}

	result := false
	for dao.Next() {
		dao.GetField(&this.Id, &this.Url, &this.Desc, &this.Creater.Id)
		result = true
	}
	
	if result {
		result = this.Creater.Query(dao)
		if !result {
			this.Creater, result = auth.QueryDefaultUser(dao)
		}
	}
	
	return result	
}


func (this *Image)delete(dao *dao.Dao) {
	sql := fmt.Sprintf("delete from image where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete image failed, sql:%s", sql)
	}
}

func (this *Image)save(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id from image where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query image failed, sql:%s", sql)
		return false
	}

	result := false;
	for dao.Next() {
		var id = 0
		result = dao.GetField(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `image` (url,description,creater) values ('%s','%s',%d)", this.Url, this.Desc, this.Creater.Id)
	} else {
		// modify
		sql = fmt.Sprintf("update `image` set url ='%s', description='%s', creater=%d where id=%d", this.Url, this.Desc, this.Creater.Id, this.Id)
	}
	
	log.Print(sql)
	
	result = dao.Execute(sql)
	
	return result
}
