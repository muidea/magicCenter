package config

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Catalog Catalog信息
type Catalog struct {
	Name        string
	Description string
	Catalog     []string
}

// Article Article信息
type Article struct {
	Name        string
	Description string
	Catalog     []string

	Content string
}

// Config config info
type Config struct {
	Catalogs map[string]*Catalog
	Articles map[string]*Article
}

// New 新建Config
func New() *Config {
	cfg := &Config{Catalogs: map[string]*Catalog{}, Articles: map[string]*Article{}}
	return cfg
}

// Load 加载config file
func (s *Config) Load(fileName string) bool {
	val, err := LoadXML(fileName)
	if err != nil {
		log.Printf("load xml exception, err:%s", err.Error())
		return false
	}

	err = s.parse(val)
	if err != nil {
		log.Printf("parse exception, err:%s", err.Error())
		return false
	}

	return true
}

func (s *Config) parse(app *App) error {
	for _, val := range app.Content.Catalogs.Catalog {
		catalog := &Catalog{Name: val.Name, Description: val.Description}
		if len(val.Catalog) > 0 {
			catalog.Catalog = strings.Split(val.Catalog, ",")
		} else {
			catalog.Catalog = []string{}
		}

		_, ok := s.Catalogs[val.ID]
		if ok {
			msg := fmt.Sprintf("[catalog] duplicate catalog, id:%s, name:%s", val.ID, val.Name)
			return errors.New(msg)
		}

		for _, c := range catalog.Catalog {
			_, ok := s.Catalogs[c]
			if !ok {
				msg := fmt.Sprintf("[catalog] no exist parent catalog, name:%s, catalog:%s", val.Name, c)
				return errors.New(msg)
			}
		}

		s.Catalogs[val.Name] = catalog
	}

	for _, val := range app.Content.Articles.Article {
		article := &Article{Name: val.Name, Description: val.Description, Content: val.Content}
		if len(val.Catalog) > 0 {
			article.Catalog = strings.Split(val.Catalog, ",")
		} else {
			article.Catalog = []string{}
		}

		_, ok := s.Articles[val.ID]
		if ok {
			msg := fmt.Sprintf("[article] duplicate article, name:%s", val.Name)
			return errors.New(msg)
		}

		for _, c := range article.Catalog {
			_, ok := s.Catalogs[c]
			if !ok {
				msg := fmt.Sprintf("[article] no exist parent catalog, name:%s, catalog:%s", val.Name, c)
				return errors.New(msg)
			}
		}

		s.Articles[val.ID] = article
	}

	return nil
}
