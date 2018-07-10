package resource

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/foundation/util"
)

func loadResourceOID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(oid) from common_resource`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// Resource 资源对象
// 用于表示可用于访问的信息(article,catalog,image,link)
type Resource interface {
	ID() int
	// RId 资源对应信息的ID
	RId() int
	// RName 资源名称
	RName() string
	// RDescription 资源描述
	RDescription() string
	// RType 资源类型
	RType() string
	// RCreateDate 创建时间
	RCreateDate() string
	// RRelative 关联的资源
	Relative() []Resource
	// ROwner 资源拥有者
	ROwner() int
	// UpdateName 更新Name
	UpdateName(name string)
	// UpdateDescription 更新Description
	UpdateDescription(desc string)
	// ResetRelative 重置关联资源
	ResetRelative()
	// AppendRelative 追加关联资源
	AppendRelative(r Resource)
}

// CreateSimpleRes 创建新的资源
func CreateSimpleRes(rID int, rType, rName, rDescription, rCreateDate string, rOwner int) Resource {
	res := &simpleRes{oid: allocResourceOID()}
	res.rid = rID
	res.rType = rType
	res.rName = rName
	res.rDescription = rDescription
	res.rCreateDate = rCreateDate
	res.rOwner = rOwner

	return res
}

// simpleRes 简单资源对象
type simpleRes struct {
	oid          int
	rid          int
	rName        string
	rDescription string
	rType        string
	rCreateDate  string
	rOwner       int
	relative     []Resource
}

func (s *simpleRes) ID() int {
	return s.oid
}

// RId 资源ID
func (s *simpleRes) RId() int {
	return s.rid
}

// RName 资源名
func (s *simpleRes) RName() string {
	return s.rName
}

// RDescription 资源描述
func (s *simpleRes) RDescription() string {
	return s.rDescription
}

// RType 资源类型
func (s *simpleRes) RType() string {
	return s.rType
}

// RCreateDate 创建时间
func (s *simpleRes) RCreateDate() string {
	return s.rCreateDate
}

// Relative 相关联的资源
func (s *simpleRes) Relative() []Resource {
	return s.relative
}

// ROwner 资源拥有者
func (s *simpleRes) ROwner() int {
	return s.rOwner
}

func (s *simpleRes) UpdateName(name string) {
	s.rName = name
}

func (s *simpleRes) UpdateDescription(desc string) {
	s.rDescription = desc
}

func (s *simpleRes) ResetRelative() {
	s.relative = []Resource{}
}

// AppendRelative 追加关联对象
func (s *simpleRes) AppendRelative(r Resource) {
	s.relative = append(s.relative, r)
}

func (s *simpleRes) setID(id int) {
	s.oid = id
}

// QueryResourceByID 查询资源
func QueryResourceByID(helper dbhelper.DBHelper, rid int, rType string) (Resource, bool) {
	sql := fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where id =%d and type ='%s'`, rid, rType)
	helper.Query(sql)

	res := simpleRes{}
	result := false
	if helper.Next() {
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		result = true
	}
	helper.Finish()

	if result {
		res.relative = relativeResource(helper, res.oid)
	}

	return &res, result
}

// QueryResourceByName 查询资源
func QueryResourceByName(helper dbhelper.DBHelper, name, rType string) (Resource, bool) {
	sql := fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where name ='%s' and type ='%s'`, name, rType)
	helper.Query(sql)

	res := simpleRes{}
	result := false
	if helper.Next() {
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		result = true
	}
	helper.Finish()

	if result {
		res.relative = relativeResource(helper, res.oid)
	}

	return &res, result
}

// QueryResourceByType 查询指定类型的资源
func QueryResourceByType(helper dbhelper.DBHelper, rType string) []Resource {
	resList := []simpleRes{}

	sql := fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where type ='%s' order by type`, rType)
	helper.Query(sql)
	defer helper.Finish()
	for helper.Next() {
		res := simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)

		resList = append(resList, res)
	}

	retVal := []Resource{}
	for idx := range resList {
		cur := resList[idx]
		cur.relative = relativeResource(helper, cur.oid)

		retVal = append(retVal, &cur)
	}

	return retVal
}

// QueryResourceByUser 查询指定用户的资源
func QueryResourceByUser(helper dbhelper.DBHelper, uids []int) []Resource {
	resList := []simpleRes{}

	UserStr := util.IntArray2Str(uids)

	sql := fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where owner in ('%s') order by type`, UserStr)
	helper.Query(sql)
	defer helper.Finish()
	for helper.Next() {
		res := simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)

		resList = append(resList, res)
	}

	retVal := []Resource{}
	for idx := range resList {
		cur := resList[idx]
		cur.relative = relativeResource(helper, cur.oid)

		retVal = append(retVal, &cur)
	}

	return retVal
}

// relativeResource 查询关联的资源,即以oid的子资源
func relativeResource(helper dbhelper.DBHelper, oid int) []Resource {
	sql := fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.dst and rr.src =%d`, oid)
	helper.Query(sql)
	defer helper.Finish()

	resultList := []Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resultList = append(resultList, res)
	}

	return resultList
}

// QueryRelativeResource 查询关联的资源
func QueryRelativeResource(helper dbhelper.DBHelper, rid int, rType string) []Resource {
	oid := -1
	sql := fmt.Sprintf(`select oid from common_resource where id=%d and type='%s'`, rid, rType)
	helper.Query(sql)

	ok := false
	if helper.Next() {
		helper.GetValue(&oid)
		ok = true
	}
	helper.Finish()

	if ok {
		return relativeResource(helper, oid)
	}

	return []Resource{}
}

