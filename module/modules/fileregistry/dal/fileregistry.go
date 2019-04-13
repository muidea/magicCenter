package dal

import (
	"fmt"

	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCommon/model"
)

// SaveFileSummary 保存文件信息
func SaveFileSummary(helper dbhelper.DBHelper, fileSummary model.FileSummary) bool {
	sql := fmt.Sprintf("insert into common_fileregistry (filetoken, filename, filepath, uploaddate) values ('%s','%s','%s','%s')", fileSummary.FileToken, fileSummary.FileName, fileSummary.FilePath, fileSummary.UploadDate)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// RemoveFileSummary 删除文件信息
func RemoveFileSummary(helper dbhelper.DBHelper, fileToken string) bool {
	sql := fmt.Sprintf("delete from common_fileregistry where filetoken = '%s' and reserveflag = 0", fileToken)

	_, result := helper.Execute(sql)
	return result
}

// ReserveFileSummary 预留文件信息
func ReserveFileSummary(helper dbhelper.DBHelper, fileToken string) bool {
	sql := fmt.Sprintf("update common_fileregistry set reserveflag = 1 where filetoken ='%s'", fileToken)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// DisReserveFileSummary 取消预留文件信息
func DisReserveFileSummary(helper dbhelper.DBHelper, fileToken string) bool {
	sql := fmt.Sprintf("update common_fileregistry set reserveflag = 0 where filetoken ='%s'", fileToken)
	num, result := helper.Execute(sql)
	return num == 1 && result
}

// FindFileSummary 查找指定文件信息
func FindFileSummary(helper dbhelper.DBHelper, filetoken string) (model.FileSummary, bool) {
	sql := fmt.Sprintf("select filename, filepath, uploaddate, reserveflag from common_fileregistry where filetoken = '%s'", filetoken)
	helper.Query(sql)
	defer helper.Finish()

	fileSummary := model.FileSummary{}
	retVal := false
	if helper.Next() {
		fileSummary.FileToken = filetoken
		helper.GetValue(&fileSummary.FileName, &fileSummary.FilePath, &fileSummary.UploadDate, &fileSummary.ReserveFlag)
		retVal = true
	}

	return fileSummary, retVal
}
