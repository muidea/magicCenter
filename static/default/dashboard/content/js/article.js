var article = {};

article.constructArticleListlView = function(articleList, catalogList) {
    var articleListView = new Array();
    var offset = 0;
    for (var i = 0; i < articleList.length; ++i) {
        var curArticle = articleList[i];
        var catalogNames = "";
        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curCatalog = catalogList[idx];
            for (var j = 0; j < curArticle.Catalog; j++) {
                var val = curArticle.Catalog[j];
                if (curCatalog.ID == val) {
                    if (catalogNames.length > 0) {
                        catalogNames += ", ";
                    }
                    catalogNames += curCatalog.Name;
                }
            }
        }
        var view = {
            ID: curArticle.ID,
            Name: curArticle.Name,
            Catalog: catalogNames,
            CreateDate: curArticle.CreateDate
        };

        articleListView[offset++] = view;
    }

    return articleListView;
};


article.constructArticleEditView = function(catalogList) {
    var catalogListView = new Array();
    var ii = 0;
    for (var idx = 0; idx < catalogList.length; ++idx) {
        var curCatalog = catalogList[idx];

        var view = {
            ID: curCatalog.ID,
            Name: curCatalog.Name
        };

        catalogListView[ii++] = view;
    }

    return catalogListView;
};

article.updateListArticleVM = function(articleList) {
    article.listVM.articles = articleList;
};

article.updateEditArticleVM = function(curArticle, catalogList) {
    article.editVM.article = curArticle;
    article.editVM.catalogs = catalogList;
};

// 加载全部的Article
article.getAllArticlesAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/content/article/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Article);
            }
        }
    });
};

// 加载全部Catalog
article.getAllCatalogsAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/content/catalog/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Catalog);
            }
        }
    });
};

article.loadData = function(callBack) {
    var getAllCatalogsCallBack = function(errCode, catalogList) {
        if (errCode != 0) {
            return;
        }

        article.catalogs = catalogList;
        if (callBack != null) {
            callBack()
        }
    };

    var getAllArticlesCallBack = function(errCode, articleList) {
        if (errCode != 0) {
            return;
        }
        article.curArticle = { ID: -1, Name: "", Content: "", Catalog: [] };
        article.articles = articleList;
        article.getAllCatalogsAction(getAllCatalogsCallBack);
    };

    // 加载完成
    article.getAllArticlesAction(getAllArticlesCallBack);
}

article.refreshArticleListView = function(articleList, catalogList) {
    var articlesView = article.constructArticleListlView(articleList, catalogList);
    article.updateListArticleVM(articlesView);
};

article.refreshArticleEditView = function(curArticle, catalogList) {
    var articleView = {
        ID: curArticle.ID,
        Name: curArticle.Name,
        Content: curArticle.Content,
        Catalog: curArticle.Catalog
    };
    console.log(curArticle);

    var catalogView = article.constructArticleEditView(catalogList);
    article.updateEditArticleVM(articleView, catalogView);
};

$(document).ready(function() {
    article.listVM = avalon.define({
        $id: "article-List",
        articles: []
    });

    article.editVM = avalon.define({
        $id: "article-Edit",
        article: {},
        catalogs: []
    });


    $("#selectArticle-button").click(
        function() {
            var selectArticleList = new Array();
            var offset = 0;
            $("#selectArticle-List .article_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectArticleList[offset++] = id;
                }
            );

            var unSelectArticleList = new Array();
            offset = 0;
            $("#selectArticle-List .article_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectArticleList[offset++] = id;
                }
            );

            article.statusArticlesAction(
                selectArticleList,
                unSelectArticleList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < article.modules.length; idx++) {
                                var curModule = article.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    article.loadData(function() {
                        article.refreshArticleListView(article.articles, article.catalogs);

                        article.refreshArticleEditView(article.curArticle, article.catalogs);
                    })
                });
        }
    );

    article.loadData(function() {
        article.refreshArticleListView(article.articles, article.catalogs);

        article.refreshArticleEditView(article.curArticle, article.catalogs);
    })
});