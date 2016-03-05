package common

import (
	"log"
	"fmt"
	"webcenter/util/modelhelper"
)

type Resource interface {
	Id() int
	Name() string
	Type() int
	Relative() []Resource
}

type simpleRes struct {
	rid   int
	rname string
	tid   int
	relative []Resource
}

func (this *simpleRes) Id() int {
	return this.rid
}

func (this *simpleRes) Name() string {
	return this.rname
}

func (this *simpleRes) Type() int {
	return this.tid
}

func (this *simpleRes) Relative() []Resource {
	return this.relative
}

func (this *simpleRes) queryRelative(model modelhelper.Model, recurse bool) {
	sql := fmt.Sprintf(`select rr.dst id, rr.dstType type, r.name name from resource_relative rr, resource r where rr.dst = r.id and rr.dstType = r.type and rr.src =%d and rr.srcType =%d`,this.rid, this.tid)
	model.Query(sql)
	
	presList := []simpleRes{}
	for model.Next() {
		pres := simpleRes{}
		model.GetValue(&pres.rid, &pres.tid, &pres.rname)
		
		if recurse {
			presList = append(presList, pres)
		} else {
			this.relative = append(this.relative, &pres)
		}
	}
	
	if recurse {
		for _, r := range presList {
			r.queryRelative(model, recurse)
			
			this.relative = append(this.relative, &r)
		}
	}
}

func NewSimpleRes(id int, name string, tid int) Resource {
	res := simpleRes{}
	res.rid = id
	res.rname = name
	res.tid = tid
	res.relative = []Resource{}
	
	return &res
}

func QueryResource(model modelhelper.Model, id int, tid int, recurse bool) (Resource, bool) {
	sql := fmt.Sprintf(`select id, type, name from resource where id =%d and type =%d`, id, tid)
	model.Query(sql)
	
	res := simpleRes{}
	result := false
	if model.Next() {
		model.GetValue(&res.rid, &res.tid, &res.rname)
		result = true
	}
	
	res.queryRelative(model, recurse)
	return &res, result
}

func QueryReferenceResource(model modelhelper.Model, id int, tid int, recurse bool) []Resource {
	sql := fmt.Sprintf(`select r.id, r.type, r.name from resource r, resource_relative rr where r.id = rr.src and r.type = rr.srcType and rr.dst = %d and rr.dstType = %d`, id, tid)
	model.Query(sql)
	
	resultList := []Resource{}
	resList := []simpleRes{}
	for model.Next() {
		res := simpleRes{}
		model.GetValue(&res.rid, &res.tid, &res.rname)
		resList = append(resList, res)
	}
	
	for _, r := range resList {
		r.queryRelative(model,recurse)
		
		resultList = append(resultList, &r)
	}
	
	return resultList
}

func SaveResource(model modelhelper.Model, res Resource) bool {
	sql := fmt.Sprintf(`select id from resource where id=%d and type=%d`, res.Id(), res.Type())
	model.Query(sql)

	result := false
	for model.Next() {
		var id = 0
		model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into resource (name,type,id) values ('%s',%d, %d)`, res.Name(), res.Type(), res.Id())
	} else {
		// modify
		sql = fmt.Sprintf(`update resource set name ='%s' where type=%d and id=%d`, res.Name(), res.Type(), res.Id())
	}

	result = model.Execute(sql)
	if result {
		saveResourceRelative(model, res)
	}

	return result
}

func DeleteResource(model modelhelper.Model, res Resource) bool {
	sql := fmt.Sprintf(`delete from resource where type=%d and id=%d`, res.Type(), res.Id())
	result := model.Execute(sql)
	if result {
		deleteResourceRelative(model, res)
	} else {
		panic("execute failed")
	}

	return result
}

func saveResourceRelative(model modelhelper.Model, res Resource) bool {
	result := false

	deleteResourceRelative(model, res)

	for _, rr := range res.Relative() {
		result = false				
		sql := fmt.Sprintf(`select id from resource_relative where src=%d and srcType=%d and dst=%d and dstType=%d`, res.Id(), res.Type(), rr.Id(), rr.Type())
		model.Query(sql)

		for model.Next() {
			var id = 0
			model.GetValue(&id)
			result = true
		}

		log.Printf("result:%b, sql:%s", result, sql)

		if !result {
			// insert
			sql = fmt.Sprintf(`insert into resource_relative (src,srcType,dst,dstType) values (%d, %d, %d, %d)`, res.Id(), res.Type(), rr.Id(), rr.Type())
			result = model.Execute(sql)
			if !result {
				panic("execute failed")
			}
		}
	}

	return result
}

func deleteResourceRelative(model modelhelper.Model, res Resource) bool {
	sql := fmt.Sprintf(`delete from resource_relative where src=%d and srcType=%d`, res.Id(), res.Type())
	result := model.Execute(sql)
	if !result {
		panic("execute failed")
	}

	return result
}