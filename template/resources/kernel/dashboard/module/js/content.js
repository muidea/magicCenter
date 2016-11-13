var content = {
    moduleList: {},
    articleList: {},
    catalogList: {},
    linkList: {},
    currentModule: {}
};

$(document).ready(function() {

    $("#block-Form .button").click(function() {
        var articleList = "";
        var catalogList = "";
        var linkList = "";
        var ownerModule = $("#block-Form .block-owner").val();
        var block_id = $("#block-Form .block-id").val();
        var blockArray = $("#block-List table tbody tr td :checkbox:checked");
        var articleArray = $("#block-Form .article-List table tbody tr td :checkbox:checked");
        var catalogArray = $("#block-Form .catalog-List table tbody tr td :checkbox:checked");
        var linkArray = $("#block-Form .link-List table tbody tr td :checkbox:checked");

        if (blockArray.length == 0) {
            $("#content-List .alert-Info .content").html("请选择一个功能块");
            $("#content-List .alert-Info").modal();
            return;
        }

        for (var ii = 0; ii < articleArray.length; ++ii) {
            var chk = articleArray[ii];
            articleList += $(chk).attr("value");
            articleList += ",";
        }

        for (var ii = 0; ii < catalogArray.length; ++ii) {
            var chk = catalogArray[ii];
            catalogList += $(chk).attr("value");
            catalogList += ",";
        }

        for (var ii = 0; ii < linkArray.length; ++ii) {
            var chk = linkArray[ii];
            linkList += $(chk).attr("value");
            linkList += ",";
        }

        $.post("/admin/system/ajaxBlockItem/", {
            'module-id': ownerModule,
            "block-id": block_id,
            "article-list": articleList,
            "catalog-list": catalogList,
            "link-list": linkList
        }, function(result) {

            content.currentModule = result.Module;
            if (result.ErrCode > 0) {
                $("#content-List .alert-Info .content").html(result.Reason);
                $("#content-List .alert-Info").modal();
            } else {
                content.refreshBlockContentView();
            }

        }, "json");
    });

    $("#block-Form button.reset").click(function() {

    });
});

content.getModeListView = function() {
    return $("#module-List table");
}

content.getBlockListView = function() {
    return $("#block-List table");
}

content.getArticleListView = function() {
    return $("#block-Form .article-List table");
}

content.getCatalogListView = function() {
    return $("#block-Form .catalog-List table");
}

content.getLinkListView = function() {
    return $("#block-Form .link-List table");
}

content.initialize = function() {
    content.fillModuleView();
};

content.fillModuleView = function() {
    var modueListView = content.getModeListView();

    $(modueListView).find("tbody tr").remove();
    for (var ii = 0; ii < content.moduleList.length; ++ii) {
        var info = content.moduleList[ii];
        var trContent = content.constructModuleItem(info);
        $(modueListView).find("tbody").append(trContent);
    }
};

content.constructModuleItem = function(mod) {
    var tr = document.createElement("tr");
    tr.setAttribute("class", "module");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = mod.Name;
    tr.appendChild(nameTd);

    var descriptionTd = document.createElement("td");
    descriptionTd.innerHTML = mod.Description
    tr.appendChild(descriptionTd);

    var editTd = document.createElement("td");
    var settingButton = document.createElement("input");
    settingButton.setAttribute("type", "button");
    settingButton.setAttribute("onclick", "content.maintainContent('/admin/system/queryModuleContent/?id=" + mod.Id + "'); return false;");
    settingButton.setAttribute("value", "内容管理");
    settingButton.setAttribute("class", "button");
    editTd.appendChild(settingButton);
    tr.appendChild(editTd);
    return tr;
};

content.maintainContent = function(maintainUrl) {
    $.get(maintainUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#module-List .alert-Info .content").html(result.Reason);
            $("#module-List .alert-Info").modal();
            return
        }
        content.currentModule = result.Module;
        content.articleList = result.Articles;
        content.catalogList = result.Catalogs;
        content.linkList = result.Links;
        content.refreshBlockContentView();

        $("#block-Form .block-owner").val(result.Module.Id);

        $("#module-Content .content-header .nav .content-List").find("a").trigger("click");
    }, "json");
};

content.constructBlockItem = function(block) {

    var tr = document.createElement("tr");
    tr.setAttribute("class", "block");

    var nameTd = document.createElement("td");
    var chkBox = document.createElement("input");
    chkBox.setAttribute("type", "checkbox");
    chkBox.setAttribute("value", block.Id);
    nameTd.appendChild(chkBox);
    var label = document.createElement("span");
    label.innerHTML = block.Name;
    nameTd.appendChild(label);
    tr.appendChild(nameTd);

    var numberTd = document.createElement("td");
    var numVal = 0;
    if (block.Article) {
        numVal += block.Article.length;
    }
    if (block.Catalog) {
        numVal += block.Catalog.length;
    }
    if (block.Link) {
        numVal += block.Link.length;
    }
    numberTd.innerHTML = numVal;

    tr.appendChild(numberTd);

    tr.setAttribute("onclick", "content.selectModuleBlock(" + block.Id + "); return false;");

    return tr;
};

