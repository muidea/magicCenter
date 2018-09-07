package main

import (
	"flag"
	"fmt"
	"log"

	"muidea.com/magicCenter/tools/config"
	"muidea.com/magicCommon/agent"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

var centerServer = "127.0.0.1:8888"
var endpointName = ""
var endpointID = ""
var authToken = ""
var configFile = "config.xml"

func main() {
	flag.StringVar(&centerServer, "CenterSvr", centerServer, "magicCenter address")
	flag.StringVar(&endpointName, "EndpointName", endpointName, "endpoint name")
	flag.StringVar(&endpointID, "EndpointID", endpointID, "endpoint id")
	flag.StringVar(&authToken, "AuthToken", authToken, "endpoint authtoken")
	flag.Parse()

	if len(centerServer) == 0 || len(endpointID) == 0 || len(endpointName) == 0 || len(authToken) == 0 {
		fmt.Printf("---------------magicSetupTool---------------\n")
		flag.PrintDefaults()
		return
	}

	cfg := config.New()
	ok := cfg.Load(configFile)
	if !ok {
		return
	}

	agent := agent.New()
	authToken, sessionID, ok := agent.Start(centerServer, endpointID, authToken)
	if !ok {
		log.Printf("illegal runtime param, centerServer:%s, endpointID:%s, authToken:%s", centerServer, endpointID, authToken)
		return
	}
	defer agent.Stop()

	_, ok = agent.FetchSummary(endpointName, model.CATALOG, authToken, sessionID, nil)
	if ok {
		return
	}

	catalog := common_const.SystemContentCatalog.CatalogUnit()
	rootView, ok := agent.CreateCatalog(endpointName, "auto setup catalog description", []model.CatalogUnit{*catalog}, authToken, sessionID)
	if !ok {
		log.Printf("create root catalog failed")
		return
	}
	catalogInfo := map[string]model.SummaryView{}
	articleInfo := map[string]model.SummaryView{}

	for key, val := range cfg.Catalogs {
		catalogs := findCatalog(val.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, *rootView.CatalogUnit())
		}

		sub, ok := agent.CreateCatalog(val.Name, val.Description, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create catalog failed, name:%s", val.Name)
			return
		}

		catalogInfo[key] = sub
	}

	for key, val := range cfg.Articles {
		catalogs := findCatalog(val.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, *rootView.CatalogUnit())
		}

		sub, ok := agent.CreateArticle(val.Name, val.Content, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create article failed, name:%s", val.Name)
			return
		}

		articleInfo[key] = sub
	}
}

func findCatalog(ids []string, catalogInfo map[string]model.SummaryView) []model.CatalogUnit {
	rets := []model.CatalogUnit{}

	for _, v := range ids {
		catalog, ok := catalogInfo[v]
		if ok {
			rets = append(rets, *catalog.CatalogUnit())
		}
	}

	return rets
}
