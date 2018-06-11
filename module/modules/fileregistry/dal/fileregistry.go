package dal

import (
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/model"
)

// SaveFileSummary 保存文件信息
func SaveFileSummary(helper dbhelper.DBHelper, fileSummary model.FileSummary) bool {
	sql := fmt.Sprintf("insert into common_fileregistry (accesstoken, filename, filepath, uploaddate) values ('%s','%s','%s','%s')", fileSummary.AccessToken, fileSummary.FileName, fileSummary.FilePath, fileSummary.UploadDate)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// RemoveFileSummary 删除文件信息
func RemoveFileSummary(helper dbhelper.DBHelper, accessToken string) bool {
	sql := fmt.Sprintf("delete from common_fileregistry where accesstoken = '%s' and reserveflag = 0", accessToken)

	_, result := helper.Execute(sql)
	return result
}

// ReserveFileSummary 预留文件信息
func ReserveFileSummary(helper dbhelper.DBHelper, accessToken string) bool {
	sql := fmt.Sprintf("update common_fileregistry set reserveflag = 1 where accesstoken ='%s'", accessToken)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// DisReserveFileSummary 取消预留文件信息
func DisReserveFileSummary(helper dbhelper.DBHelper, accessToken string) bool {
	sql := fmt.Sprintf("update common_fileregistry set reserveflag = 0 where accesstoken ='%s'", accessToken)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// FindFileSummary 查找指定文件信息
func FindFileSummary(helper dbhelper.DBHelper, accesstoken string) (model.FileSummary, bool) {
	sql := fmt.Sprintf("select filename, filepath, uploaddate, reserveflag from common_fileregistry where accesstoken = '%s'", accesstoken)
	helper.Query(sql)
	defer helper.Finish()

	fileSummary := model.FileSummary{}
	retVal := false
	if helper.Next() {
		fileSummary.AccessToken = accesstoken
		helper.GetValue(&fileSummary.FileName, &fileSummary.FilePath, &fileSummary.UploadDate, &fileSummary.ReserveFlag)
		retVal = true
	}

	return fileSummary, retVal
}