content.selectModuleBlock = function(blockId) {
    $("#block-List table>tbody>tr>td>input").prop("checked", false);
    $("#block-Form table>tbody>tr>td>input").prop("checked", false);

    $("#block-List table tbody tr td input").filter("[value=" + blockId + "]").prop("checked", true);
    $("#block-Form .block-id").val(blockId);

    var block = null;
    for (var ii = 0; ii < content.currentModule.Blocks.length; ++ii) {
        var cur = content.currentModule.Blocks[ii];
        if (cur.Id == blockId) {
            block = cur;
            break;
        }
    }

    if (block) {
        content.selectBlockContent(block);
    }
};

content.constructArticleItem = function(article) {

    var tr = document.createElement("tr");

    var nameTd = document.createElement("td");
    var chkBox = document.createElement("input");
    chkBox.setAttribute("type", "checkbox");
    chkBox.setAttribute("value", article.Id);
    chkBox.setAttribute("class", "article");
    nameTd.appendChild(chkBox);
    var label = document.createElement("span");
    label.innerHTML = article.Title;
    nameTd.appendChild(label);
    tr.appendChild(nameTd);

    return tr;
};

content.constructCatalogItem = function(catalog) {

    var tr = document.createElement("tr");

    var nameTd = document.createElement("td");
    var chkBox = document.createElement("input");
    chkBox.setAttribute("type", "checkbox");
    chkBox.setAttribute("value", catalog.Id);
    chkBox.setAttribute("class", "catalog");
    nameTd.appendChild(chkBox);
    var label = document.createElement("span");
    label.innerHTML = catalog.Name;
    nameTd.appendChild(label);
    tr.appendChild(nameTd);

    return tr;
};

content.constructLinkItem = function(link) {

    var tr = document.createElement("tr");

    var nameTd = document.createElement("td");
    var chkBox = document.createElement("input");
    chkBox.setAttribute("type", "checkbox");
    chkBox.setAttribute("value", link.Id);
    chkBox.setAttribute("class", "link");
    nameTd.appendChild(chkBox);
    var label = document.createElement("span");
    label.innerHTML = link.Name;
    nameTd.appendChild(label);
    tr.appendChild(nameTd);

    return tr;
};

content.refreshBlockContentView = function() {
    var blockListView = content.getBlockListView();
    $(blockListView).find("tbody tr").remove();
    if (content.currentModule && content.currentModule.Blocks) {
        for (var ii = 0; ii < content.currentModule.Blocks.length; ++ii) {
            var block = content.currentModule.Blocks[ii];
            var trContent = content.constructBlockItem(block);
            $(blockListView).find("tbody").append(trContent);
        }
    }

    var articleListView = content.getArticleListView();
    $(articleListView).find("tbody tr").remove();
    if (content.articleList) {
        for (var ii = 0; ii < content.articleList.length; ++ii) {
            var article = content.articleList[ii];
            var trContent = content.constructArticleItem(article);
            $(articleListView).find("tbody").append(trContent);
        }
    }

    var catalogListView = content.getCatalogListView();
    $(catalogListView).find("tbody tr").remove();
    if (content.catalogList) {
        for (var ii = 0; ii < content.catalogList.length; ++ii) {
            var catalog = content.catalogList[ii];
            var trContent = content.constructCatalogItem(catalog);
            $(catalogListView).find("tbody").append(trContent);
        }
    }

    var linkListView = content.getLinkListView();
    $(linkListView).find("tbody tr").remove();
    if (content.linkList) {
        for (var ii = 0; ii < content.linkList.length; ++ii) {
            var link = content.linkList[ii];
            var trContent = content.constructLinkItem(link);
            $(linkListView).find("tbody").append(trContent);
        }
    }
};

content.selectBlockContent = function(block) {
    $("#block-Form .block-id").val(block.Id);

    for (var ii = 0; ii < block.Article.length; ++ii) {
        var article = block.Article[ii];
        $("#block-Form .article-List table tbody tr td input").filter("[value=" + article.Rid + "]").prop("checked", true);
    }

    for (var ii = 0; ii < block.Catalog.length; ++ii) {
        var catalog = block.Catalog[ii];
        $("#block-Form .catalog-List table tbody tr td input").filter("[value=" + catalog.Rid + "]").prop("checked", true);
    }

    for (var ii = 0; ii < block.Link.length; ++ii) {
        var link = block.Link[ii];
        $("#block-Form .link-List table tbody tr td input").filter("[value=" + link.Rid + "]").prop("checked", true);
    }
};