// referenceResource 查询引用了指定Res的资源列表
// rID Res ID
// rType Res 类型
// referenceType 待查询的资源类型，值为""表示查询所有类型
func referenceResource(helper dbhelper.DBHelper, oid int, referenceType string) []Resource {
	sql := ""
	if referenceType == "" {
		sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d`, oid)
	} else {
		sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and r.type ='%s'`, oid, referenceType)
	}
	helper.Query(sql)
	defer helper.Finish()

	resultList := []Resource{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resultList = append(resultList, res)
	}

	return resultList
}

// QueryReferenceResource 查询引用了指定Res的资源列表
// rID Res ID
// rType Res 类型
// referenceType 待查询的资源类型，值为""表示查询所有类型
func QueryReferenceResource(helper dbhelper.DBHelper, rID int, rType, referenceType string) []Resource {
	oid := -1
	sql := fmt.Sprintf(`select oid from common_resource where id=%d and type='%s'`, rID, rType)
	helper.Query(sql)

	ok := false
	if helper.Next() {
		helper.GetValue(&oid)
		ok = true
	}
	helper.Finish()

	if ok {
		return referenceResource(helper, oid, referenceType)
	}

	return []Resource{}
}

// CreateResource 新建资源
func CreateResource(helper dbhelper.DBHelper, res Resource, enableTransaction bool) bool {
	if !enableTransaction {
		helper.BeginTransaction()
	}

	result := false
	for {
		result = false
		sql := fmt.Sprintf(`select oid from common_resource where id=%d and type='%s'`, res.RId(), res.RType())
		helper.Query(sql)

		if helper.Next() {
			helper.Finish()
			// 说明对应的资源已经存在
			break
		}
		helper.Finish()

		// insert
		sql = fmt.Sprintf(`insert into common_resource (oid, id,name,description,type,createtime,owner) values (%d, %d, '%s', '%s', '%s', '%s', %d)`, res.ID(), res.RId(), res.RName(), res.RDescription(), res.RType(), res.RCreateDate(), res.ROwner())
		_, result = helper.Execute(sql)
		if !result {
			// 插入失败
			break
		}

		result = saveResourceRelative(helper, res)
		break
	}

	if !enableTransaction {
		if result {
			helper.Commit()
		} else {
			helper.Rollback()
		}
	}

	return result
}

// SaveResource 保存资源
func SaveResource(helper dbhelper.DBHelper, res Resource, enableTransaction bool) bool {
	if !enableTransaction {
		helper.BeginTransaction()
	}

	sql := fmt.Sprintf(`update common_resource set name ='%s', description ='%s', createtime ='%s', owner=%d where oid=%d`, res.RName(), res.RDescription(), res.RCreateDate(), res.ROwner(), res.ID())

	// 这里只需要没有出错就可以了，不需要判断是否真实的更新了记录
	// 原因是由于resource本身并没有变化，所以update是没有记录更新的
	_, result := helper.Execute(sql)
	if result {
		saveResourceRelative(helper, res)
	}

	if !enableTransaction {
		if result {
			helper.Commit()
		} else {
			helper.Rollback()
		}
	}

	return result
}

// DeleteResource 删除资源
func DeleteResource(helper dbhelper.DBHelper, res Resource, enableTransaction bool) bool {
	result := false
	if !enableTransaction {
		helper.BeginTransaction()
	}

	for {
		sql := fmt.Sprintf(`delete from common_resource where oid=%d`, res.ID())
		_, result = helper.Execute(sql)
		if result {
			result = deleteResourceRelative(helper, res)
		}

		break
	}

	if !enableTransaction {
		if result {
			helper.Commit()
		} else {
			helper.Rollback()
		}
	}

	return result
}

// 保存关联资源
func saveResourceRelative(helper dbhelper.DBHelper, res Resource) bool {
	result := true

	deleteResourceRelative(helper, res)

	for _, rr := range res.Relative() {
		result = false

		sql := fmt.Sprintf(`select oid from common_resource where oid=%d`, rr.ID())
		helper.Query(sql)

		found := false
		if helper.Next() {
			found = true
		}
		helper.Finish()

		if !found {
			break
		}

		// insert
		sql = fmt.Sprintf(`insert into common_resource_relative (src, dst) values (%d, %d)`, res.ID(), rr.ID())
		_, result = helper.Execute(sql)
		if !result {
			break
		}
	}

	return result
}

// 删除关联资源
func deleteResourceRelative(helper dbhelper.DBHelper, res Resource) bool {
	sql := fmt.Sprintf(`delete from common_resource_relative where src=%d`, res.ID())
	_, result := helper.Execute(sql)

	return result
}

// GetLastResource 获取最新的资源
func GetLastResource(helper dbhelper.DBHelper, count int) []Resource {
	sql := fmt.Sprintf(`select id, name, description, type, createtime, owner from common_resource order by createtime desc limit %d`, count)
	helper.Query(sql)
	defer helper.Finish()

	resList := []simpleRes{}
	for helper.Next() {
		res := simpleRes{}
		helper.GetValue(&res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resList = append(resList, res)
	}

	retVal := []Resource{}
	for _, v := range resList {
		res := v
		res.relative = relativeResource(helper, v.oid)

		retVal = append(retVal, &res)
	}

	return retVal
}
