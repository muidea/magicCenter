package resource

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/def"
	common_util "muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
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

	CatalogUnit() *model.CatalogUnit
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

func (s *simpleRes) CatalogUnit() *model.CatalogUnit {
	return &model.CatalogUnit{ID: s.rid, Type: s.rType}
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
		res.relative, _ = relativeResource(helper, res.oid, nil)
	}

	return &res, result
}

// QueryResourceByName 查询资源
func QueryResourceByName(helper dbhelper.DBHelper, rName, rType string, filter *common_util.PageFilter) ([]Resource, int) {
	totalCount := 0
	resultList := []Resource{}

	sql := fmt.Sprintf(`select count(oid) from common_resource where name ='%s' and type ='%s'`, rName, rType)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		limitVal = filter.PageSize
		offsetVal = filter.PageSize * (filter.PageNum - 1)
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return resultList, totalCount
	}

	sql = fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where name ='%s' and type ='%s' order by createtime desc limit %d offset %d`, rName, rType, limitVal, offsetVal)
	helper.Query(sql)

	resList := []*simpleRes{}
	for helper.Next() {
		res := simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resList = append(resList, &res)
	}
	helper.Finish()

	for _, val := range resList {
		val.relative, _ = relativeResource(helper, val.oid, nil)

		resultList = append(resultList, val)
	}

	return resultList, totalCount
}

// QueryResourceByType 查询指定类型的资源
func QueryResourceByType(helper dbhelper.DBHelper, rType string, filter *def.Filter) ([]Resource, int) {
	totalCount := 0
	resultList := []Resource{}

	sql := fmt.Sprintf(`select count(oid) from common_resource where type ='%s' order by type`, rType)
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			sql = fmt.Sprintf(`select count(oid) from common_resource where type ='%s' and (description like '%%%s%%' or name like '%%%s%%') order by type`, rType, filterValue, filterValue)
		}
	}
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		if filter.PageFilter != nil {
			limitVal = filter.PageFilter.PageSize
			offsetVal = filter.PageFilter.PageSize * (filter.PageFilter.PageNum - 1)
		}
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return resultList, totalCount
	}

	resList := []*simpleRes{}
	sql = fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where type ='%s' order by type limit %d offset %d`, rType, limitVal, offsetVal)
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			sql = fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where type ='%s' and (description like '%%%s%%' or name like '%%%s%%') order by type limit %d offset %d`, rType, filterValue, filterValue, limitVal, offsetVal)
		}
	}
	helper.Query(sql)
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)

		resList = append(resList, res)
	}
	helper.Finish()

	for _, val := range resList {
		val.relative, _ = relativeResource(helper, val.oid, nil)

		resultList = append(resultList, val)
	}

	return resultList, totalCount
}

// QueryResourceByIDs 查询指定类型的资源
func QueryResourceByIDs(helper dbhelper.DBHelper, rIDs []int, rType string, filter *common_util.PageFilter) ([]Resource, int) {
	totalCount := 0
	resultList := []Resource{}
	if len(rIDs) == 0 {
		return resultList, totalCount
	}

	ids := common_util.IntArray2Str(rIDs)
	sql := fmt.Sprintf(`select count(oid) from common_resource where id in (%s) and type ='%s' order by type`, ids, rType)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		limitVal = filter.PageSize
		offsetVal = filter.PageSize * (filter.PageNum - 1)
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return resultList, totalCount
	}

	resList := []*simpleRes{}
	sql = fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where id in (%s) and type ='%s' order by type limit %d offset %d`, ids, rType, limitVal, offsetVal)
	helper.Query(sql)
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)

		resList = append(resList, res)
	}
	helper.Finish()

	for _, val := range resList {
		val.relative, _ = relativeResource(helper, val.oid, nil)

		resultList = append(resultList, val)
	}

	return resultList, totalCount
}

