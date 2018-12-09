package handler

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/content/dal"
	common_util "muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

type commentActionHandler struct {
}

func (i *commentActionHandler) findCommentByCatalog(catalog model.CatalogUnit, pageFilter *common_util.PageFilter) ([]model.CommentDetail, int) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryCommentByCatalog(dbhelper, catalog, pageFilter)
}

func (i *commentActionHandler) createComment(subject, content, createDate string, catalog []model.CatalogUnit, author int) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.CreateComment(dbhelper, subject, content, createDate, author, catalog)
}

func (i *commentActionHandler) saveComment(comment model.CommentDetail) (model.Summary, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.SaveComment(dbhelper, comment)
}

func (i *commentActionHandler) disableComment(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DisableCommentByID(dbhelper, id)
}

func (i *commentActionHandler) destroyComment(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteCommentByID(dbhelper, id)
}
