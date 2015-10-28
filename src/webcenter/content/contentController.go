package content

import (
    "webcenter/session"
    "webcenter/common"
	"webcenter/auth"
	"log"
)

type QueryAllContentParam struct {
	session *session.Session
	accessCode string	
}

type QueryAllContentResult struct {
	common.Result
	ArticleInfo []ArticleInfo
	Catalog []Catalog
	Link []Link
	Image []Image
}

type QueryArticleParam struct {
	session *session.Session
	accessCode string
	id int
}

type QueryArticleResult struct {
	common.Result
	Article Article
}

type EditArticleParam struct {
	session *session.Session
	accessCode string
	id int
}

type EditArticleResult struct {
	common.Result
	Article Article
	Catalog []Catalog
}

type DeleteArticleParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteArticleResult struct {
	common.Result
}

type QueryAllArticleParam struct {
	session *session.Session
	accessCode string	
}

type QueryAllArticleResult struct {
	common.Result
	ArticleInfo []ArticleInfo
}

type SubmitArticleParam struct {
	session *session.Session
	accessCode string
	id int
	title string
	content string
	catalog int
	submitDate string	
}

type SubmitArticleResult struct {
	common.Result
}

type QueryCatalogParam struct {
	session *session.Session
	accessCode string
	id int
}

type QueryCatalogResult struct {
	common.Result
	Catalog Catalog
}


type EditCatalogParam struct {
	session *session.Session
	accessCode string
	id int
}

type EditCatalogResult struct {
	common.Result
	Catalog Catalog
	ParentCatalog []Catalog
}

type DeleteCatalogParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteCatalogResult struct {
	common.Result
}

type QueryAllCatalogParam struct {
	session *session.Session
	accessCode string	
}

type QueryAllCatalogResult struct {
	common.Result
	Catalog []Catalog
}

type SubmitCatalogParam struct {
	session *session.Session
	accessCode string
	id int
	name string
	pid int
	submitDate string	
}

type SubmitCatalogResult struct {
	common.Result
}


type QueryLinkParam struct {
	session *session.Session
	accessCode string
	id int
}

type QueryLinkResult struct {
	common.Result
	Link Link
}


type EditLinkParam struct {
	session *session.Session
	accessCode string
	id int
}

type EditLinkResult struct {
	common.Result
	Link Link
}

type DeleteLinkParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteLinkResult struct {
	common.Result
}

type QueryAllLinkParam struct {
	session *session.Session
	accessCode string	
}

type QueryAllLinkResult struct {
	common.Result
	Link []Link
}

type SubmitLinkParam struct {
	session *session.Session
	accessCode string
	id int
	name string
	url string
	logo string
	style int
}

type SubmitLinkResult struct {
	common.Result
}


type QueryAllImageParam struct {
	session *session.Session
	accessCode string	
}

type QueryAllImageResult struct {
	common.Result
	Image []Image
}

type DeleteImageParam struct {
	session *session.Session
	accessCode string
	id int
}

type DeleteImageResult struct {
	common.Result
}


type SubmitImageParam struct {
	session *session.Session
	accessCode string
	id int
	url string
	desc string
}

type SubmitImageResult struct {
	common.Result
}


type contentController struct {
}
 
func (this *contentController)queryAllContentAction(param QueryAllContentParam) QueryAllContentResult {
	result := QueryAllContentResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	result.ArticleInfo = model.GetAllArticleInfo()
	result.Catalog = model.GetAllCatalog()
	result.Link = model.GetAllLink()
	result.Image = model.GetAllImage()
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *contentController)queryAllArticleAction(param QueryAllArticleParam) QueryAllArticleResult {
	result := QueryAllArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	result.ArticleInfo = model.GetAllArticleInfo()
	result.ErrCode = 0

	model.Release()
	
	return result
}

func (this *contentController)queryArticleAction(param QueryArticleParam) QueryArticleResult {
	result := QueryArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	article, found := model.GetArticle(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Article = article
	}
	
	model.Release()
	
	return result
}

func (this *contentController)editArticleAction(param EditArticleParam) EditArticleResult {
	result := EditArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	article, found := model.GetArticle(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.Catalog = model.GetAllCatalog()

		result.ErrCode = 0
		result.Article = article
	}
	
	model.Release()
	
	return result
}

func (this *contentController)deleteArticleAction(param DeleteArticleParam) DeleteArticleResult {
	result := DeleteArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	model.DeleteArticle(param.id)
	result.ErrCode = 0
	
	model.Release()
	
	return result
}
 

func (this *contentController)submitArticleAction(param SubmitArticleParam) SubmitArticleResult {
	result := SubmitArticleResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}

	article := newArticle()
	article.Id = param.id
	article.Title = param.title
	article.Content = param.content
	article.Author.Id = user.Id
	article.Catalog.Id = param.catalog
	article.CreateDate = param.submitDate	
	
	if !model.SaveArticle(article) {
		result.ErrCode = 1
		result.Reason = "保存文章失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存文章成功"
	}
	
	model.Release()

	return result
}