// QueryResourceByUser 查询指定用户的资源
func QueryResourceByUser(helper dbhelper.DBHelper, uids []int, filter *def.Filter) ([]Resource, int) {
	totalCount := 0
	resultList := []Resource{}

	userStr := common_util.IntArray2Str(uids)

	sql := fmt.Sprintf(`select count(oid) from common_resource where owner in (%s) order by type`, userStr)
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			sql = fmt.Sprintf(`select count(oid) from common_resource where owner in (%s) and (description like '%%%s%%' or name like '%%%s%%') order by type`, userStr, filterValue, filterValue)
		}
	}
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		if filter.PageFilter != nil {
			limitVal = filter.PageFilter.PageSize
			offsetVal = filter.PageFilter.PageSize * (filter.PageFilter.PageNum - 1)
		}
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return resultList, totalCount
	}

	resList := []*simpleRes{}
	sql = fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where owner in (%s) order by type limit %d offset %d`, userStr, limitVal, offsetVal)
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			sql = fmt.Sprintf(`select oid, id, name, description, type, createtime, owner from common_resource where owner in (%s) and (description like '%%%s%%' or name like '%%%s%%') order by type limit %d offset %d`, userStr, filterValue, filterValue, limitVal, offsetVal)
		}
	}
	helper.Query(sql)
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)

		resList = append(resList, res)
	}
	helper.Finish()

	for _, val := range resList {
		val.relative, _ = relativeResource(helper, val.oid, nil)

		resultList = append(resultList, val)
	}

	return resultList, totalCount
}

// relativeResource 查询关联的资源,即以oid的子资源
func relativeResource(helper dbhelper.DBHelper, oid int, filter *def.Filter) ([]Resource, int) {
	totalCount := 0
	resultList := []Resource{}

	sql := fmt.Sprintf(`select count(r.oid) from common_resource r, common_resource_relative rr where r.oid = rr.dst and rr.src =%d`, oid)
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			sql = fmt.Sprintf(`select count(r.oid) from common_resource r, common_resource_relative rr where r.oid = rr.dst and rr.src =%d and (r.description like '%%%s%%' or r.name like '%%%s%%')`, oid, filterValue, filterValue)
		}
	}
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		if filter.PageFilter != nil {
			limitVal = filter.PageFilter.PageSize
			offsetVal = filter.PageFilter.PageSize * (filter.PageFilter.PageNum - 1)
		}
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return resultList, totalCount
	}

	sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.dst and rr.src =%d order by r.createtime desc limit %d offset %d`, oid, limitVal, offsetVal)
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.dst and rr.src =%d and (r.description like '%%%s%%' or r.name like '%%%s%%') order by r.createtime desc limit %d offset %d`, oid, filterValue, filterValue, limitVal, offsetVal)
		}
	}
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resultList = append(resultList, res)
	}

	return resultList, totalCount
}

// QueryRelativeResource 查询关联的资源
func QueryRelativeResource(helper dbhelper.DBHelper, rid int, rType string, filter *def.Filter) ([]Resource, int) {
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
		return relativeResource(helper, oid, filter)
	}

	return []Resource{}, 0
}

// referenceResource 查询引用了指定Res的资源列表
// rID Res ID
// rType Res 类型
// referenceType 待查询的资源类型，值为""表示查询所有类型
func referenceResource(helper dbhelper.DBHelper, oid int, referenceType string, filter *def.Filter) ([]Resource, int) {
	totalCount := 0
	resultList := []Resource{}

	sql := ""
	if referenceType == "" {
		sql = fmt.Sprintf(`select count(r.oid) from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d`, oid)
	} else {
		sql = fmt.Sprintf(`select count(r.oid) from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and r.type ='%s'`, oid, referenceType)
	}
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			if referenceType == "" {
				sql = fmt.Sprintf(`select count(r.oid) from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and (r.description like '%%%s%%' or r.name like '%%%s%%') `, oid, filterValue, filterValue)
			} else {
				sql = fmt.Sprintf(`select count(r.oid) from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and r.type ='%s' and (r.description like '%%%s%%' or r.name like '%%%s%%') `, oid, referenceType, filterValue, filterValue)
			}
		}
	}
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := totalCount
	offsetVal := 0
	if filter != nil {
		if filter.PageFilter != nil {
			limitVal = filter.PageFilter.PageSize
			offsetVal = filter.PageFilter.PageSize * (filter.PageFilter.PageNum - 1)
		}
	}
	if offsetVal < 0 {
		offsetVal = 0
	}
	if offsetVal >= totalCount {
		return resultList, totalCount
	}

	if referenceType == "" {
		sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d order by r.createtime desc limit %d offset %d`, oid, limitVal, offsetVal)
	} else {
		sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and r.type ='%s' order by r.createtime desc limit %d offset %d`, oid, referenceType, limitVal, offsetVal)
	}
	if filter != nil {
		if filter.ContentFilter != nil && filter.ContentFilter.FilterValue != "" {
			filterValue := filter.ContentFilter.FilterValue
			if referenceType == "" {
				sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and (r.description like '%%%s%%' or r.name like '%%%s%%') order by r.createtime desc limit %d offset %d`, oid, filterValue, filterValue, limitVal, offsetVal)
			} else {
				sql = fmt.Sprintf(`select r.oid, r.id, r.name, r.description, r.type, r.createtime, r.owner from common_resource r, common_resource_relative rr where r.oid = rr.src and rr.dst = %d and r.type ='%s' and (r.description like '%%%s%%' or r.name like '%%%s%%') order by r.createtime desc limit %d offset %d`, oid, referenceType, filterValue, filterValue, limitVal, offsetVal)
			}
		}
	}
	helper.Query(sql)

	resList := []*simpleRes{}
	for helper.Next() {
		res := &simpleRes{}
		helper.GetValue(&res.oid, &res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resList = append(resList, res)
	}
	helper.Finish()

	for _, val := range resList {
		val.relative, _ = relativeResource(helper, val.oid, nil)

		resultList = append(resultList, val)
	}

	return resultList, totalCount
}

