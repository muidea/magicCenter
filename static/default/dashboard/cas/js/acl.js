var acl = { module: [] };

acl.acl2AclView = function(acls, filterAclFun) {
    var aclListView = new Array();
    for (var i = 0; i < acls.length; ++i) {
        var curAcl = acls[i];

        if (!filterAclFun(curAcl)) {
            continue;
        }

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

acl.filterModuleAclView = function(modules) {
    var aclListView = new Array();
    var ii = 0;
    for (var aclIdx = 0; aclIdx < acl.acl.length; ++aclIdx) {
        var curAcl = acl.acl[aclIdx];

        for (var idx = 0; idx < modules.length; ++idx) {
            var curModule = modules[idx];

            if (curAcl.Module == curModule.ID) {
                var view = {
                    ID: curAcl.ID,
                    URL: curAcl.URL,
                    Method: curAcl.Method,
                    Module: curModule.Name,
                    ModuleID: curModule.ID,
                    Enable: curAcl.Enable
                }

                aclListView[ii++] = view;
            }
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
    acl.listVM.acl = acl.acl2AclView(acls, function(item) {
        return item.Enable > 0;
    });
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
        acl.editVM.acl = acl.filterModuleAclView(selectModuleArray)
        for (var offset = 0; offset < acl.editVM.acl.length; ++offset) {
            var curAcl = acl.editVM.acl[offset];
            if (curAcl.Enable > 0) {
                console.log($("#selectAcl-List .acl_" + curAcl.ID));

                $("#selectAcl-List .acl_" + curAcl.ID).prop("checked", true);
            }
        }
    });

    $("#selectAcl-button").click(
        function() {
            var aclList = new Array();
            var offset = 0;
            var selectAcl = $("#selectAcl-List .acl-id:checked");
            for (var idx = 0; idx < acl.editVM.acl.length; ++idx) {
                var curAcl = acl.editVM.acl[idx];
                selectAcl.each(
                    function() {
                        var id = $(this).val();
                        if (id == curAcl.ID) {
                            var val = {
                                ID: curAcl.ID,
                                URL: curAcl.URL,
                                Method: curAcl.Method,
                                Module: curAcl.ModuleID
                            }

                            aclList[offset++] = val;
                        }
                    }
                );
            }

            console.log(aclList);
        }
    );
});