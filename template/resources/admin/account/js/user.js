var user = {
    userInfos: {},
    groupInfos: {}
};

$(document).ready(function() {

    // 绑定表单提交事件处理器
    $("#user-Content .user-Form").submit(function() {
        var options = {
            beforeSubmit: showRequest,
            success: showResponse,
            dataType: "json"
        };

        function showRequest() {}

        function showResponse(result) {

            if (result.ErrCode > 0) {
                $("#user-Edit .alert-Info .content").html(result.Reason);
                $("#user-Edit .alert-Info").modal();
            } else {
                user.refreshUser();
                user.fillUserListView();
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

    $("#user-Content .user-Form button.reset").click(function() {
        $("#user-Edit .user-Form .user-account").val("");
        $("#user-Edit .user-Form .user-email").val("");
        $("#user-Edit .user-Form .user-group").prop("checked", false);
        $("#user-Edit .user-Form .user-id").val("-1");
    });
});

user.initialize = function() {
    user.fillUserListView();

    user.fillGroupListView();
};

user.getUserListView = function() {
    return $("#user-List table");
};

user.getGroupListView = function() {
    return $("#user-Edit .user-Form .user-group");
};

user.constructUserItem = function(userInfo) {
    var tr = document.createElement("tr");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = userInfo.Account;
    tr.appendChild(nameTd);

    var emailTd = document.createElement("td");
    emailTd.innerHTML = userInfo.Email;
    tr.appendChild(emailTd);

    var groupTd = document.createElement("td");
    var groups = "";
    for (var ii = 0; ii < userInfo.Groups.length;) {
        var gid = userInfo.Groups[ii++];
        for (var jj = 0; jj < user.groupInfos.length;) {
            var group = user.groupInfos[jj++];
            if (group.ID == gid) {
                groups += group.Name;

                if (ii < userInfo.Groups.length) {
                    groups += ",";
                }
                break;
            }
        }
    }
    groupTd.innerHTML = groups;
    tr.appendChild(groupTd);

    var statusTd = document.createElement("td");
    if (userInfo.Status == 0) {
        statusTd.innerHTML = "正常";
    } else {
        statusTd.innerHTML = "锁定";
    }
    tr.appendChild(statusTd);

    var opTd = document.createElement("td");
    var editLink = document.createElement("a");
    editLink.setAttribute("class", "edit");
    editLink.setAttribute("href", "#editUser");
    editLink.setAttribute("onclick", "user.editUser('/account/queryUser/?id=" + userInfo.ID + "'); return false");
    var editImage = document.createElement("img");
    editImage.setAttribute("src", "/resources/admin/images/pencil.png");
    editImage.setAttribute("alt", "Edit");
    editLink.appendChild(editImage);
    opTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteUser");
    deleteLink.setAttribute("onclick", "user.deleteUser('/account/deleteUser/?id=" + userInfo.ID + "'); return false;");
    var deleteImage = document.createElement("img");
    deleteImage.setAttribute("src", "/resources/admin/images/cross.png");
    deleteImage.setAttribute("alt", "Delete");
    deleteLink.appendChild(deleteImage);
    opTd.appendChild(deleteLink);
    tr.appendChild(opTd);

    return tr;
};

user.fillUserListView = function() {
    var userListView = user.getUserListView();

    $(userListView).find("tbody tr").remove();
    for (var ii = 0; ii < user.userInfos.length; ++ii) {
        var userInfo = user.userInfos[ii];
        var trContent = user.constructUserItem(userInfo);

        $(userListView).find("tbody").append(trContent);
    }

    $("#user-Edit .user-Form .user-account").val("");
    $("#user-Edit .user-Form .user-email").val("");
    $("#user-Edit .user-Form .user-group").prop("checked", false);
    $("#user-Edit .user-Form .user-id").val("-1");
};

user.constructGroupItem = function(group) {
    var label = document.createElement("label");

    var chk = document.createElement("input");
    chk.setAttribute("type", "checkbox");
    chk.setAttribute("name", "user-group");
    chk.setAttribute("class", "user-group");
    chk.setAttribute("value", group.ID);
    label.appendChild(chk);

    var span = document.createElement("span");
    span.innerHTML = group.Name;
    label.appendChild(span);
    label.setAttribute("class", "text-center");

    return label;
};

user.fillGroupListView = function() {
    var groupListView = user.getGroupListView();

    $(groupListView).find("label").remove();
    for (var ii = 0; ii < user.groupInfos.length; ++ii) {
        var cur = user.groupInfos[ii];
        var label = user.constructGroupItem(cur);
        $(groupListView).append(label);
    }
}

user.refreshUser = function() {
    $.get("/account/queryAllUser/", {}, function(result) {
        user.userInfos = result.Users;

        user.fillUserListView();
    }, "json");
};

user.editUser = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#user-List .alert-Info .content").html(result.Reason);
            $("#user-List .alert-Info").modal();
            return
        }

        $("#user-Edit .user-Form .user-account").val(result.User.Name);
        $("#user-Edit .user-Form .user-email").val(result.User.Email);
        $("#user-Edit .user-Form .user-id").val(result.User.ID);

        var groupListView = user.getGroupListView();
        for (var ii = 0; ii < result.User.Groups.length; ++ii) {
            var gid = result.User.Groups[ii];
            $(groupListView).find("input ").filter("[value=" + gid + "]").prop("checked", true);
        }

        $("#user-Content .content-header .nav .user-Edit").find("a").trigger("click");
    }, "json");
}

user.deleteUser = function(deleteUrl) {
    $.get(deleteUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#user-List .alert-Info .content").html(result.Reason);
            $("#user-List .alert-Info").modal();
            return
        }

        user.refreshUser();
    }, "json");
};