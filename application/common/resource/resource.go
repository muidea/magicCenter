package resource

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
)

// Resource 资源对象
// 用于表示可用于访问的信息(article,catalog,image,link)
type Resource interface {
	// RId 资源对应信息的ID
	RId() int
	// RName 资源名称
	RName() string
	// RType 资源类型
	RType() string
	// RRelative 关联的资源
	Relative() []Resource
	// AppendRelative 追加关联资源
	AppendRelative(r Resource)
}

// CreateSimpleRes 创建新的资源
func CreateSimpleRes(rID int, rType, rName string) Resource {
	res := &simpleRes{}
	res.rid = rID
	res.rtype = rType
	res.rname = rName

	return res
}

// simpleRes 简单资源对象
type simpleRes struct {
	rid      int
	rname    string
	rtype    string
	relative []Resource
}

// RId 资源ID
func (s *simpleRes) RId() int {
	return s.rid
}

// RName 资源名
func (s *simpleRes) RName() string {
	return s.rname
}

// RType 资源类型
func (s *simpleRes) RType() string {
	return s.rtype
}

// Relative 相关联的资源
func (s *simpleRes) Relative() []Resource {
	return s.relative
}

// AppendRelative 追加关联对象
func (s *simpleRes) AppendRelative(r Resource) {
	s.relative = append(s.relative, r)
}

// QueryResource 查询资源
func QueryResource(helper dbhelper.DBHelper, rid int, rtype string) (Resource, bool) {
	sql := fmt.Sprintf(`select id, type, name from common_resource where id =%d and type ='%s'`, rid, rtype)
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
func QueryRelativeResource(helper dbhelper.DBHelper, rid int, rtype string) []Resource {
	sql := fmt.Sprintf(`select distinct(r.oid), r.name, r.type, r.id from common_resource r, common_resource_relative rr where r.id = rr.dst and r.type = rr.dsttype and rr.src =%d and rr.srctype ='%s'`, rid, rtype)
	helper.Query(sql)

	resultList := []Resource{}
	for helper.Next() {
		res := &simpleRes{}
		oid := 0
		helper.GetValue(&oid, &res.rname, &res.rtype, &res.rid)
		resultList = append(resultList, res)
	}

	return resultList
}

// QueryReferenceResource 查询引用了指定Res的资源列表
// rID Res ID
// rType Res 类型
// referenceType 待查询的资源类型，值为""表示查询所有类型
func QueryReferenceResource(helper dbhelper.DBHelper, rID int, rType, referenceType string) []Resource {
	sql := ""
	if referenceType == "" {
		sql = fmt.Sprintf(`select r.id, r.type, r.name from common_resource r, common_resource_relative rr where r.id = rr.src and r.type = rr.srctype and rr.dst = %d and rr.dsttype = '%s'`, rID, rType)
	} else {
		sql = fmt.Sprintf(`select r.id, r.type, r.name from common_resource r, common_resource_relative rr where r.id = rr.src and r.type = rr.srctype and rr.dst = %d and rr.dsttype = '%s' and rr.srctype ='%s'`, rID, rType, referenceType)
	}
	helper.Query(sql)

	resultList := []Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.rid, &res.rtype, &res.rname)
		resultList = append(resultList, res)
	}

	return resultList
}

// SaveResource 保存资源
func SaveResource(helper dbhelper.DBHelper, res Resource) bool {
	sql := fmt.Sprintf(`select id from common_resource where id=%d and type='%s'`, res.RId(), res.RType())
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into common_resource (name,type,id) values ('%s','%s', %d)`, res.RName(), res.RType(), res.RId())
	} else {
		// modify
		sql = fmt.Sprintf(`update common_resource set name ='%s' where type='%s' and id=%d`, res.RName(), res.RType(), res.RId())
	}

	// 这里只需要没有出错就可以了，不需要判断是否真实的更新了记录
	// 原因是由于resource本身并没有变化，所以update是没有记录更新的
	_, result = helper.Execute(sql)
	if result {
		saveResourceRelative(helper, res)
	}

	return result
}

// DeleteResource 删除资源
func DeleteResource(helper dbhelper.DBHelper, res Resource) bool {
	sql := fmt.Sprintf(`delete from common_resource where type='%s' and id=%d`, res.RType(), res.RId())
	_, result := helper.Execute(sql)
	if result {
		deleteResourceRelative(helper, res)
	} else {
		panic("execute failed")
	}

	return result
}

// 保存关联资源
func saveResourceRelative(helper dbhelper.DBHelper, res Resource) bool {
	result := false

	deleteResourceRelative(helper, res)

	for _, rr := range res.Relative() {
		result = false
		sql := fmt.Sprintf(`select id from common_resource_relative where src=%d and srctype='%s' and dst=%d and dsttype='%s'`, res.RId(), res.RType(), rr.RId(), rr.RType())
		helper.Query(sql)

		if helper.Next() {
			var id = 0
			helper.GetValue(&id)
			result = true
		}

		if !result {
			// insert
			sql = fmt.Sprintf(`insert into common_resource_relative (src,srctype,dst,dsttype) values (%d, '%s', %d, '%s')`, res.RId(), res.RType(), rr.RId(), rr.RType())
			_, result = helper.Execute(sql)
			if !result {
				panic("execute failed")
			}
		}
	}

	return result
}

// 删除关联资源
func deleteResourceRelative(helper dbhelper.DBHelper, res Resource) bool {
	sql := fmt.Sprintf(`delete from common_resource_relative where src=%d and srctype='%s'`, res.RId(), res.RType())
	_, result := helper.Execute(sql)
	if !result {
		panic("execute failed")
	}

	return result
}
