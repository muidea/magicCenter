var acl = { module: [] };

acl.acl2AclView = function(acls) {
    var aclListView = new Array();
    for (var i = 0; i < acls.length; ++i) {
        var curAcl = acls[i];

        var curModule = null;
        for (var idx = 0; idx < acl.module.length; ++idx) {
            var mod = acl.module[idx];
            if (mod.ID == curAcl.Module) {
                curModule = mod;
                break;
            }
        }
        var view = {
            ID: curAcl.ID,
            URL: curAcl.URL,
            Method: curAcl.Method,
            Module: curModule.Name
        };

        aclListView[i] = view;
    }

    return aclListView;
};

// 加载全部的Module
acl.getModules = function(callBack) {
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

acl.getModulesCallBack = function(errCode, modules) {
    if (errCode != 0) {
        return;
    }

    acl.module = modules;
    acl.getAcls(acl.getAclsCallBack);
};

// 加载全部已经定义的ACL
acl.getAcls = function(callBack) {
    var acls = [];
    $.ajax({
        type: "GET",
        url: "/cas/acl/?module=all",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.ACLs);
            }
        }
    });
};

acl.getAclsCallBack = function(errCode, acls) {
    if (errCode != 0) {
        return;
    }

    acl.acl = acls;
    acl.listVM.acl = acl.acl2AclView(acls);
};

acl.updateEditVM = function(modules) {
    acl.editVM.module = modules
}

$(document).ready(function() {

    acl.listVM = avalon.define({
        $id: "acl-List",
        acl: []
    });

    acl.editVM = avalon.define({
        $id: "acl-Edit",
        module: [],
        acl: {}
    });

    // 加载完成
    acl.getModules(acl.getModulesCallBack);

    $("#selectModule").click(
        function() {
            acl.updateEditVM(acl.module);
        }
    );
});