package dal


import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/content/model"
	"magiccenter/kernel/account/dal"
)

func QueryAllImage(helper modelhelper.Model) []model.ImageDetail {
	imageList := []model.ImageDetail{}
	sql := fmt.Sprintf(`select id, name, url, description, creater from image`)
	helper.Query(sql)

	for helper.Next() {
		image := model.ImageDetail{}
		helper.GetValue(&image.Id, &image.Name, &image.Url, &image.Desc, &image.Creater.Id)
		
		imageList = append(imageList, image)
	}
	
	for i, _ := range imageList {
		image := &imageList[i]
		
		user, found := dal.QueryUserById(helper, image.Creater.Id)
		if found {
			image.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, image.Id, model.IMAGE)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			
			image.Catalog = append(image.Catalog, catalog)
		}		
	}
	
	return imageList
}

func QueryImageByCatalog(helper modelhelper.Model, id int) []model.ImageDetail {
	imageList := []model.ImageDetail{}
	
	resList := QueryReferenceResource(helper, id, model.CATALOG, model.IMAGE)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name, url, description, creater from image where id =%d`, r.RId())
		helper.Query(sql)
		
		if helper.Next() {
			image := model.ImageDetail{}
			helper.GetValue(&image.Id, &image.Name, &image.Url, &image.Desc, &image.Creater.Id)			
			imageList = append(imageList, image)
		}
	}
	
	for i, _ := range imageList {
		image := &imageList[i]
		
		user, found := dal.QueryUserById(helper, image.Creater.Id)
		if found {
			image.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, id, model.IMAGE)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			
			image.Catalog = append(image.Catalog, catalog)
		}		
	}
			
	return imageList	
}

func QueryImageByRang(helper modelhelper.Model, begin int,offset int) []model.ImageDetail {
	imageList := []model.ImageDetail{}
	sql := fmt.Sprintf(`select id, name, url, description, creater from image order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		image := model.ImageDetail{}
		helper.GetValue(&image.Id, &image.Name, &image.Url, &image.Desc, &image.Creater.Id)
		
		imageList = append(imageList, image)
	}
	
	for i, _ := range imageList {
		image := &imageList[i]
		
		user, found := dal.QueryUserById(helper, image.Creater.Id)
		if found {
			image.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, image.Id, model.IMAGE)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			
			image.Catalog = append(image.Catalog, catalog)
		}		
	}	
	
	return imageList
}

func QueryImageById(helper modelhelper.Model, id int) (model.ImageDetail, bool) {
	image := model.ImageDetail{}
	
	sql := fmt.Sprintf(`select id, name, url, description, creater from image where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&image.Id, &image.Name, &image.Url, &image.Desc, &image.Creater.Id)
		result = true
	}

	if result {
		user, found := dal.QueryUserById(helper, image.Creater.Id)
		if found {
			image.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, id, model.IMAGE)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			
			image.Catalog = append(image.Catalog, catalog)
		}
	}
	
	return image, result
}

func DeleteImageById(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf(`delete from image where id =%d`, id)	
	num, result := helper.Execute(sql)
	if num > 0 && result {
		img := model.ImageDetail{}
		img.Id = id
		result  = DeleteResource(helper, &img)
	}
		
	return result	
}


func SaveImage(helper modelhelper.Model, image model.ImageDetail) bool {
	sql := fmt.Sprintf(`select id from image where id=%d`, image.Id)
	helper.Query(sql)

	result := false;
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into image (name,url,description,creater) values ('%s','%s','%s',%d)`, image.Name, image.Url, image.Desc, image.Creater.Id)
		_, result = helper.Execute(sql)
		sql = fmt.Sprintf(`select id from image where url='%s' and description ='%s' and creater=%d`, image.Url, image.Desc, image.Creater.Id)
		
		helper.Query(sql)
		result = false
		if helper.Next() {
			helper.GetValue(&image.Id)
			result = true
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update image set name='%s', url ='%s', description='%s', creater=%d where id=%d`, image.Name, image.Url, image.Desc, image.Creater.Id, image.Id)
		_, result = helper.Execute(sql)
	}
	
	if result {
		result = SaveResource(helper, &image)
	}
		
	return result	
}


