package config

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

// Item Item xml结构定义
type Item struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	Catalog     string `xml:"catalog,attr"`
	Content     string `xml:"content,attr"`
}

// Catalogs catalogs xml结构定义
type Catalogs struct {
	Catalog []Item `xml:"catalog"`
}

// Articles articles xml结构定义
type Articles struct {
	Article []Item `xml:"article"`
}

// Content content xml结构定义
type Content struct {
	Catalogs Catalogs `xml:"catalogs"`
	Articles Articles `xml:"articles"`
}

// App app xml结构定义
type App struct {
	Name    string  `xml:"name,attr"`
	Content Content `xml:"content"`
}

// LoadXML 加载xml
func LoadXML(fileName string) (*App, error) {
	app := &App{}

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("open config failed, fileName:%s, err:%s", fileName, err.Error())
		return app, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read config content failed, fileName:%s, err:%s", fileName, err.Error())
		return app, err
	}

	err = xml.Unmarshal(data, app)
	if err != nil {
		log.Printf("unmarshal exception, fileName:%s, err:%s", fileName, err.Error())
		return app, err
	}

	return app, nil
}
