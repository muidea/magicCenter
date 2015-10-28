package content

import (
	"fmt"
	"log"
	"muidea.com/dao"
	"webcenter/auth"
)

type Link struct {
	Id int
	Name string
	Url string
	Logo string
	Style int
	Creater auth.User
}

func newLink() Link {
	link := Link{}
	link.Id = -1
	link.Creater = auth.NewUser()

	return link
}

func GetAllLink(dao * dao.Dao) []Link {
	linkList := []Link{}
	sql := fmt.Sprintf("select id, name, url, logo, style, creater from link")
	if !dao.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
		return linkList
	}

	for dao.Next() {
		link := newLink()
		dao.GetField(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Style, &link.Creater.Id)
		
		linkList = append(linkList, link)
	}
	
	for i :=0; i<len(linkList); i++ {
		link := &linkList[i]
		if !link.Creater.Query(dao) {
			link.Creater, _ = auth.QueryDefaultUser(dao)
		}
	}
	
	return linkList
}

func (this *Link)Query(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id, name, url, logo, style, creater from link where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
		return false
	}

	result := false
	for dao.Next() {
		dao.GetField(&this.Id, &this.Name, &this.Url, &this.Logo, &this.Style, &this.Creater.Id)
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


func (this *Link)delete(dao *dao.Dao) {
	sql := fmt.Sprintf("delete from link where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete link failed, sql:%s", sql)
	}
}

func (this *Link)save(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id from link where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query link failed, sql:%s", sql)
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
		sql = fmt.Sprintf("insert into link (name,url,logo,style, creater) values ('%s','%s','%s',%d, %d)", this.Name, this.Url, this.Logo, this.Style, this.Creater.Id)
	} else {
		// modify
		sql = fmt.Sprintf("update link set name ='%s', url ='%s', logo='%s', style=%d, creater=%d where id=%d", this.Name, this.Url, this.Logo, this.Style, this.Creater.Id, this.Id)
	}
	
	result = dao.Execute(sql)
	
	return result
}
