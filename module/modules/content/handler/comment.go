package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	"muidea.com/magicCommon/model"
)

type commentActionHandler struct {
	dbhelper dbhelper.DBHelper
}

func (i *commentActionHandler) findCommentByCatalog(catalog model.CatalogUnit) []model.CommentDetail {
	return dal.QueryCommentByCatalog(i.dbhelper, catalog)
}

func (i *commentActionHandler) createComment(subject, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	return dal.CreateComment(i.dbhelper, subject, content, createDate, author, catalog)
}

func (i *commentActionHandler) saveComment(comment model.CommentDetail) (model.Summary, bool) {
	return dal.SaveComment(i.dbhelper, comment)
}

func (i *commentActionHandler) disableComment(id int) bool {
	return dal.DisableCommentByID(i.dbhelper, id)
}

func (i *commentActionHandler) destroyComment(id int) bool {
	return dal.DeleteCommentByID(i.dbhelper, id)
}
