package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

// SaveFileInfo 保存文件信息
func SaveFileInfo(helper dbhelper.DBHelper, fileInfo model.FileInfo) bool {
	sql := fmt.Sprintf("insert into fileregistry (accesstoken, filename, filepath, uploaddate) values ('%s','%s','%s','%s')", fileInfo.AccessToken, fileInfo.FileName, fileInfo.FilePath, fileInfo.UploadDate)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// RemoveFileInfo 删除文件信息
func RemoveFileInfo(helper dbhelper.DBHelper, accessToken string) bool {
	sql := fmt.Sprintf("delete from fileregistry where accesstoken = '%s' and reserveflag = 0", accessToken)

	_, result := helper.Execute(sql)
	return result
}

// ReserveFileInfo 预留文件信息
func ReserveFileInfo(helper dbhelper.DBHelper, accessToken string) bool {
	sql := fmt.Sprintf("update fileregistry set reserveflag = 1 where accesstoken ='%s'", accessToken)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// DisReserveFileInfo 取消预留文件信息
func DisReserveFileInfo(helper dbhelper.DBHelper, accessToken string) bool {
	sql := fmt.Sprintf("update fileregistry set reserveflag = 0 where accesstoken ='%s'", accessToken)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// FindFileInfo 查找指定文件信息
func FindFileInfo(helper dbhelper.DBHelper, accesstoken string) (model.FileInfo, bool) {
	sql := fmt.Sprintf("select filename, filepath, uploaddate, reserveflag from fileregistry where accesstoken = '%s'", accesstoken)
	helper.Query(sql)

	fileInfo := model.FileInfo{}
	retVal := false
	if helper.Next() {
		fileInfo.AccessToken = accesstoken
		helper.GetValue(&fileInfo.FileName, &fileInfo.FilePath, &fileInfo.UploadDate, &fileInfo.ReserveFlag)
		retVal = true
	}

	return fileInfo, retVal
}
