package util

import "testing"

func TestMergePath(t *testing.T) {
	rootPath := "root"
	basePath := "base"
	url := "/static/12345/"

	fullPath := MergePath(rootPath, basePath, url)
	if fullPath != "root/base/12345/" {
		t.Errorf("MergePath failed, fullPath:%s", fullPath)
	}

	rootPath = "root"
	basePath = "base"
	url = "static/12345/"
	fullPath = MergePath(rootPath, basePath, url)
	if fullPath != "root/base/12345/" {
		t.Errorf("MergePath failed, fullPath:%s", fullPath)
	}

	rootPath = "root/"
	basePath = "/base/"
	url = "static/12345"
	fullPath = MergePath(rootPath, basePath, url)
	if fullPath != "root/base/12345" {
		t.Errorf("MergePath failed, fullPath:%s", fullPath)
	}
}
