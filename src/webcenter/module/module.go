package module

import (
	"webcenter/util/modelhelper"
)


type Module interface {
	ID() string
	Name() string
	Description() string
	
	Enable()
	Disable()
	EnableStatus() int
	
	Default()
	Undefault()
	DefaultStatus() int
	
	Internal() bool
	StyleFlag() int
	
	AssignUrls(urls []string)
	Urls() []string
}

type module struct {
	id string
	name string
	description string
	enableFlag int
	defaultFlag int
	styleFlag int
	blocks []Block
	urls []string
}

func (this *module)ID() string {
	return this.id
}

func (this *module)Name() string {
	return this.name
}

func (this *module)Description() string {
	return this.description
}

func (this *module)Enable() {
	this.enableFlag = 1
}

func (this *module)Disable() {
	this.enableFlag = 0
}

func (this *module)EnableStatus() int {
	return this.enableFlag
}

func (this *module)Default() {
	this.defaultFlag = 1
}

func (this *module)Undefault() {
	this.defaultFlag = 0
}

func (this *module)DefaultStatus() int {
	return this.defaultFlag
}

func (this *module)Internal() bool {
	return this.styleFlag == 0
}

func (this *module)StyleFlag() int {
	return this.styleFlag
}

func (this *module)AssignUrls(urls []string) {
	this.urls = urls
}

func (this *module)Urls() []string {
	return this.urls
}

func newModule(id,name,description string) Module {
	e := &module{}
	e.id = id
	e.name = name
	e.description = description
	
	return e
}


func QueryAllModule() []Module {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	return queryAllModule(helper)
}


