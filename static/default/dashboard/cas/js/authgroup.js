var authgroup = {};

authgroup.constructAclListlView = function(aclList, authGroupList, moduleList) {
    var aclListView = new Array();
    var offset = 0;
    for (var i = 0; i < aclList.length; ++i) {
        var curAcl = aclList[i];

        var curModule = null;
        for (var idx = 0; idx < moduleList.length; ++idx) {
            var mod = moduleList[idx];
            if (mod.ID == curAcl.Module) {
                curModule = mod;
                break;
            }
        }

        var curAuthGroup = "";
        for (var idx = 0; idx < authGroupList.length; ++idx) {
            var cur = authGroupList[idx];
            for (var ii = 0; ii < curAcl.AuthGroup.length; ++ii) {
                var id = curAcl.AuthGroup[ii];
                if ((cur.ID == id) && (cur.Module == curModule.ID)) {
                    if (curAuthGroup.length > 0) {
                        curAuthGroup += ",";
                    }

                    curAuthGroup += cur.Name;
                }
            }
        }

        var view = {
            ID: curAcl.ID,
            URL: curAcl.URL,
            Method: curAcl.Method,
            Module: curModule.Name,
            AuthGroup: curAuthGroup
        };

        aclListView[offset++] = view;
    }

    return aclListView;
};

authgroup.constructAclEditView = function(aclList, authGroupList, moduleList) {
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

authgroup.updateListAclVM = function(aclList) {
    authgroup.listVM.acls = aclList;
}

authgroup.updateEditModuleVM = function(moduleList) {
    authgroup.editVM.modules = moduleList;
};

authgroup.updateEditAclVM = function(aclList) {
    authgroup.editVM.acls = aclList;

    // 将已经enable的authgroup设置上checked标示
    for (var offset = 0; offset < authgroup.editVM.acls.length; ++offset) {
        var curAcl = authgroup.editVM.acls[offset];
        if (curAcl.Status > 0) {
            $("#selectAcl-List .acl_" + curAcl.ID).prop("checked", true);
        }
    }
};

// 加载全部的Module
authgroup.getAllModulesAction = function(callBack) {
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

// 获取全部的AuthGroup
authgroup.getAllAuthGroupsAction = function(callBack) {
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

// 加载全部已经定义的ACL
authgroup.getAllAclsAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/cas/acl/?module=all&status=1",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.ACLs);
            }
        }
    });
};

// 更新ACL对应的AuthGroup
authgroup.updateAclAuthGroupAction = function(aclID, authGroups, callBack) {
    var strAuthGroups = "";
    for (var ii = 0; ii < authGroups.length; ++ii) {
        strAuthGroups += authGroups[ii];
        strAuthGroups += ",";
    }

    $.ajax({
        type: "POST",
        url: "/cas/acl/authgroup/",
        data: { "acl-id": aclID, "acl-authgroup": strAuthGroups },
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode);
            }
        }
    });
};

authgroup.loadData = function(callBack) {
    var getAllAclsCallBack = function(errCode, aclList) {
        if (errCode != 0) {
            return;
        }

        authgroup.acls = aclList;
        if (callBack != null) {
            callBack()
        }
    };

    var getAllAuthGroupsCallBack = function(errCode, authGroupList) {
        if (errCode != 0) {
            return;
        }

        authgroup.authGroups = authGroupList;
        authgroup.getAllAclsAction(getAllAclsCallBack)
    };

    var getAllModulesCallBack = function(errCode, moduleList) {
        if (errCode != 0) {
            return;
        }

        authgroup.modules = moduleList;
        authgroup.getAllAuthGroupsAction(getAllAuthGroupsCallBack);
    };

    // 加载完成
    authgroup.getAllModulesAction(getAllModulesCallBack);
}

authgroup.refreshAclListView = function(aclList, authGroupList, moduleList) {
    var aclsView = authgroup.constructAclListlView(aclList, authGroupList, moduleList);
    authgroup.updateListAclVM(aclsView);
};

authgroup.refreshAclEditView = function(aclList, authGroupList, moduleList) {
    var moduleNames = "";
    var offset = 0;
    for (var offset = 0; offset < moduleList.length; ++offset) {
        var curModule = moduleList[offset];
        if (moduleNames.length > 0) {
            moduleNames += ", ";
        }
        moduleNames += curModule.Name;
    }

    $("#authgroup-Edit .authgroup-selectModule").val(moduleNames);
    authgroup.updateEditAclVM(authgroup.constructAclEditView(aclList, moduleList));
};

$(document).ready(function() {
    authgroup.listVM = avalon.define({
        $id: "authgroup-List",
        acls: []
    });

    authgroup.editVM = avalon.define({
        $id: "authgroup-Edit",
        modules: [],
        acls: []
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

            authgroup.updateAclAuthGroupAction(
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
                            for (var idx = 0; idx < authgroup.modules.length; idx++) {
                                var curModule = authgroup.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    authgroup.loadData(function() {
                        authgroup.refreshAclListView(authgroup.acls, authgroup.authGroups, authgroup.modules);

                        authgroup.refreshAclEditView(authgroup.acls, selectModuleArray);
                    })
                });
        }
    );

    authgroup.loadData(function() {
        console.log(authgroup.acls);
        authgroup.refreshAclListView(authgroup.acls, authgroup.authGroups, authgroup.modules);
    })
});