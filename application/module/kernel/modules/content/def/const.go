package def

import "muidea.com/magicCenter/application/common"

// ID 模块ID
const ID = common.CotentModuleID

// Name 模块名称
const Name = "Magic Content"

// Description 模块描述信息
const Description = "Magic 内容管理模块"

// URL 模块Url
const URL = "/content"

// GetArticleDetail 查询指定文章
const GetArticleDetail = "/article/:id"

// PostArticle 新建文章
const PostArticle = "/article/"

// PutArticle 更新文章
const PutArticle = "/article/:id"

// DeleteArticle 删除文章
const DeleteArticle = "/article/:id"

// GetArticleList 查询文章列表
const GetArticleList = "/articles/"

// GetCatalogDetail 查询指定分类
const GetCatalogDetail = "/catalog/:id"

// PostCatalog 新建分类
const PostCatalog = "/catalog/"

// PutCatalog 更新分类
const PutCatalog = "/catalog/:id"

// DeleteCatalog 删除分类
const DeleteCatalog = "/catalog/:id"

// QueryCatalogByName 查询指定分类
const QueryCatalogByName = "/catalog/"

// GetCatalogList 查询分类列表
const GetCatalogList = "/catalogs/"

// GetLinkDetail 查询指定链接
const GetLinkDetail = "/link/:id"

// PostLink 新建链接
const PostLink = "/link/"

// PutLink 更新链接
const PutLink = "/link/:id"

// DeleteLink 删除链接
const DeleteLink = "/link/:id"

// GetLinkList 查询链接列表
const GetLinkList = "/links/"

// GetMediaDetail 查询指定文件
const GetMediaDetail = "/media/:id"

// PostMedia 新建文件
const PostMedia = "/media/"

// PutMedia 更新文件
const PutMedia = "/media/:id"

// DeleteMedia 删除文件
const DeleteMedia = "/media/:id"

// GetMediaList 查询文件列表
const GetMediaList = "/medias/"
