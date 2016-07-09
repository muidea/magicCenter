var page = {
    moduleList: {},
    currentModule: {}
};

$(document).ready(function () {

});

page.getModuleListView = function () {
    return $("#module-List table");
};

page.getModulePageListView = function () {
    return $("#page-List table");
};

page.getBlockListView = function() {
    return $("#page-List .page-Form .page-block");
};

page.initialize = function () {
    page.fillModuleListView();
};

page.fillModuleListView = function () {
    var moduleListView = page.getModuleListView()
    $(moduleListView).find("tbody tr").remove();
    for (var ii = 0; ii < page.moduleList.length; ++ii) {
        var info = page.moduleList[ii];
        var trModule = page.constructModuleItem(info);
        $(moduleListView).find("tbody").append(trModule);
    }
};

page.fillPageListView = function () {
    var pageListView = page.getModulePageListView();
    $(pageListView).find("tbody tr").remove();
    for (var ii = 0; ii < page.currentModule.Pages.length; ++ii) {
        var info = page.currentModule.Pages[ii];
        var trPage = page.constructPageItem(info);
        $(pageListView).find("tbody").append(trPage);
    }

    var blockListView = page.getBlockListView();
    $(blockListView).find("label").remove();
    for (var ii =0; ii < page.currentModule.Blocks.length; ++ii) {
        var cur = page.currentModule.Blocks[ii];
        var label = page.constructBlockItem(cur);
        $(blockListView).append(label);
    }
}

page.constructModuleItem = function (mod) {
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
    settingButton.setAttribute("onclick", "page.maintainPage('/admin/system/queryModuleDetail/?id=" + mod.Id + "'); return false;");
    settingButton.setAttribute("value", "页面设置");
    settingButton.setAttribute("class", "button");
    editTd.appendChild(settingButton);
    tr.appendChild(editTd);
    return tr;
};

page.constructPageItem = function (page) {
    var tr = document.createElement("tr");
    tr.setAttribute("class", "module");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = page.Url;
    tr.appendChild(nameTd);

    var blocks = "";
    if (page.Blocks) {
        for (var ii = 0; ii < page.Blocks.length;) {
            var block = page.Blocks[ii++];
            blocks += block.Name;
            if (ii < page.Blocks.length) {
                blocks += ",";
            }
        }
    } else {
        blocks = "-";
    }
    var blocksTd = document.createElement("td");
    blocksTd.innerHTML = blocks;
    tr.appendChild(blocksTd);
    tr.setAttribute("onclick", "page.selectPage('" + page.Url + "'); return false;");
    return tr;
};

page.constructBlockItem = function(block) {
    var label = document.createElement("label");

    var chk = document.createElement("input");
    chk.setAttribute("type", "checkbox");
    chk.setAttribute("value", block.Id);
    label.appendChild(chk);

    var span = document.createElement("span");
    span.innerHTML = block.Name;
    label.appendChild(span);
    label.setAttribute("class", "text-center");

    return label;
}

page.selectPage = function (pageUrl) {
    if (!page.currentModule) {
        return;
    }

    var blockListView = page.getBlockListView();
    $(blockListView).find("input ").prop("checked", false);
    for (var ii = 0; ii < page.currentModule.Pages.length; ++ii) {
        var cur = page.currentModule.Pages[ii];
        if (cur.Url == pageUrl) {
            $("#page-List .page-Form .page-url").val(cur.Url);
            $("#page-List .page-Form .page-owner").val(cur.Owner);

            for (var ii =0; ii < cur.Blocks.length; ++ii) {
                var block = cur.Blocks[ii];
                $(blockListView).find("input ").filter("[value=" + block.Id + "]").prop("checked", true);
            }

            break;
        }
    }

}

page.maintainPage = function (maintainUrl) {
    $.get(maintainUrl, {}, function (result) {

        if (result.ErrCode > 0) {
            return
        }

        page.currentModule = result.Module;
        page.fillPageListView();
        $("#module-Page .content-header .nav .page-List").find("a").trigger("click");
    }, "json");
};