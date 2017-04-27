var article = {};

article.constructArticleListlView = function(articleList, catalogList) {
    var articleListView = new Array();
    var offset = 0;
    for (var i = 0; i < articleList.length; ++i) {
        var curArticle = articleList[i];
        var catalogNames = "";

        if (curArticle.Catalog) {
            for (var idx = 0; idx < catalogList.length; ++idx) {
                var curCatalog = catalogList[idx];
                for (var j = 0; j < curArticle.Catalog.length; j++) {
                    var val = curArticle.Catalog[j];
                    if (curCatalog.ID == val) {
                        if (catalogNames.length > 0) {
                            catalogNames += ", ";
                        }
                        catalogNames += curCatalog.Name;
                        break;
                    }
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

article.saveArticleAction = function(title, content, catalogs, callBack) {
    $.ajax({
        type: "POST",
        url: "/content/article/",
        data: { "article-title": title, "article-content": content, "article-catalog": catalogs },
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Article);
            }
        }
    });
};

article.updateArticleAtion = function(id, title, content, catalogs, callBack) {
    $.ajax({
        type: "PUT",
        url: "/content/article/" + id + "/",
        data: { "article-title": title, "article-content": content, "article-catalog": catalogs },
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Article);
            }
        }
    });
}

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


    $("#article-Edit .submit").click(
        function() {
            var id = $("#article-Edit .article-id").val();
            var title = $("#article-Edit .article-title").val();
            var content = $("#article-Edit .article-content").val();
            var catalog = "";
            $("#article-Edit .catalog-item:checked").each(
                function() {
                    var id = $(this).val();
                    if (catalog.length > 0) {
                        catalog += ",";
                    }
                    catalog += id;
                });

            var callBack = function(errCode, data) {
                if (errCode > 0) {
                    return;
                }

                $("#article-Edit .article-id").val(-1);
                $("#article-Edit .article-title").val("");
                $("#article-Edit .article-content").val("");
                $("#article-Edit .catalog-item").prop("checked", false);

                article.loadData(function() {
                    article.refreshArticleListView(article.articles, article.catalogs);

                    article.refreshArticleEditView(article.curArticle, article.catalogs);
                })
            };

            if (id == -1) {
                article.saveArticleAction(title, content, catalog, callBack);
            } else {
                article.updateArticleAtion(id, title, content, catalog, callBack);
            }
        }
    );

    article.loadData(function() {
        article.refreshArticleListView(article.articles, article.catalogs);

        article.refreshArticleEditView(article.curArticle, article.catalogs);
    })
});