var group = {
    groupInfos: {}
};


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

    var userCountTd = document.createElement("td");
    userCountTd.innerHTML = groupInfo.UserCount;
    tr.appendChild(userCountTd);

    var createrlTd = document.createElement("td");
    createrlTd.innerHTML = groupInfo.Creater.Name;
    tr.appendChild(createrlTd);

    var opTd = document.createElement("td");
    var editLink = document.createElement("a");
    editLink.setAttribute("class", "edit");
    editLink.setAttribute("href", "#editGroup");
    editLink.setAttribute("onclick", "group.editGroup('/admin/account/queryGroup/?id=" + groupInfo.Id + "'); return false");
    var editImage = document.createElement("img");
    editImage.setAttribute("src", "/resources/admin/images/pencil.png");
    editImage.setAttribute("alt", "Edit");
    editLink.appendChild(editImage);
    opTd.appendChild(editLink);

    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteGroup");
    deleteLink.setAttribute("onclick", "group.deleteGroup('/admin/account/deleteGroup/?id=" + groupInfo.Id + "'); return false;");
    var deleteImage = document.createElement("img");
    deleteImage.setAttribute("src", "/resources/admin/images/cross.png");
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
};

group.editGroup = function(editUrl) {
    $.get(editUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            return
        }

        $("#group-Edit .group-Form .group-name").val(result.Group.Name);
        $("#group-Edit .group-Form .group-id").val(result.Group.Id);
        $("#group-Content .content-header .nav .group-Edit").find("a").trigger("click");
    }, "json");
}

group.deleteGroup = function(deleteUrl) {

};