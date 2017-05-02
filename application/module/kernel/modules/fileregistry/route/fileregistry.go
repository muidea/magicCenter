package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/api/def"
	"muidea.com/magicCenter/foundation/net"
)

// AppendFileRegistryRoute 追加FileRegistry路由
func AppendFileRegistryRoute(routes []common.Route) []common.Route {
	route := createUploadFileRoute()

	routes = append(routes, route)
	return routes
}

func createUploadFileRoute() common.Route {
	return &uploadFileRoute{}
}

type uploadFileRoute struct {
}

type uploadFileResult struct {
	common.Result
	FilePath string
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

	result := uploadFileResult{}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
