package link


import (
	"fmt"
	"log"
	"webcenter/modelhelper"
)

type Link struct {
	Id int
	Name string
	Url string
	Logo string
	Style int
	Creater int
}

func newLink() Link {
	link := Link{}
	link.Id = -1
	link.Creater = -1

	return link
}

func QueryAllLink(model modelhelper.Model) []Link {
	linkList := []Link{}
	sql := fmt.Sprintf("select id, name, url, logo, style, creater from link")
	if !model.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
		return linkList
	}

	for model.Next() {
		link := newLink()
		model.GetValue(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Style, &link.Creater)
		
		linkList = append(linkList, link)
	}
		
	return linkList
}

func QueryLink(model modelhelper.Model, id int) (Link, bool) {
	link := Link{}
	sql := fmt.Sprintf("select id, name, url, logo, style, creater from link where id=%d", id)
	if !model.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
		return link, false
	}

	result := false
	for model.Next() {
		model.GetValue(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Style, &link.Creater)
		result = true
	}
	
	return link, result	
}

func DeleteLink(model modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from link where id =%d", id)
	result := model.Execute(sql)
	if !result {
		log.Printf("delete link failed, sql:%s", sql)
	}
	
	return result
}

func SaveLink(model modelhelper.Model, link Link) bool {
	sql := fmt.Sprintf("select id from link where id=%d", link.Id)
	if !model.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
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
		sql = fmt.Sprintf("insert into link (name,url,logo,style, creater) values ('%s','%s','%s',%d, %d)", link.Name, link.Url, link.Logo, link.Style, link.Creater)
	} else {
		// modify
		sql = fmt.Sprintf("update link set name ='%s', url ='%s', logo='%s', style=%d, creater=%d where id=%d", link.Name, link.Url, link.Logo, link.Style, link.Creater, link.Id)
	}
	
	result = model.Execute(sql)
	if !result {
		log.Printf("execute failed, sql:%s", sql)
	}
	
	return result
}


func (this *Link)Query(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id, name, url, logo, style, creater from link where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
		return false
	}

	result := false
	for model.Next() {
		model.GetValue(&this.Id, &this.Name, &this.Url, &this.Logo, &this.Style, &this.Creater)
		result = true
	}
	
	return result	
}


func (this *Link)delete(model modelhelper.Model) {
	sql := fmt.Sprintf("delete from link where id =%d", this.Id)
	if !model.Execute(sql) {
		log.Printf("delete link failed, sql:%s", sql)
	}
}

func (this *Link)save(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id from link where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
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
		sql = fmt.Sprintf("insert into link (name,url,logo,style, creater) values ('%s','%s','%s',%d, %d)", this.Name, this.Url, this.Logo, this.Style, this.Creater)
	} else {
		// modify
		sql = fmt.Sprintf("update link set name ='%s', url ='%s', logo='%s', style=%d, creater=%d where id=%d", this.Name, this.Url, this.Logo, this.Style, this.Creater, this.Id)
	}
	
	result = model.Execute(sql)
	if !result {
		log.Printf("execute failed, sql:%s", sql)
	}
	
	return result
}

