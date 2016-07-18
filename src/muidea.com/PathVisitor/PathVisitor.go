package PathVisitor

import (
	"os"
	"path/filepath"
)

// WalkPath 遍历指定目录
func WalkPath(filePath string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		fileList = append(fileList, path)
		return nil
	})

	return fileList, err
}
