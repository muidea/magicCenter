var article = {
    articleInfo: {},
    catalogInfo: {},
    userInfo: {}
};

$(document).ready(function() {
    // 绑定表单提交事件处理器
    $("#article-Content .article-Form").submit(function() {
        var options = {
            beforeSubmit: showRequest,
            success: showResponse,
            dataType: "json"
        };

        function showRequest() {}

        function showResponse(result) {

            if (result.ErrCode > 0) {
                $("#article-Edit .alert-Info .content").html(result.Reason);
                $("#article-Edit .alert-Info").modal();
            } else {
                article.refreshArticle();
            }
        }

        function validate() {
            var result = true

            $("#article-Content .article-Form .article-title").parent().find("span").remove();
            var title = $("#article-Content .article-Form .article-title").val();
            if (title.length == 0) {
                $("#article-Content .article-Form .article-title").parent().append("<span class=\"input-notification error\">请输入标题</span>");
                result = false;
            }

            $("#article-Content .article-Form .article-content").parent().find("span").remove();
            var content = $("#article-Content .article-Form .article-content").val();
            if (content.length == 0) {
                $("#article-Content .article-Form .article-content").parent().append("<span class=\"input-notification error\">请输入内容</span>");
                result = false;
            }

            return result;
        }

        if (!validate()) {
            return false;
        }

        //提交表单
        $(this).ajaxSubmit(options);

        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });

    $("#article-Content .article-Form button.reset").click(function() {
        $("#article-Edit .article-Form .article-id").val(-1);
        $("#article-Edit .article-Form .article-title").val("");
        //$("#article-Edit .article-Form .article-content").wysiwyg("setContent", "");
        $("#article-Edit .article-Form .article-content").val("");
        $("#article-Edit .article-Form .article-catalog input").prop("checked", false);
    });
});

article.initialize = function() {
    article.refreshCatalog();
    article.fillArticleView();
};

article.refreshCatalog = function() {
    $("#article-Edit .article-Form .article-catalog").children().remove();
    for (var ii = 0; ii < article.catalogInfo.length; ++ii) {
        var catalog = article.catalogInfo[ii];
        $("#article-Edit .article-Form .article-catalog").append("<label><input type='checkbox' name='article-catalog' value=" + catalog.ID + "> </input>" + catalog.Name + "</label> ");
    }
};

article.refreshArticle = function() {
    $.get("/content/queryAllArticleSummary/", {}, function(result) {
        article.articleInfo = result.Articles;

        article.fillArticleView();
    }, "json");
};

article.fillArticleView = function() {
    $("#article-List table tbody tr").remove();
    var articleInfoList = article.articleInfo;
    for (var ii = 0; ii < articleInfoList.length; ++ii) {
        var articleInfo = articleInfoList[ii];
        var trContent = article.constructArticleItem(articleInfo);
        $("#article-List table tbody").append(trContent);
    }

    $("#article-Edit .article-Form .article-id").val(-1);
    $("#article-Edit .article-Form .article-title").val("");
    //$("#article-Edit .article-Form .article-content").wysiwyg("setContent", "");
    $("#article-Edit .article-Form .article-content").val("");
    $("#article-Edit .article-Form .article-catalog input").prop("checked", false);

};

article.constructArticleItem = function(articleInfo) {
    var tr = document.createElement("tr");
    tr.setAttribute("class", "article");

    var titleTd = document.createElement("td");
    var titleLink = document.createElement("a");
    titleLink.setAttribute("class", "edit");
    titleLink.setAttribute("href", "#queryArticle");
    titleLink.setAttribute("onclick", "article.editArticle('/content/queryArticle/?id=" + articleInfo.ID + "'); return false;");
    titleLink.innerHTML = articleInfo.Title;
    titleTd.appendChild(titleLink);
    tr.appendChild(titleTd);

    var cataLogTd = document.createElement("td");
    var catalogs = "";
    if (articleInfo.Catalog) {
        for (var ii = 0; ii < articleInfo.Catalog.length;) {
            var cid = articleInfo.Catalog[ii++];
            for (var jj = 0; jj < article.catalogInfo.length;) {
                var catalog = article.catalogInfo[jj++];
                if (catalog.ID == cid) {
                    catalogs += catalog.Name;
                    if (ii < articleInfo.Catalog.length) {
                        catalogs += ",";
                    }
                    break;
                }
            }
        }
    }
    catalogs = catalogs.length == 0 ? '-' : catalogs;
    cataLogTd.innerHTML = catalogs;
    tr.appendChild(cataLogTd);

    var authorTd = document.createElement("td");
    var authorValue = "-";
    for (var ii = 0; ii < article.userInfo.length; ii++) {
        var author = article.userInfo[ii];
        if (author.ID == articleInfo.Author) {
            authorValue = author.Name;
            break;
        }
    }
    authorTd.innerHTML = authorValue;
    tr.appendChild(authorTd);

    var createDateTd = document.createElement("td");
    createDateTd.innerHTML = articleInfo.CreateDate;
    tr.appendChild(createDateTd);

    var editTd = document.createElement("td");
    var editLink = document.createElement("a");
    editLink.setAttribute("class", "edit");
    editLink.setAttribute("href", "#editArticle");
    editLink.setAttribute("onclick", "article.editArticle('/content/queryArticle/?id=" + articleInfo.ID + "'); return false");
    var editImage = document.createElement("img");
    editImage.setAttribute("src", "/resources/admin/images/pencil.png");
    editImage.setAttribute("alt", "Edit");
    editLink.appendChild(editImage);
    editTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteArticle");
    deleteLink.setAttribute("onclick", "article.deleteArticle('/content/deleteArticle/?id=" + articleInfo.ID + "'); return false;");
    var deleteImage = document.createElement("img");
    deleteImage.setAttribute("src", "/resources/admin/images/cross.png");
    deleteImage.setAttribute("alt", "Delete");
    deleteLink.appendChild(deleteImage);
    editTd.appendChild(deleteLink);

    tr.appendChild(editTd);

    return tr;
};

article.editArticle = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#article-List .alert-Info .content").html(result.Reason);
            $("#article-List .alert-Info").modal();
            return
        }

        $("#article-Edit .article-Form .article-id").val(result.Article.ID);
        $("#article-Edit .article-Form .article-title").val(result.Article.Title);
        //$("#article-Edit .article-Form .article-content").wysiwyg("setContent", result.Article.Content);
        $("#article-Edit .article-Form .article-content").val(result.Article.Content);
        $("#article-Edit .article-Form .article-catalog input").prop("checked", false);

        if (result.Article.Catalog) {
            for (var ii = 0; ii < result.Article.Catalog.length; ++ii) {
                var ca = result.Article.Catalog[ii];
                $("#article-Edit .article-Form .article-catalog input").filter("[value=" + ca + "]").prop("checked", true);
            }
        }

        $("#article-Content .content-header .nav .article-Edit").find("a").trigger("click");
    }, "json");
};

article.deleteArticle = function(deleteUrl) {
    $.get(deleteUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#article-List .alert-Info .content").html(result.Reason);
            $("#article-List .alert-Info").modal();
            return
        }

        article.refreshArticle();
    }, "json");
};