package module

import (
	"webcenter/util/modelhelper"
)

type Page interface {
	Url() string
	Blocks() []Block
}

type page struct {
	url string
	blocks []Block
}

func (this *page)Url() string {
	return this.url
}

func (this *page)Blocks() []Block {
	return this.blocks
}

func AddPageBlock(url string, block int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()

	return addPageBlock(helper, url, block)	
}

func RemovePageBlock(url string, block int) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()

	removePageBlock(helper, url, block)
}

func QueryPage(url string) Page {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	p := &page{}
	p.url = url
	p.blocks = queryPageBlock(helper, url)
		
	return p
}
