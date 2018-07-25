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

	_, ok = agent.FetchSummary(endpointName, model.CATALOG, authToken, sessionID)
	if ok {
		return
	}

	catalog := model.Catalog{ID: common_const.SystemContentCatalog.ID, Name: common_const.SystemContentCatalog.Name}
	rootView, ok := agent.CreateCatalog(endpointName, "auto setup catalog description", []model.Catalog{catalog}, authToken, sessionID)
	if !ok {
		log.Printf("create root catalog failed")
		return
	}
	catalogInfo := map[string]model.SummaryView{}
	articleInfo := map[string]model.SummaryView{}
	catalogInfo[endpointName] = rootView
	for _, c := range cfg.Catalogs {
		catalogs := findCatalog(c.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, model.Catalog{ID: rootView.ID, Name: rootView.Name})
		}

		sub, ok := agent.CreateCatalog(c.Name, c.Description, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create catalog failed, name:%s", c.Name)
			return
		}
		catalogInfo[c.Name] = sub
	}

	for _, c := range cfg.Articles {
		catalogs := findCatalog(c.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, model.Catalog{ID: rootView.ID, Name: rootView.Name})
		}

		sub, ok := agent.CreateArticle(c.Name, c.Description, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create article failed, name:%s", c.Name)
			return
		}

		articleInfo[c.Name] = sub
	}
}

func findCatalog(names []string, catalogInfo map[string]model.SummaryView) []model.Catalog {
	rets := []model.Catalog{}

	for _, v := range names {
		catalog, ok := catalogInfo[v]
		if ok {
			rets = append(rets, model.Catalog{ID: catalog.ID, Name: catalog.Name})
		}
	}

	return rets
}
