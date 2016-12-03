var group = {
    groupInfos: {}
};

$(document).ready(function() {

    // 绑定表单提交事件处理器
    $("#group-Content .group-Form").submit(function() {
        var options = {
            beforeSubmit: showRequest,
            success: showResponse,
            dataType: "json"
        };

        function showRequest() {}

        function showResponse(result) {
            console.log(result);

            if (result.ErrCode > 0) {
                $("#group-Edit .alert-Info .content").html(result.Reason);
                $("#group-Edit .alert-Info").modal();
            } else {
                group.refreshGroups();
            }
        }

        function validate() {
            var result = true;
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

    $("#group-Edit .group-Form button.reset").click(
        function() {
            $("#group-Edit .group-Form .group-id").val("");
            $("#group-Edit .group-Form .group-name").val("");
        });
});

group.initialize = function() {
    group.fillGroupListView();
};

group.getGroupListView = function() {
    return $("#group-List table");
};

group.constructGroupItem = function(groupInfo) {
    var tr = document.createElement("tr");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = groupInfo.Name;
    tr.appendChild(nameTd);

    var descriptionTd = document.createElement("td");
    descriptionTd.innerHTML = groupInfo.Description;
    tr.appendChild(descriptionTd);

    var typeInfoTd = document.createElement("td");
    if (groupInfo.Type == 0) {
        typeInfoTd.innerHTML = "普通组";
    } else {
        typeInfoTd.innerHTML = "管理员组";
    }
    tr.appendChild(typeInfoTd);

    var opTd = document.createElement("td");
    var editLink = document.createElement("a");
    editLink.setAttribute("class", "edit");
    editLink.setAttribute("href", "#editGroup");
    editLink.setAttribute("onclick", "group.editGroup('/account/queryGroup/?id=" + groupInfo.ID + "'); return false");
    var editImage = document.createElement("img");
    editImage.setAttribute("src", "/common/images/pencil.png");
    editImage.setAttribute("alt", "Edit");
    editLink.appendChild(editImage);
    opTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteGroup");
    deleteLink.setAttribute("onclick", "group.deleteGroup('/account/deleteGroup/?id=" + groupInfo.ID + "'); return false;");
    var deleteImage = document.createElement("img");
    deleteImage.setAttribute("src", "/common/images/cross.png");
    deleteImage.setAttribute("alt", "Delete");
    deleteLink.appendChild(deleteImage);
    opTd.appendChild(deleteLink);
    tr.appendChild(opTd);

    return tr;
};

group.fillGroupListView = function() {
    var groupListView = group.getGroupListView();

    $(groupListView).find("tbody tr").remove();
    for (var ii = 0; ii < group.groupInfos.length; ++ii) {
        var groupInfo = group.groupInfos[ii];
        var trContent = group.constructGroupItem(groupInfo);

        $(groupListView).find("tbody").append(trContent);
    }

    $("#group-Edit .group-Form .group-id").val("");
    $("#group-Edit .group-Form .group-name").val("");
    $("#group-Edit .group-Form .group-description").val("");
};

group.refreshGroups = function() {
    $.get("/account/queryAllGroup/", {}, function(result) {
        group.groupInfos = result.Groups;

        group.fillGroupListView();
    }, "json");
};

group.editGroup = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#group-List .alert-Info .content").html(result.Reason);
            $("#group-List .alert-Info").modal();
            return
        }

        $("#group-Edit .group-Form .group-name").val(result.Group.Name);
        $("#group-Edit .group-Form .group-description").val(result.Group.Description);
        $("#group-Edit .group-Form .group-id").val(result.Group.ID);
        $("#group-Content .content-header .nav .group-Edit").find("a").trigger("click");
    }, "json");
}

group.deleteGroup = function(deleteUrl) {
    $.get(deleteUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#group-List .alert-Info .content").html(result.Reason);
            $("#group-List .alert-Info").modal();
            return
        }

        group.refreshGroups();
    }, "json");
};