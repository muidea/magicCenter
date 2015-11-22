package application

import (
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

var g_staticPath string = "static"
var g_resourcePath string = "template"
var g_uploadPath string = "upload"
var g_router martini.Router = martini.NewRouter()
var g_martiniFrame *martini.Martini = martini.New()
var g_app *application = nil

func AppInstance() Application {
	if (g_app == nil) {
		g_app = &application{}
		
		g_app.construct()
	}
	
	return g_app 
}

func UploadPath() string {
	return g_uploadPath
}

func ResourcePath() string {
	return g_resourcePath
}

func StaticPath() string {
	return g_staticPath
}

func RegisterGetHandler(pattern string, h interface{}) {
	g_router.Get(pattern, h)	
}

func RegisterPostHandler(pattern string, h interface{}) {
	g_router.Post(pattern, h)
}

func BindStatic(path string) {
	g_martiniFrame.Use(martini.Static(path))
}

func (app application) construct() {
	log.Print("application construct")
	
	g_app.uploadPath = g_uploadPath
	g_app.martiniRouter = g_router
	g_app.martiniFrame = g_martiniFrame
}

func (app application) Run() {
	log.Print("application Run")
	
	app.martiniFrame.Use(martini.Logger())
	app.martiniFrame.Use(martini.Recovery())
	app.martiniFrame.MapTo(app.martiniRouter, (*martini.Routes)(nil))
	app.martiniFrame.Action(app.martiniRouter.Handle)
	app.martinInstance = &martini.ClassicMartini{app.martiniFrame, app.martiniRouter}
	app.martinInstance.Run()
}