// QueryReferenceResource 查询引用了指定Res的资源列表
// rID Res ID
// rType Res 类型
// referenceType 待查询的资源类型，值为""表示查询所有类型
func QueryReferenceResource(helper dbhelper.DBHelper, rID int, rType, referenceType string, filter *def.Filter) ([]Resource, int) {
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
		return referenceResource(helper, oid, referenceType, filter)
	}

	return []Resource{}, 0
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
// []Resource 当前页内容
// int 总条数
func GetLastResource(helper dbhelper.DBHelper, count int, pageFilger *common_util.PageFilter) ([]Resource, int) {
	retVal := []Resource{}
	sql := fmt.Sprint("select count(id) from common_resource order by createtime desc")
	totalCount := 0
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&totalCount)
	}
	helper.Finish()

	limitVal := count
	offsetVal := 0
	if pageFilger != nil {
		limitVal = pageFilger.PageSize
		offsetVal = pageFilger.PageSize * (pageFilger.PageNum - 1)
	}
	if offsetVal < 0 {
		offsetVal = 0
	}

	if offsetVal >= totalCount {
		return retVal, totalCount
	}

	sql = fmt.Sprintf(`select id, name, description, type, createtime, owner from common_resource order by createtime desc limit %d offset %d`, limitVal, offsetVal)
	helper.Query(sql)

	resList := []simpleRes{}
	for helper.Next() {
		res := simpleRes{}
		helper.GetValue(&res.rid, &res.rName, &res.rDescription, &res.rType, &res.rCreateDate, &res.rOwner)
		resList = append(resList, res)
	}
	helper.Finish()

	for _, v := range resList {
		res := v
		res.relative, _ = relativeResource(helper, v.oid, nil)

		retVal = append(retVal, &res)
	}

	return retVal, totalCount
}
