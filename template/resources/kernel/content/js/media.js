var media = {
    mediaInfo: {},
    catalogInfo: {},
    userInfo: {}
};


$(document).ready(function() {

    // 绑定表单提交事件处理器
    $("#media-Content .media-Form").submit(function() {
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
            if (result.ErrCode > 0) {
                $("#media-Edit .alert-Info .content").html(result.Reason);
                $("#media-Edit .alert-Info").modal();
            } else {
                media.refreshMedia();
            }
        }

        function validate() {
            var result = true

            $("#media-Content .media-Form .media-url").parent().find("span").remove();
            var url = $("#media-Content .media-Form .media-url").val();
            if (url.length == 0) {
                $("#media-Content .media-Form .media-url").parent().append("<span class=\"input-notification error png_bg\">请输入分类名</span>");
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

    $("#media-Content .media-Form button.reset").click(function() {
        $("#media-Edit .media-Form .media-id").val(-1);
        $("#media-Edit .media-Form .media-name").val("");
        $("#media-Edit .media-Form .media-url").val("");
        //$("#media-Edit .media-Form .media-desc").wysiwyg("setContent", "");
        $("#media-Edit .media-Form .media-desc").val("");
        $("#media-Edit .media-Form .media-catalog input").prop("checked", false);
    });
});


media.initialize = function() {

    media.refreshCatalog();

    media.fillMediaView();

};

media.assignDefaltName = function() {
    var url = $("#media-Edit .media-Form .media-url").val();

    console.log(url);

    if (url.length == 0) {
        return;
    }

    var arr = url.split('\\');
    var fileName = arr[arr.length - 1];
    arr = fileName.split('.');
    $("#media-Edit .media-Form .media-name").val(arr[0]);
}

media.refreshCatalog = function() {
    $("#media-Edit .media-Form .media-catalog").children().remove();
    for (var ii = 0; ii < media.catalogInfo.length; ++ii) {
        var catalog = media.catalogInfo[ii];
        $("#media-Edit .media-Form .media-catalog").append("<label><input type='checkbox' name='media-catalog' value=" + catalog.ID + "> </input>" + catalog.Name + "</label> ");
    }
};


media.refreshMedia = function() {
    $.get("/content/queryAllMedia/", {}, function(result) {
        media.mediaInfo = result.Medias;

        media.fillMediaView();
    }, "json");
};


media.fillMediaView = function() {
    $("#media-List table tbody tr").remove();
    if (media.mediaInfo) {
        for (var ii = 0; ii < media.mediaInfo.length; ++ii) {
            var info = media.mediaInfo[ii];
            var trContent = media.constructMediaItem(info);
            $("#media-List table tbody").append(trContent);
        }
    }

    $("#media-Edit .media-Form .media-id").val(-1);
    $("#media-Edit .media-Form .media-name").val("");
    $("#media-Edit .media-Form .media-url").val("");
    //$("#media-Edit .media-Form .media-desc").wysiwyg("setContent", "");
    $("#media-Edit .media-Form .media-desc").val("");
    $("#media-Edit .media-Form .media-catalog input").prop("checked", false);
};


media.constructMediaItem = function(img) {
    console.log(img);

    var tr = document.createElement("tr");
    tr.setAttribute("class", "media");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = img.Name;
    tr.appendChild(nameTd);

    var urlTd = document.createElement("td");
    urlTd.innerHTML = img.URL;
    tr.appendChild(urlTd);

    var descTd = document.createElement("td");
    descTd.innerHTML = img.Desc;
    tr.appendChild(descTd);

    var catalogTd = document.createElement("td");
    var catalogs = ""
    if (img.Catalog) {
        for (var ii = 0; ii < img.Catalog.length;) {
            var cid = img.Catalog[ii++];
            for (var jj = 0; jj < media.catalogInfo.length;) {
                var catalog = media.catalogInfo[jj++];
                if (catalog.ID == cid) {
                    catalogs += catalog.Name;
                    if (ii < img.Catalog.length) {
                        catalogs += ",";
                    }
                    break;
                }
            }
        }
    }
    catalogs = catalogs.length == 0 ? '-' : catalogs;
    catalogTd.innerHTML = catalogs;
    tr.appendChild(catalogTd);

    var createTd = document.createElement("td");
    var createrValue = "-";
    for (var ii = 0; ii < media.userInfo.length;) {
        var user = media.userInfo[ii++];
        if (user.ID == img.Creater) {
            createrValue = user.Name;
            break;
        }
    }
    createTd.innerHTML = createrValue;
    tr.appendChild(createTd);

    var editTd = document.createElement("td");
    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteMedia");
    deleteLink.setAttribute("onclick", "media.deleteMedia('/content/deleteMedia/?id=" + img.ID + "'); return false;");
    var deleteMedia = document.createElement("img");
    deleteMedia.setAttribute("src", "/resources/admin/images/cross.png");
    deleteMedia.setAttribute("alt", "Delete");
    deleteLink.appendChild(deleteMedia);
    editTd.appendChild(deleteLink);

    tr.appendChild(editTd);

    return tr;
};

media.editMedia = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#media-List .alert-Info .content").html(result.Reason);
            $("#media-List .alert-Info").modal();
            return
        }

        $("#media-Edit .media-Form .media-id").val(result.Media.ID);
        $("#media-Edit .media-Form .media-name").val(result.Media.Name);
        //$("#media-Edit .media-Form .media-url").val(result.Media.Url);
        //$("#media-Edit .media-Form .media-desc").wysiwyg("setContent", result.Media.Desc);
        $("#media-Edit .media-Form .media-desc").val(result.Media.Desc);

        $("#media-Edit .media-Form .media-catalog input").prop("checked", false);
        if (result.Media.Catalog) {
            for (var ii = 0; ii < result.Media.Catalog.length; ++ii) {
                var ca = result.Media.Catalog[ii];
                $("#media-Edit .media-Form .media-catalog input").filter("[value=" + ca + "]").prop("checked", true);
            }
        }

        $("#media-Content .content-header .nav .media-Edit").find("a").trigger("click");
    }, "json");
};

media.deleteMedia = function(deleteUrl) {
    $.get(deleteUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#media-List .alert-Info .content").html(result.Reason);
            $("#media-List .alert-Info").modal();
            return;
        }

        media.refreshMedia();
    }, "json");
};