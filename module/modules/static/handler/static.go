package handler

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/static/util"
	"github.com/muidea/magicCommon/model"
)

// CreateStaticHandler 新建StaticHandler
func CreateStaticHandler(configuration common.Configuration, sessionRegistry common.SessionRegistry, moduleHub common.ModuleHub) common.StaticHandler {
	staticPath, _ := configuration.GetOption(model.StaticPath)

	var fileRegistryHandler common.FileRegistryHandler
	module, ok := moduleHub.FindModule(common.FileRegistryModuleID)
	if ok {
		entryPoint := module.EntryPoint()
		switch entryPoint.(type) {
		case common.FileRegistryHandler:
			fileRegistryHandler = entryPoint.(common.FileRegistryHandler)
		}
	}
	if fileRegistryHandler == nil {
		panic("can\\'t find fileregistryHandler")
	}

	i := impl{rootPath: staticPath, fileRegistryHandler: fileRegistryHandler}

	return &i
}

type impl struct {
	rootPath            string
	fileRegistryHandler common.FileRegistryHandler
}

func (i *impl) HandleResource(w http.ResponseWriter, r *http.Request) {
	fullPath := util.MergePath(i.rootPath, r.URL.Path)
	source := r.URL.Query().Get("source")
	if len(source) > 0 {
		rootPath, fileInfo, ok := i.fileRegistryHandler.FindFile(source)
		if ok {
			fullPath = path.Join(rootPath, fileInfo.FilePath)
		}
	}

	_, err := os.Stat(fullPath)
	if err == nil || os.IsExist(err) {
		filePath, fileName := path.Split(fullPath)
		dir := http.Dir(filePath)
		file, err := dir.Open(fileName)
		if err != nil {
			return
		}
		defer file.Close()

		fi, err := file.Stat()
		if err != nil || fi.IsDir() {
			return
		}

		http.ServeContent(w, r, fullPath, fi.ModTime(), file)
	} else {
		log.Printf("no found resource, path:%s", fullPath)
		http.Redirect(w, r, "/404.html", 404)
	}
}
