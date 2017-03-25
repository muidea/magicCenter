package dal

import "muidea.com/magicCenter/application/common/dbhelper"
import "muidea.com/magicCenter/application/common/model"
import "fmt"

// InsertAuthGroup 新增AuthGroup
func InsertAuthGroup(helper dbhelper.DBHelper, authGroup model.AuthGroup) (model.AuthGroup, bool) {
	sql := fmt.Sprintf("insert into authgroup (name, description, module) values ('%s','%s','%s')", authGroup.Name, authGroup.Description, authGroup.Module)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return authGroup, false
	}

	ok = false
	sql = fmt.Sprintf("select id from authgroup where name='%s' and module='%s'", authGroup.Name, authGroup.Module)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&authGroup.ID)
		ok = true
	}

	return authGroup, ok
}

// DeleteAuthGroup 删除指定AuthGroup
func DeleteAuthGroup(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from authgroup where id=%d", id)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// GetAllAuthGroup 获取所有AuthGroup
func GetAllAuthGroup(helper dbhelper.DBHelper) []model.AuthGroup {
	authGroups := []model.AuthGroup{}
	sql := fmt.Sprint("select id,name,description,module from authgroup")
	helper.Query(sql)
	for helper.Next() {
		authGroup := model.AuthGroup{}
		helper.GetValue(&authGroup.ID, &authGroup.Name, &authGroup.Description, &authGroup.Module)
		authGroups = append(authGroups, authGroup)
	}

	return authGroups
}
