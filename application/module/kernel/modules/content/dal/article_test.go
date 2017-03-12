package dal

import (
	"log"
	"testing"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

func TestArticle(t *testing.T) {
	log.Print("------------------TestArticle--------------------")
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ar := model.Article{}
	ar.Title = "testing"
	ar.Content = "123456789"
	ar.Author = 0
	ar.CreateDate = "2016-08-08 00:00:00"
	ar.Catalog = append(ar.Catalog, 8)

	summary, ret := SaveArticle(helper, ar)
	if !ret {
		t.Error("SaveArticle failed")
		return
	}

	arInfo, found := QueryArticleByID(helper, summary.ID)
	if !found {
		t.Error("QueryArticleByID failed")
		return
	}

	if arInfo.CreateDate != "2016-08-08 00:00:00" {
		t.Error("QueryArticleByID failed, invalid createDate")
		return
	}

	arInfo.Content = "0987654321"
	summary, ret = SaveArticle(helper, arInfo)
	if !ret {
		t.Error("SaveArticle failed")
		return
	}

	arSummarys := QueryAllArticleSummary(helper)
	if len(arSummarys) < 1 {
		t.Error("QueryAllArticleSummary failed")
	}

	ret = DeleteArticle(helper, arInfo.ID)
	if !ret {
		t.Error("DeleteArticle failed")
		return
	}

}
