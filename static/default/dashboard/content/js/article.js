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

article.constructArticleEditView = function(articleList, catalogList) {
    var articleListView = new Array();
    var ii = 0;
    for (var articleIdx = 0; articleIdx < articleList.length; ++articleIdx) {
        var curArticle = articleList[articleIdx];

        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curModule = catalogList[idx];

            if (curArticle.Module == curModule.ID) {
                var view = {
                    ID: curArticle.ID,
                    URL: curArticle.URL,
                    Method: curArticle.Method,
                    Status: curArticle.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                articleListView[ii++] = view;
            }
        }
    }

    return articleListView;
}

article.updateListArticleVM = function(articleList) {
    article.listVM.articles = articleList;
}

article.updateEditModuleVM = function(catalogList) {
    article.editVM.modules = catalogList;
};

article.updateEditArticleVM = function(articleList) {
    article.editVM.articles = articleList;

    // 将已经enable的article设置上checked标示
    for (var offset = 0; offset < article.editVM.articles.length; ++offset) {
        var curArticle = article.editVM.articles[offset];
        if (curArticle.Status > 0) {
            $("#selectArticle-List .article_" + curArticle.ID).prop("checked", true);
        }
    }
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

article.refreshArticleEditView = function(article, catalogList) {};

$(document).ready(function() {
    article.listVM = avalon.define({
        $id: "article-List",
        articles: []
    });

    article.editVM = avalon.define({
        $id: "article-Edit",
        modules: [],
        articles: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        article.updateEditModuleVM(article.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
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
        article.refreshArticleEditView(article.articles, selectModuleArray);
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

                        article.refreshArticleEditView(article.articles, selectModuleArray);
                    })
                });
        }
    );

    article.loadData(function() {
        article.refreshArticleListView(article.articles, article.catalogs);
    })
});