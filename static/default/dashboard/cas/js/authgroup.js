var authgroup = { module: [], authGroup: [] };

authgroup.getModules = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/dashboard/module/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Module);
            }
        }
    });
};

authgroup.getModulesCallBack = function(errCode, modules) {
    if (errCode != 0) {
        return;
    }

    authgroup.module = modules;
};

authgroup.getAuthGroups = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/cas/authgroup/?module=all",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.AuthGroup);
            }
        }
    });
}

authgroup.getAuthGroupsCallBack = function(errCode, authgroups) {
    if (errCode != 0) {
        return;
    }

    authgroup.authGroup = authgroups;
}

$(document).ready(function() {

    authgroup.listVM = avalon.define({
        $id: "authgroup-List",
        acl: []
    });

    authgroup.editVM = avalon.define({
        $id: "authgroup-Edit",
        acl: {}
    });

    // 加载完成
    authgroup.getModules(authgroup.getModulesCallBack);
});