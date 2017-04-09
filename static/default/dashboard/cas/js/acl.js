var acl = {};

acl.constructAclListlView = function(acls, modules, filterAclFun) {
    var aclListView = new Array();
    var offset = 0;
    for (var i = 0; i < acls.length; ++i) {
        var curAcl = acls[i];
        if (!filterAclFun(curAcl)) {
            continue;
        }

        var curModule = null;
        for (var idx = 0; idx < modules.length; ++idx) {
            var mod = modules[idx];
            if (mod.ID == curAcl.Module) {
                curModule = mod;
                break;
            }
        }
        var view = {
            ID: curAcl.ID,
            URL: curAcl.URL,
            Method: curAcl.Method,
            Status: curAcl.Status,
            Module: curModule.Name,
            ModuleID: curModule.ID
        };

        aclListView[offset++] = view;
    }

    return aclListView;
};

acl.constructAclEditView = function(acls, modules) {
    var aclListView = new Array();
    var ii = 0;
    for (var aclIdx = 0; aclIdx < acls.length; ++aclIdx) {
        var curAcl = acls[aclIdx];

        for (var idx = 0; idx < modules.length; ++idx) {
            var curModule = modules[idx];

            if (curAcl.Module == curModule.ID) {
                var view = {
                    ID: curAcl.ID,
                    URL: curAcl.URL,
                    Method: curAcl.Method,
                    Status: curAcl.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                aclListView[ii++] = view;
            }
        }
    }

    return aclListView;
}

acl.updateListAclVM = function(acls) {
    acl.listVM.acls = acls;
}

acl.updateEditModuleVM = function(modules) {
    acl.editVM.modules = modules;
};

acl.updateEditAclVM = function(acls) {
    acl.editVM.acls = acls;
};

// 加载全部的Module
acl.getAllModulesAction = function(callBack) {
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

// 加载全部已经定义的ACL
acl.getAllAclsAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/cas/acl/?module=all&status=-1",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.ACLs);
            }
        }
    });
};

acl.statusAclsAction = function(enableAcls, disableAcls, callBack) {
    var strEnableAlcs = "";
    for (var ii = 0; ii < enableAcls.length; ++ii) {
        strEnableAlcs += enableAcls[ii];
        strEnableAlcs += ",";
    }

    var strDisableAlcs = "";
    for (var ii = 0; ii < disableAcls.length; ++ii) {
        strDisableAlcs += disableAcls[ii];
        strDisableAlcs += ",";
    }

    $.ajax({
        type: "POST",
        url: "/cas/acl/status/",
        data: { "enable-list": strEnableAlcs, "disable-list": strDisableAlcs },
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode);
            }
        }
    });
};

acl.refreshView = function() {
    var getAllAclsCallBack = function(errCode, acls) {
        if (errCode != 0) {
            return;
        }

        acl.acls = acls;
        var aclsView = acl.constructAclListlView(acl.acls, acl.modules, function(item) {
            return item.Status > 0;
        });
        acl.updateListAclVM(aclsView);
    };

    var getAllModulesCallBack = function(errCode, modules) {
        if (errCode != 0) {
            return;
        }

        acl.modules = modules;
        acl.getAllAclsAction(getAllAclsCallBack);
    };

    // 加载完成
    acl.getAllModulesAction(getAllModulesCallBack);
};

$(document).ready(function() {

    acl.listVM = avalon.define({
        $id: "acl-List",
        acls: []
    });

    acl.editVM = avalon.define({
        $id: "acl-Edit",
        modules: [],
        acls: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        acl.updateEditModuleVM(acl.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleNames = "";
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < acl.modules.length; idx++) {
                    var curModule = acl.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                        if (selectModuleNames.length > 0) {
                            selectModuleNames += ", ";
                        }
                        selectModuleNames += curModule.Name;
                    }
                }
            }
        );

        $("#acl-Edit .acl-selectModule").val(selectModuleNames);
        acl.updateEditAclVM(acl.constructAclEditView(acl.acls, selectModuleArray));

        // 将已经enable的acl设置上checked标示
        for (var offset = 0; offset < acl.editVM.acls.length; ++offset) {
            var curAcl = acl.editVM.acls[offset];
            if (curAcl.Status > 0) {
                $("#selectAcl-List .acl_" + curAcl.ID).prop("checked", true);
            }
        }
    });

    $("#selectAcl-button").click(
        function() {
            var selectAclList = new Array();
            var offset = 0;
            $("#selectAcl-List .acl_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectAclList[offset++] = id;
                }
            );

            var unSelectAclList = new Array();
            offset = 0;
            $("#selectAcl-List .acl_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectAclList[offset++] = id;
                }
            );

            acl.statusAclsAction(
                selectAclList,
                unSelectAclList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    acl.refreshView();
                });

        }
    );

    acl.refreshView();
});