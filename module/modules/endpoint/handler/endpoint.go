package handler

import (
	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/endpoint/dal"
	"muidea.com/magicCommon/model"
)

// CreateEndpointHandler 新建CASHandler
func CreateEndpointHandler() common.EndpointHandler {
	dbhelper, _ := dbhelper.NewHelper()

	i := impl{
		dbhelper: dbhelper}

	return &i
}

type impl struct {
	dbhelper dbhelper.DBHelper
}

func (i *impl) QueryAllEndpoint() []model.Endpoint {
	return dal.QueryAllEndpoint(i.dbhelper)
}

func (i *impl) QueryEndpointByID(id string) (model.Endpoint, bool) {
	return dal.QueryEndpointByID(i.dbhelper, id)
}

func (i *impl) InsertEndpoint(id, name, description string, user []int, status int, authToken string) (model.Endpoint, bool) {
	return dal.InsertEndpoint(i.dbhelper, id, name, description, user, status, authToken)
}

func (i *impl) UpdateEndpoint(endpoint model.Endpoint) (model.Endpoint, bool) {
	return dal.UpdateEndpoint(i.dbhelper, endpoint)
}

func (i *impl) DeleteEndpoint(id string) bool {
	return dal.DeleteEndpoint(i.dbhelper, id)
}

func (i *impl) GetSummary() model.EndpointSummary {
	result := model.EndpointSummary{}
	allEndpoint := i.QueryAllEndpoint()
	endpointItem := model.UnitSummary{Name: "终端", Type: "endpoint", Count: len(allEndpoint)}
	result = append(result, endpointItem)

	return result
}
