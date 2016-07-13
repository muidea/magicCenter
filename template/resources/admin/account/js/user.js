var user = {
    userInfos: {},
    groupInfos: {}
};


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
        var group = userInfo.Groups[ii++];
        groups += group.Name;

        if (ii < userInfo.Groups.length) {
            groups += ",";
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
    editLink.setAttribute("onclick", "user.editUser('/admin/account/queryUser/?id=" + userInfo.Id + "'); return false");
    var editImage = document.createElement("img");
    editImage.setAttribute("src", "/resources/admin/images/pencil.png");
    editImage.setAttribute("alt", "Edit");
    editLink.appendChild(editImage);
    opTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteUser");
    deleteLink.setAttribute("onclick", "user.deleteUser('/admin/account/deleteUser/?id=" + userInfo.Id + "'); return false;");
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
};

user.constructGroupItem = function(group) {
    var label = document.createElement("label");

    var chk = document.createElement("input");
    chk.setAttribute("type", "checkbox");
    chk.setAttribute("name", "user-group");
    chk.setAttribute("value", group.Id);
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

user.editUser = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            return
        }

        $("#user-Edit .user-Form .user-account").val(result.User.Name);
        $("#user-Edit .user-Form .user-email").val(result.User.Email);
        $("#user-Edit .user-Form .user-id").val(result.User.Id);

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
            return
        }

    }, "json");
};