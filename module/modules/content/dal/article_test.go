package dal

import (
	"log"
	"testing"

	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCommon/model"
)

func TestArticle(t *testing.T) {
	log.Print("------------------TestArticle--------------------")
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ar := model.ArticleDetail{}
	ar.Name = "testing"
	ar.Content = "123456789"
	ar.Creater = 0
	ar.CreateDate = "2016-08-08 00:00:00"
	ar.Catalog = append(ar.Catalog, 8)
	ar.ID = 10

	summary, ret := CreateArticle(helper, ar.Name, ar.Content, ar.Catalog, ar.Creater, ar.CreateDate)
	if !ret {
		t.Error("CreateArticle failed")
		return
	}

	arDetail, found := QueryArticleByID(helper, summary.ID)
	if !found {
		t.Error("QueryArticleByID failed")
		return
	}

	if arDetail.CreateDate != "2016-08-08 00:00:00" {
		t.Error("QueryArticleByID failed, invalid createDate")
		return
	}

	arDetail.Content = "0987654321"
	summary, ret = SaveArticle(helper, arDetail)
	if !ret {
		t.Error("SaveArticle failed")
		return
	}

	arSummarys := QueryAllArticleSummary(helper)
	if len(arSummarys) < 1 {
		t.Error("QueryAllArticleSummary failed")
	}

	ret = DeleteArticle(helper, arDetail.ID)
	if !ret {
		t.Error("DeleteArticle failed")
		return
	}

}
