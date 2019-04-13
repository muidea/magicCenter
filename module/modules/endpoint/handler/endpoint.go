package handler

import (
	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/module/modules/endpoint/dal"
	"github.com/muidea/magicCommon/model"
)

// CreateEndpointHandler 新建CASHandler
func CreateEndpointHandler() common.EndpointHandler {

	i := impl{}

	return &i
}

type impl struct {
}

func (i *impl) QueryAllEndpoint() []model.Endpoint {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllEndpoint(dbhelper)
}

func (i *impl) QueryEndpointByID(id string) (model.Endpoint, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryEndpointByID(dbhelper, id)
}

func (i *impl) InsertEndpoint(id, name, description string, user []int, status int, authToken string) (model.Endpoint, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.InsertEndpoint(dbhelper, id, name, description, user, status, authToken)
}

func (i *impl) UpdateEndpoint(endpoint model.Endpoint) (model.Endpoint, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.UpdateEndpoint(dbhelper, endpoint)
}

func (i *impl) DeleteEndpoint(id string) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteEndpoint(dbhelper, id)
}

func (i *impl) GetSummary() model.EndpointSummary {
	result := model.EndpointSummary{}
	allEndpoint := i.QueryAllEndpoint()
	endpointItem := model.UnitSummary{Name: "终端", Type: "endpoint", Count: len(allEndpoint)}
	result = append(result, endpointItem)

	return result
}