func (this *contentController)queryAllCatalogAction(param QueryAllCatalogParam) QueryAllCatalogResult {
	result := QueryAllCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	result.Catalog = model.GetAllCatalog()
	result.ErrCode = 0
	model.Release()
	
	return result
}
 
func (this *contentController)queryCatalogAction(param QueryCatalogParam) QueryCatalogResult {
	result := QueryCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	catalog, found := model.GetCatalog(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Catalog = catalog
	}

	model.Release()

	return result
}

func (this *contentController)editCatalogAction(param EditCatalogParam) EditCatalogResult {
	result := EditCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	catalog, found := model.GetCatalog(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ParentCatalog = model.GetAvalibleParentCatalog(param.id)
		result.ErrCode = 0
		result.Catalog = catalog
	}

	model.Release()

	return result
}

func (this *contentController)deleteCatalogAction(param DeleteCatalogParam) DeleteCatalogResult {
	result := DeleteCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	articleInfoList := model.GetArticleByCatalog(param.id)
	if (len(articleInfoList) >0) {
		result.ErrCode = 1
		result.Reason = "该分类被引用，无法立即删除"
		return result
	}
	
	model.DeleteCatalog(param.id)
	result.ErrCode = 0
	
	model.Release()
	
	return result
}


func (this *contentController)submitCatalogAction(param SubmitCatalogParam) SubmitCatalogResult {
	result := SubmitCatalogResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create userModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}

	catalog := newCatalog()
	catalog.Id = param.id
	catalog.Name = param.name
	catalog.Pid = param.pid
	catalog.Creater.Id = user.Id
	
	if !model.SaveCatalog(catalog) {
		result.ErrCode = 1
		result.Reason = "保存分类失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存分类成功"
	}
	
	model.Release()

	return result
}


func (this *contentController)queryAllLinkAction(param QueryAllLinkParam) QueryAllLinkResult {
	result := QueryAllLinkResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	result.Link = model.GetAllLink()
	result.ErrCode = 0
	model.Release()
	
	return result
}
 
func (this *contentController)queryLinkAction(param QueryLinkParam) QueryLinkResult {
	result := QueryLinkResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	link, found := model.GetLink(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Link = link
	}

	model.Release()

	return result
}

func (this *contentController)editLinkAction(param EditLinkParam) EditLinkResult {
	result := EditLinkResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	link, found := model.GetLink(param.id)
	if !found {
		result.ErrCode = 1
		result.Reason = "指定对象不存在"
	} else {
		result.ErrCode = 0
		result.Link = link
	}

	model.Release()

	return result
}

func (this *contentController)deleteLinkAction(param DeleteLinkParam) DeleteLinkResult {
	result := DeleteLinkResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	model.DeleteLink(param.id)
	result.ErrCode = 0
	
	model.Release()
	
	return result
}


func (this *contentController)submitLinkAction(param SubmitLinkParam) SubmitLinkResult {
	result := SubmitLinkResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}

	link := newLink()
	link.Id = param.id
	link.Name = param.name
	link.Url = param.url
	link.Logo = param.logo
	link.Style = param.style
	link.Creater.Id = user.Id
	
	if !model.SaveLink(link) {
		result.ErrCode = 1
		result.Reason = "保存链接失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存链接成功"
	}
	
	model.Release()

	return result
}


func (this *contentController)queryAllImageAction(param QueryAllImageParam) QueryAllImageResult {
	result := QueryAllImageResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
	
	result.Image = model.GetAllImage()
	result.ErrCode = 0
	model.Release()
	
	return result
}

func (this *contentController)deleteImageAction(param DeleteImageParam) DeleteImageResult {
	result := DeleteImageResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}
		
	model.DeleteImage(param.id)
	result.ErrCode = 0
	
	model.Release()
	
	return result
}


func (this *contentController)submitImageAction(param SubmitImageParam) SubmitImageResult {
	result := SubmitImageResult{}
	
	model, err := NewModel()
	if err != nil {
		log.Print("create contentModel failed")
		
		result.ErrCode = 1
		result.Reason = "创建Model失败"
		return result
	}
	defer model.Release()

	session := param.session
	account, found := session.GetAccount()
	if !found {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result
	}
	
	user, _ := auth.QueryUserByAccount(account, model.dao)
	if !user.IsAdmin() {
		log.Print("illegal authorization")
		
		result.ErrCode = 1
		result.Reason = "权限不足"
		return result		
	}

	image := newImage()
	image.Id = param.id
	image.Url = param.url
	image.Desc = param.desc
	image.Creater.Id = user.Id
	
	if !model.SaveImage(image) {
		result.ErrCode = 1
		result.Reason = "保存链接失败"
	} else {
		result.ErrCode = 0
		result.Reason = "保存链接成功"
	}
	
	model.Release()

	return result
}


