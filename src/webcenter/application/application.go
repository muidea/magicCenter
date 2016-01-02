package application

import (
	"os"
	"log"
	"martini"
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
var webDomain string = "127.0.0.1:8888"
var webMailServer string = "smtp.126.com:25"
var webMailAccount string = "rangh@126.com"
var webMailPassword string = "hRangh@13924"

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

func Domain() string {
	return webDomain
}

func MailServer() string {
	return webMailServer
}

func MailAccount() string {
	return webMailAccount
}

func MailPassword() string {
	return webMailPassword
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

func RegisterPostHandler(pattern string, h interface{}) {
	router.Post(pattern, h)
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
	
	instance.martiniFrame.Use(martini.Logger())
	instance.martiniFrame.Use(martini.Recovery())
	instance.martiniFrame.MapTo(instance.martiniRouter, (*martini.Routes)(nil))
	instance.martiniFrame.Action(instance.martiniRouter.Handle)
	instance.martinInstance = &martini.ClassicMartini{instance.martiniFrame, instance.martiniRouter}
	instance.martinInstance.Run()
}


