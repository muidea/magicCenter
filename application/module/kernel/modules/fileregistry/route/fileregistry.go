package route

import (
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/fileregistry/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendFileRegistryRoute 追加FileRegistry路由
func AppendFileRegistryRoute(routes []common.Route, fileRegistryHandler common.FileRegistryHandler) []common.Route {
	route := createUploadFileRoute(fileRegistryHandler)
	routes = append(routes, route)

	rt := createDeleteFileRoute(fileRegistryHandler)
	routes = append(routes, rt)

	return routes
}

func createUploadFileRoute(fileRegistryHandler common.FileRegistryHandler) common.Route {
	return &uploadFileRoute{fileRegistryHandler: fileRegistryHandler}
}

func createDeleteFileRoute(fileRegistryHandler common.FileRegistryHandler) common.Route {
	return &deleteFileRoute{fileRegistryHandler: fileRegistryHandler}
}

type uploadFileRoute struct {
	fileRegistryHandler common.FileRegistryHandler
}

func (i *uploadFileRoute) Method() string {
	return common.POST
}

func (i *uploadFileRoute) Pattern() string {
	return net.JoinURL(def.URL, "/fileregistry/")
}

func (i *uploadFileRoute) Handler() interface{} {
	return i.uploadFileHandler
}

func (i *uploadFileRoute) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("uploadFileHandler")

	i.fileRegistryHandler.UploadFile(w, r)
}

type deleteFileRoute struct {
	fileRegistryHandler common.FileRegistryHandler
}

func (i *deleteFileRoute) Method() string {
	return common.DELETE
}

func (i *deleteFileRoute) Pattern() string {
	return net.JoinURL(def.URL, "/fileregistry/[a-zA-Z]+[a-zA-Z0-9]*")
}

func (i *deleteFileRoute) Handler() interface{} {
	return i.deleteFileHandler
}

func (i *deleteFileRoute) deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteFileHandler")

	i.fileRegistryHandler.DeleteFile(w, r)
}
