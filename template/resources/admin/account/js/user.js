var user = {
    userInfos: {},
    groupInfos: {}
};


user.initialize = function() {
    user.fillUserListView();
};

user.getUserListView = function() {
    return $("#user-List table");
};

user.constructUserItem = function(userInfo) {
    var tr = document.createElement("tr");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = userInfo.Account;
    tr.appendChild(nameTd);

    var nickTd = document.createElement("td");
    nickTd.innerHTML = userInfo.Name;
    tr.appendChild(nickTd);

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
    editLink.setAttribute("href", "#editArticle");
    editLink.setAttribute("onclick", "article.editArticle('/admin/content/editArticle/?id=" + userInfo.Id + "'); return false");
    var editImage = document.createElement("img");
    editImage.setAttribute("src", "/resources/admin/images/pencil.png");
    editImage.setAttribute("alt", "Edit");
    editLink.appendChild(editImage);
    opTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteArticle");
    deleteLink.setAttribute("onclick", "article.deleteArticle('/admin/content/deleteArticle/?id=" + userInfo.Id + "'); return false;");
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