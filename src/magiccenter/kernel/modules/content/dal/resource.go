package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/content/model"
)

type simpleRes struct {
	rid   int
	rname string
	tid   int
	relative []model.Resource
}

func (this *simpleRes) RId() int {
	return this.rid
}

func (this *simpleRes) RName() string {
	return this.rname
}

func (this *simpleRes) RType() int {
	return this.tid
}

func (this *simpleRes) RRelative() []model.Resource {
	return this.relative
}

func CreateSimpleRes(rid, rtype int, rname string) model.Resource {
	res := &simpleRes{}
	res.rid = rid
	res.tid = rtype
	res.rname = rname
	
	return res
}

func QueryResource(helper modelhelper.Model, id int, tid int) (model.Resource, bool) {
	sql := fmt.Sprintf(`select id, type, name from resource where id =%d and type =%d`, id, tid)
	helper.Query(sql)
	
	res := simpleRes{}
	result := false
	if helper.Next() {
		helper.GetValue(&res.rid, &res.tid, &res.rname)
		result = true
	}
	
	if result {
		res.relative = QueryRelativeResource(helper, id, tid) 
	}
	
	return &res, result
}

func QueryRelativeResource(helper modelhelper.Model, id int, tid int) []model.Resource {
	sql := fmt.Sprintf(`select rr.dst id, rr.dstType type, r.name name from resource_relative rr, resource r where rr.dst = r.id and rr.dstType = r.type and rr.src =%d and rr.srcType =%d`, id, tid)
	helper.Query(sql)
	
	resultList := []model.Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.rid, &res.tid, &res.rname)		
		resultList = append(resultList, res)
	}
	
	return resultList
}

// 查询引用了指定Res的资源列表
// id Res ID
// tid Res 类型
// rType 待查询的资源类型，值为-1表示查询所有类型
func QueryReferenceResource(helper modelhelper.Model, id, tid, rtype int) []model.Resource {
	sql := ""
	if rtype == -1 {
		sql = fmt.Sprintf(`select r.id, r.type, r.name from resource r, resource_relative rr where r.id = rr.src and r.type = rr.srcType and rr.dst = %d and rr.dstType = %d`, id, tid)
	} else {
		sql = fmt.Sprintf(`select r.id, r.type, r.name from resource r, resource_relative rr where r.id = rr.src and r.type = rr.srcType and rr.dst = %d and rr.dstType = %d and rr.srcType =%d`, id, tid, rtype)
	}
	helper.Query(sql)
	
	resultList := []model.Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.rid, &res.tid, &res.rname)
		resultList = append(resultList, res)
	}
	
	return resultList
}

func SaveResource(helper modelhelper.Model, res model.Resource) bool {
	sql := fmt.Sprintf(`select id from resource where id=%d and type=%d`, res.RId(), res.RType())
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into resource (name,type,id) values ('%s',%d, %d)`, res.RName(), res.RType(), res.RId())
	} else {
		// modify
		sql = fmt.Sprintf(`update resource set name ='%s' where type=%d and id=%d`, res.RName(), res.RType(), res.RId())
	}

	_, result = helper.Execute(sql)
	if result {
		saveResourceRelative(helper, res)
	}

	return result
}

func DeleteResource(helper modelhelper.Model, res model.Resource) bool {
	sql := fmt.Sprintf(`delete from resource where type=%d and id=%d`, res.RType(), res.RId())
	_, result := helper.Execute(sql)
	if result {
		deleteResourceRelative(helper, res)
	} else {
		panic("execute failed")
	}

	return result
}

func saveResourceRelative(helper modelhelper.Model, res model.Resource) bool {
	result := false

	deleteResourceRelative(helper, res)

	for _, rr := range res.RRelative() {
		result = false				
		sql := fmt.Sprintf(`select id from resource_relative where src=%d and srcType=%d and dst=%d and dstType=%d`, res.RId(), res.RType(), rr.RId(), rr.RType())
		helper.Query(sql)

		if helper.Next() {
			var id = 0
			helper.GetValue(&id)
			result = true
		}

		if !result {
			// insert
			sql = fmt.Sprintf(`insert into resource_relative (src,srcType,dst,dstType) values (%d, %d, %d, %d)`, res.RId(), res.RType(), rr.RId(), rr.RType())
			_, result = helper.Execute(sql)
			if !result {
				panic("execute failed")
			}
		}
	}

	return result
}

func deleteResourceRelative(helper modelhelper.Model, res model.Resource) bool {
	sql := fmt.Sprintf(`delete from resource_relative where src=%d and srcType=%d`, res.RId(), res.RType())
	_, result := helper.Execute(sql)
	if !result {
		panic("execute failed")
	}

	return result
}

