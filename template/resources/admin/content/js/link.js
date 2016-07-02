var link = {
    errCode: 0,
    reason: '',
    linkInfo: {},
    catalogInfo: {}
};


link.initialize = function() {

    link.refreshCatalog();

    link.fillLinkView();

    // 绑定表单提交事件处理器
    $('#link-content .link-Form').submit(function() {
        var options = {
            beforeSubmit: showRequest, // pre-submit callback
            success: showResponse, // post-submit callback
            dataType: 'json' // 'xml', 'script', or 'json' (expected server response type) 
        };

        // pre-submit callback
        function showRequest() {
            //return false;
        }
        // post-submit callback
        function showResponse(result) {
            $("#link-Edit div.notification").hide();

            if (result.ErrCode > 0) {
                $("#link-Edit div.error div").html(result.Reason);
                $("#link-Edit div.error").show();
            } else {
                $("#link-Edit div.success div").html(result.Reason);
                $("#link-Edit div.success").show();

                link.refreshLink();
            }
        }

        function validate() {
            var result = true

            $("#link-content .link-Form .link-name").parent().find("span").remove();
            var name = $("#link-content .link-Form .link-name").val();
            if (name.length == 0) {
                $("#link-content .link-Form .link-name").parent().append("<span class=\"input-notification error png_bg\">请输入站点名</span>");
                result = false;
            }

            var url = $("#link-content .link-Form .link-url").val();
            if (url.length == 0) {
                $("#link-content .link-Form .link-url").parent().append("<span class=\"input-notification error png_bg\">请输入网址</span>");
                result = false;
            }

            var logo = $("#link-content .link-Form .link-logo").val();
            if (logo.length == 0) {
                $("#link-content .link-Form .link-name").parent().append("<span class=\"input-notification error png_bg\">请输入Log地址</span>");
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
};

link.refreshCatalog = function() {
    $("#link-Edit .link-Form .link-catalog").children().remove();
    for (var ii = 0; ii < link.catalogInfo.length; ++ii) {
        var catalog = link.catalogInfo[ii];
        $("#link-Edit .link-Form .link-catalog").append("<label><input type='checkbox' name='link-catalog' value=" + catalog.Id+ "> </input>" + catalog.Name + "</label> ");
    }
};


link.refreshLink = function() {
    $.get("/admin/content/queryAllLink/", {}, function(result) {
        article.errCode = result.ErrCode;
        article.reason = result.Reason;

        link.linkInfo = result.Links;

        link.fillLinkView();
    }, "json");
};


link.fillLinkView = function() {

    $("#link-List div.notification").hide();

    if (link.errCode > 0) {
        $("#link-List div.error div").html(link.reason);
        $("#link-List div.error").show();

        $("#link-List table").hide();
        return;
    }

    $("#link-List table tbody tr").remove();
    for (var ii = 0; ii < link.linkInfo.length; ++ii) {
        var info = link.linkInfo[ii];
        var trContent = link.constructLinkItem(info);
        $("#link-List table tbody").append(trContent);
    }
    $("#link-List table tbody tr:even").addClass("alt-row");
    $("#link-List table").show();

    $("#link-Edit div.notification").hide();
    $("#link-Edit .link-Form .link-id").val(-1);
    $("#link-Edit .link-Form .link-name").val("");
    $("#link-Edit .link-Form .link-url").val("");
    $("#link-Edit .link-Form .link-logo").val("");

    $("#link-Edit .link-Form .link-catalog").children().remove();
    for (var ii = 0; ii < link.catalogInfo.length; ++ii) {
        var ca = link.catalogInfo[ii];
        $("#link-Edit .link-Form .link-catalog").append("<label><input type='checkbox' name='link-catalog' value=" + ca.Id+ "> </input>" + ca.Name + "</label> ");
    }
    if (ii == 0) {
        $("#link-Edit .link-Form .link-catalog").append("<label><input type='checkbox' name='link-catalog' readonly='readonly' value='-1' onclick='return false'> </input>-</label> ");
    }
};


link.constructLinkItem = function(lnk) {
    var tr = document.createElement("tr");
    tr.setAttribute("class", "link");

    var checkBoxTd = document.createElement("td");
    var checkBox = document.createElement("input");
    checkBox.setAttribute("type", "checkbox");

    checkBoxTd.appendChild(checkBox);
    tr.appendChild(checkBoxTd);

    var nameTd = document.createElement("td");
    var nameLink = document.createElement("a");
    nameLink.setAttribute("class", "edit");
    nameLink.setAttribute("href", "#editLink");
    nameLink.setAttribute("onclick", "link.editLink('/admin/content/editLink/?id=" + lnk.Id + "'); return false;");
    nameLink.innerHTML = lnk.Name;
    nameTd.appendChild(nameLink);
    tr.appendChild(nameTd);

    var urlTd = document.createElement("td");
    urlTd.innerHTML = lnk.Url;
    tr.appendChild(urlTd);

    var catalogTd = document.createElement("td");
    var catalogs = "";
    if (lnk.Catalog) {
        for (var ii = 0; ii < lnk.Catalog.length;) {
            catalogs += lnk.Catalog[ii++].Name;
            if (ii < lnk.Catalog.length) {
                catalogs += ","
            } else {
                break;
            }
        }
    }
    catalogs = catalogs.length == 0 ? '-' : catalogs;
    catalogTd.innerHTML = catalogs;
    tr.appendChild(catalogTd);

    var editTd = document.createElement("td");
    var editLink = document.createElement("a");
    editLink.setAttribute("class", "edit");
    editLink.setAttribute("href", "#editLink");
    editLink.setAttribute("onclick", "link.editLink('/admin/content/editLink/?id=" + lnk.Id + "'); return false;");
    var editUrl = document.createElement("img");
    editUrl.setAttribute("src", "/resources/admin/images/pencil.png");
    editUrl.setAttribute("alt", "Edit");
    editLink.appendChild(editUrl);
    editTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteLink");
    deleteLink.setAttribute("onclick", "link.deleteLink('/admin/content/deleteLink/?id=" + lnk.Id + "'); return false;");
    var deleteUrl = document.createElement("img");
    deleteUrl.setAttribute("src", "/resources/admin/images/cross.png");
    deleteUrl.setAttribute("alt", "Delete");
    deleteLink.appendChild(deleteUrl);
    editTd.appendChild(deleteLink);

    tr.appendChild(editTd);

    return tr;
};

link.editLink = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        $("#link-List div.notification").hide();

        if (result.ErrCode > 0) {
            $("#link-List div.error div").html(result.Reason);
            $("#link-List div.error").show();
            return
        }

        $("#link-Edit .link-Form .link-id").val(result.Link.Id);
        $("#link-Edit .link-Form .link-name").val(result.Link.Name);
        $("#link-Edit .link-Form .link-url").val(result.Link.Url);
        $("#link-Edit .link-Form .link-logo").val(result.Link.Logo);

        $("#link-Edit .link-Form .link-catalog input").prop("checked", false);
        if (result.Link.Catalog) {
            for (var ii = 0; ii < result.Link.Catalog.length; ++ii) {
                var ca = result.Link.Catalog[ii];
                $("#link-Edit .link-Form .link-catalog input").filter("[value=" + ca.Id + "]").prop("checked", true);
            }
        }

        $("#link-content .content-box-tabs li a").removeClass('current');
        $("#link-content .content-box-tabs li a.link-Edit-tab").addClass('current');
        $("#link-Edit").siblings().hide();
        $("#link-Edit").show();
    }, "json");
};

link.deleteLink = function(deleteUrl) {
    $.get(deleteUrl, {}, function(result) {
        $("#link-List div.notification").hide();

        if (result.ErrCode > 0) {
            $("#link-List div.error div").html(result.Reason);
            $("#link-List div.error").show();
            return;
        }

        $("#link-List div.success div").html(result.Reason);
        $("#link-List div.success").show();

        link.refreshLink();
    }, "json");
};