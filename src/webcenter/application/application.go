package application

import (
	"os"
	"log"
	"martini"
	"webcenter/module"
)

type Application interface {
	Run()
}

type application struct {
	uploadPath string
	martiniRouter martini.Router
	martiniFrame *martini.Martini
	martinInstance *martini.ClassicMartini
}

var staticPath string = "static"
var resourceFilePath string = "template"
var uploadFilePath string = "upload"
var serverPort string = "8888"

var systemName string
var systemLogo string
var systemDomain string
var systemMailServer string
var systemMailAccount string
var systemMailPassword string

var router martini.Router = martini.NewRouter()
var martiniFrame *martini.Martini = martini.New()
var app *application = nil

func AppInstance() Application {
	if (app == nil) {
		app = &application{}
		
		app.construct()
	}
	
	return app 
}

func Name() string {
	return systemName
}

func UpdateName(name string) {
	systemName = name
}

func Logo() string {
	return systemLogo
}

func UpdateLogo(logo string) {
	systemLogo = logo
}

func Domain() string {
	return systemDomain
}

func UpdateDomain(domain string) {
	systemDomain = domain
}

func MailServer() string {
	return systemMailServer
}

func UpdateMailServer(server string) {
	systemMailServer = server
}

func MailAccount() string {
	return systemMailAccount
}

func UpdateMailAccount(account string) {
	systemMailAccount = account
}

func MailPassword() string {
	return systemMailPassword
}

func UpdateMailPassword(password string) {
	systemMailPassword = password
}

func UploadPath() string {
	return uploadFilePath
}

func ResourcePath() string {
	return resourceFilePath
}

func StaticPath() string {
	return staticPath
}

func RegisterGetHandler(pattern string, h interface{}) {
	router.Get(pattern, h)	
}

func UnRegisterGetHandler(pattern string, h interface{}) {
}

func RegisterPostHandler(pattern string, h interface{}) {
	router.Post(pattern, h)
}

func UnRegisterPostHandler(pattern string, h interface{}) {
}

func BindStatic(path string) {
	martiniFrame.Use(martini.Static(path))
}

func (instance application) construct() {
	log.Print("application construct")
	
	app.uploadPath = uploadFilePath
	app.martiniRouter = router
	app.martiniFrame = martiniFrame
	
	os.Setenv("PORT", serverPort)
}

func (instance application) Run() {
	log.Print("application Run")
	
	module.StarupAllModules()
	
	instance.martiniFrame.Use(martini.Logger())
	instance.martiniFrame.Use(martini.Recovery())
	instance.martiniFrame.MapTo(instance.martiniRouter, (*martini.Routes)(nil))
	instance.martiniFrame.Action(instance.martiniRouter.Handle)
	instance.martinInstance = &martini.ClassicMartini{instance.martiniFrame, instance.martiniRouter}
	instance.martinInstance.Run()
	
	module.CleanupAllModules()	
}


