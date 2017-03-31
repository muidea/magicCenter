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

acl.module2AclView = function(modules) {
    var aclListView = new Array();
    var ii = 0;
    for (var idx = 0; idx < modules.length; ++idx) {
        var curModule = modules[idx];
        for (var rIdx = 0; rIdx < curModule.Route.length; ++rIdx) {
            var curRoute = curModule.Route[rIdx];
            var view = {
                URL: curRoute.Pattern,
                Method: curRoute.Method,
                Module: curModule.Name,
                ModuleID: curModule.ID
            }

            aclListView[ii++] = view;
        }
    }

    return aclListView;
}


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

    $('#moduleListModal').on('show.bs.modal', function(e) {
        acl.updateEditVM(acl.module);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleNames = "";
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < acl.module.length; idx++) {
                    var cur = acl.module[idx];
                    if (cur.ID == id) {
                        selectModuleArray[offset++] = cur;
                        if (selectModuleNames.length > 0) {
                            selectModuleNames += ", ";
                        }
                        selectModuleNames += cur.Name;
                    }
                }
            }
        );

        $("#acl-Edit .acl-selectModule").val(selectModuleNames);
        acl.editVM.acl = acl.module2AclView(selectModuleArray)
    });

    $("#moduleListModal .btn-primary").click(
        function() {
            $('#moduleListModal').modal('hide')
        }
    );
});