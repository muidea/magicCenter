package dal

import (
	"fmt"
	"magiccenter/resource"
	"magiccenter/util/modelhelper"

	"muidea.com/util"
)

// SimpleRes 简单资源对象
type simpleRes struct {
	rid      int
	rname    string
	rtype    string
	relative []resource.Resource
}

func createSimpleRes(rid int, rtype, rname string) simpleRes {
	res := simpleRes{}
	res.rid = rid
	res.rtype = rtype
	res.rname = rname

	return res
}

func (s *simpleRes) RId() int {
	return s.rid
}

func (s *simpleRes) RName() string {
	return s.rname
}

func (s *simpleRes) RType() string {
	return s.rtype
}

func (s *simpleRes) URL() string {
	param := fmt.Sprintf("id=%d", s.rid)

	return util.JoinURL(s.rname, param)
}

func (s *simpleRes) RRelative() []resource.Resource {
	return s.relative
}

// QueryResource 查询资源
func QueryResource(helper modelhelper.Model, rid int, rtype string) (resource.Resource, bool) {
	sql := fmt.Sprintf(`select id, type, name from resource where id =%d and type ='%s'`, rid, rtype)
	helper.Query(sql)

	res := simpleRes{}
	result := false
	if helper.Next() {
		helper.GetValue(&res.rid, &res.rtype, &res.rname)
		result = true
	}

	if result {
		res.relative = QueryRelativeResource(helper, rid, rtype)
	}

	return &res, result
}

// QueryRelativeResource 查询关联的资源
func QueryRelativeResource(helper modelhelper.Model, rid int, rtype string) []resource.Resource {
	sql := fmt.Sprintf(`select rr.dst id, rr.dstType type, r.name name from resource_relative rr, resource r where rr.dst = r.id and rr.dstType = r.type and rr.src =%d and rr.srcType ='%s'`, rid, rtype)
	helper.Query(sql)

	resultList := []resource.Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.rid, &res.rtype, &res.rname)
		resultList = append(resultList, res)
	}

	return resultList
}

// QueryReferenceResource 查询引用了指定Res的资源列表
// rID Res ID
// rType Res 类型
// referenceType 待查询的资源类型，值为""表示查询所有类型
func QueryReferenceResource(helper modelhelper.Model, rID int, rType, referenceType string) []resource.Resource {
	sql := ""
	if referenceType == "" {
		sql = fmt.Sprintf(`select r.id, r.type, r.name from resource r, resource_relative rr where r.id = rr.src and r.type = rr.srcType and rr.dst = %d and rr.dstType = '%s'`, rID, rType)
	} else {
		sql = fmt.Sprintf(`select r.id, r.type, r.name from resource r, resource_relative rr where r.id = rr.src and r.type = rr.srcType and rr.dst = %d and rr.dstType = '%s' and rr.srcType ='%s'`, rID, rType, referenceType)
	}
	helper.Query(sql)

	resultList := []resource.Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.rid, &res.rtype, &res.rname)
		resultList = append(resultList, res)
	}

	return resultList
}

// SaveResource 保存资源
func SaveResource(helper modelhelper.Model, res resource.Resource) bool {
	sql := fmt.Sprintf(`select id from resource where id=%d and type='%s'`, res.RId(), res.RType())
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into resource (name,type,id) values ('%s','%s', %d)`, res.RName(), res.RType(), res.RId())
	} else {
		// modify
		sql = fmt.Sprintf(`update resource set name ='%s' where type='%s' and id=%d`, res.RName(), res.RType(), res.RId())
	}

	_, result = helper.Execute(sql)
	if result {
		saveResourceRelative(helper, res)
	}

	return result
}

// DeleteResource 删除资源
func DeleteResource(helper modelhelper.Model, res resource.Resource) bool {
	sql := fmt.Sprintf(`delete from resource where type='%s' and id=%d`, res.RType(), res.RId())
	_, result := helper.Execute(sql)
	if result {
		deleteResourceRelative(helper, res)
	} else {
		panic("execute failed")
	}

	return result
}

// 保存关联资源
func saveResourceRelative(helper modelhelper.Model, res resource.Resource) bool {
	result := false

	deleteResourceRelative(helper, res)

	for _, rr := range res.RRelative() {
		result = false
		sql := fmt.Sprintf(`select id from resource_relative where src=%d and srcType='%s' and dst=%d and dstType='%s'`, res.RId(), res.RType(), rr.RId(), rr.RType())
		helper.Query(sql)

		if helper.Next() {
			var id = 0
			helper.GetValue(&id)
			result = true
		}

		if !result {
			// insert
			sql = fmt.Sprintf(`insert into resource_relative (src,srcType,dst,dstType) values (%d, '%s', %d, '%s')`, res.RId(), res.RType(), rr.RId(), rr.RType())
			_, result = helper.Execute(sql)
			if !result {
				panic("execute failed")
			}
		}
	}

	return result
}

// 删除关联资源
func deleteResourceRelative(helper modelhelper.Model, res resource.Resource) bool {
	sql := fmt.Sprintf(`delete from resource_relative where src=%d and srcType='%s'`, res.RId(), res.RType())
	_, result := helper.Execute(sql)
	if !result {
		panic("execute failed")
	}

	return result
}
