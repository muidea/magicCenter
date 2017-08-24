package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
)

// CreateCorsHandler 新建CorsHandler
func CreateCorsHandler(configuration common.Configuration) common.CorsHandler {
	i := impl{
		configuration: configuration}

	return &i
}

type impl struct {
	configuration common.Configuration
}

func (i *impl) CheckCors(res http.ResponseWriter, req *http.Request) bool {
	header := res.Header()

	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	header.Add("Access-Control-Allow-Credentials", "true")
	header.Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	return true
}
