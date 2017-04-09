var acl = {};

acl.constructAclListlView = function(aclList, moduleList, filterAclFun) {
    var aclListView = new Array();
    var offset = 0;
    for (var i = 0; i < aclList.length; ++i) {
        var curAcl = aclList[i];
        if (!filterAclFun(curAcl)) {
            continue;
        }

        var curModule = null;
        for (var idx = 0; idx < moduleList.length; ++idx) {
            var mod = moduleList[idx];
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

acl.constructAclEditView = function(aclList, moduleList) {
    var aclListView = new Array();
    var ii = 0;
    for (var aclIdx = 0; aclIdx < aclList.length; ++aclIdx) {
        var curAcl = aclList[aclIdx];

        for (var idx = 0; idx < moduleList.length; ++idx) {
            var curModule = moduleList[idx];

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

acl.updateListAclVM = function(aclList) {
    acl.listVM.acls = aclList;
}

acl.updateEditModuleVM = function(moduleList) {
    acl.editVM.modules = moduleList;
};

acl.updateEditAclVM = function(aclList) {
    acl.editVM.acls = aclList;

    // 将已经enable的acl设置上checked标示
    for (var offset = 0; offset < acl.editVM.acls.length; ++offset) {
        var curAcl = acl.editVM.acls[offset];
        if (curAcl.Status > 0) {
            $("#selectAcl-List .acl_" + curAcl.ID).prop("checked", true);
        }
    }
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

acl.loadData = function(callBack) {
    var getAllAclsCallBack = function(errCode, aclList) {
        if (errCode != 0) {
            return;
        }

        acl.acls = aclList;
        if (callBack != null) {
            callBack()
        }
    };

    var getAllModulesCallBack = function(errCode, moduleList) {
        if (errCode != 0) {
            return;
        }

        acl.modules = moduleList;
        acl.getAllAclsAction(getAllAclsCallBack);
    };

    // 加载完成
    acl.getAllModulesAction(getAllModulesCallBack);
}

acl.refreshAclListView = function(aclList, moduleList) {
    var aclsView = acl.constructAclListlView(aclList, moduleList, function(item) {
        return item.Status > 0;
    });
    acl.updateListAclVM(aclsView);
};

acl.refreshAclEditView = function(aclList, moduleList) {
    var moduleNames = "";
    var offset = 0;
    for (var offset = 0; offset < moduleList.length; ++offset) {
        var curModule = moduleList[offset];
        if (moduleNames.length > 0) {
            moduleNames += ", ";
        }
        moduleNames += curModule.Name;
    }

    $("#acl-Edit .acl-selectModule").val(moduleNames);
    acl.updateEditAclVM(acl.constructAclEditView(aclList, moduleList));
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
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < acl.modules.length; idx++) {
                    var curModule = acl.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                    }
                }
            }
        );
        acl.refreshAclEditView(acl.acls, selectModuleArray);
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

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < acl.modules.length; idx++) {
                                var curModule = acl.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    acl.loadData(function() {
                        acl.refreshAclListView(acl.acls, acl.modules);

                        acl.refreshAclEditView(acl.acls, selectModuleArray);
                    })
                });
        }
    );

    acl.loadData(function() {
        acl.refreshAclListView(acl.acls, acl.modules);
    })
});