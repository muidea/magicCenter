package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/muidea/magicCenter/tools/config"
	"github.com/muidea/magicCommon/agent"
	common_const "github.com/muidea/magicCommon/common"
	"github.com/muidea/magicCommon/model"
)

var centerServer = "127.0.0.1:8888"
var endpointName = ""
var endpointID = ""
var authToken = ""
var configFile = "config.xml"
var mode = "init"

func main() {
	flag.StringVar(&centerServer, "CenterSvr", centerServer, "magicCenter address")
	flag.StringVar(&endpointName, "EndpointName", endpointName, "endpoint name")
	flag.StringVar(&endpointID, "EndpointID", endpointID, "endpoint id")
	flag.StringVar(&authToken, "AuthToken", authToken, "endpoint authtoken")
	flag.StringVar(&mode, "Mode", mode, "running mode")
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

	rootSumary, ok := agent.FetchSummary(endpointName, model.CATALOG, authToken, sessionID, nil)
	if mode == "repair" && ok {
		repairInitContent(&rootSumary, cfg, authToken, sessionID, agent)
		return
	}

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

	for _, val := range cfg.Catalogs {
		catalogs := findCatalog(val.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, *rootView.CatalogUnit())
		}

		sub, ok := agent.CreateCatalog(val.Name, val.Description, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create catalog failed, name:%s", val.Name)
			return
		}

		catalogInfo[val.ID] = sub
	}

	for _, val := range cfg.Articles {
		catalogs := findCatalog(val.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, *rootView.CatalogUnit())
		}

		sub, ok := agent.CreateArticle(val.Name, val.Content, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create article failed, name:%s", val.Name)
			return
		}

		articleInfo[val.ID] = sub
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

func repairInitContent(rootView *model.SummaryView, cfg *config.Config, authToken, sessionID string, agent agent.Agent) {
	catalogInfo := map[string]model.SummaryView{}
	articleInfo := map[string]model.SummaryView{}

	for _, val := range cfg.Catalogs {
		catalogs := findCatalog(val.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, *rootView.CatalogUnit())
		}

		subCatalog, ok := agent.FetchSummary(val.Name, model.CATALOG, authToken, sessionID, &catalogs[0])
		if ok {
			catalogInfo[val.ID] = subCatalog
			continue
		}

		sub, ok := agent.CreateCatalog(val.Name, val.Description, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create catalog failed, name:%s", val.Name)
			return
		}

		catalogInfo[val.ID] = sub
	}

	for _, val := range cfg.Articles {
		catalogs := findCatalog(val.Catalog, catalogInfo)
		if len(catalogs) == 0 {
			catalogs = append(catalogs, *rootView.CatalogUnit())
		}

		subArticle, ok := agent.FetchSummary(val.Name, model.ARTICLE, authToken, sessionID, &catalogs[0])
		if ok {
			articleInfo[val.ID] = subArticle
			continue
		}

		sub, ok := agent.CreateArticle(val.Name, val.Content, catalogs, authToken, sessionID)
		if !ok {
			log.Printf("create article failed, name:%s", val.Name)
			return
		}

		articleInfo[val.ID] = sub
	}
}
