var group = {};

group.constructGroupListlView = function(groupList) {
    var groupListView = new Array();
    var offset = 0;
    for (var i = 0; i < groupList.length; ++i) {
        var curGroup = groupList[i];

        var view = {
            ID: curGroup.ID,
            Name: curGroup.Name,
            Description: curGroup.Description,
            Type: curGroup.Type
        };

        groupListView[offset++] = view;
    }

    return groupListView;
};

group.constructGroupEditView = function(groupList, groupList) {
    var groupListView = new Array();
    var ii = 0;
    for (var groupIdx = 0; groupIdx < groupList.length; ++groupIdx) {
        var curGroup = groupList[groupIdx];

        for (var idx = 0; idx < groupList.length; ++idx) {
            var curModule = groupList[idx];

            if (curGroup.Module == curModule.ID) {
                var view = {
                    ID: curGroup.ID,
                    URL: curGroup.URL,
                    Method: curGroup.Method,
                    Status: curGroup.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                groupListView[ii++] = view;
            }
        }
    }

    return groupListView;
}

group.updateListGroupVM = function(groupList) {
    group.listVM.groups = groupList;
}

group.updateEditModuleVM = function(groupList) {
    group.editVM.modules = groupList;
};

group.updateEditGroupVM = function(groupList) {
    group.editVM.groups = groupList;

    // 将已经enable的group设置上checked标示
    for (var offset = 0; offset < group.editVM.groups.length; ++offset) {
        var curGroup = group.editVM.groups[offset];
        if (curGroup.Status > 0) {
            $("#selectGroup-List .group_" + curGroup.ID).prop("checked", true);
        }
    }
};

// 加载全部的Group
group.getAllGroupsAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/account/group/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Group);
            }
        }
    });
};

group.loadData = function(callBack) {
    var getAllGroupsCallBack = function(errCode, groupList) {
        if (errCode != 0) {
            return;
        }

        group.groups = groupList;
        if (callBack != null) {
            callBack()
        }
    };

    // 加载完成
    group.getAllGroupsAction(getAllGroupsCallBack);
}

group.refreshGroupListView = function(groupList, groupList) {
    var groupsView = group.constructGroupListlView(groupList, groupList);
    group.updateListGroupVM(groupsView);
};

group.refreshGroupEditView = function(group, groupList) {};

$(document).ready(function() {
    group.listVM = avalon.define({
        $id: "group-List",
        groups: []
    });

    group.editVM = avalon.define({
        $id: "group-Edit",
        modules: [],
        groups: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        group.updateEditModuleVM(group.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < group.modules.length; idx++) {
                    var curModule = group.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                    }
                }
            }
        );
        group.refreshGroupEditView(group.groups, selectModuleArray);
    });

    $("#selectGroup-button").click(
        function() {
            var selectGroupList = new Array();
            var offset = 0;
            $("#selectGroup-List .group_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectGroupList[offset++] = id;
                }
            );

            var unSelectGroupList = new Array();
            offset = 0;
            $("#selectGroup-List .group_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectGroupList[offset++] = id;
                }
            );

            group.statusGroupsAction(
                selectGroupList,
                unSelectGroupList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < group.modules.length; idx++) {
                                var curModule = group.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    group.loadData(function() {
                        group.refreshGroupListView(group.groups, group.groups);

                        group.refreshGroupEditView(group.groups, selectModuleArray);
                    })
                });
        }
    );

    group.loadData(function() {
        group.refreshGroupListView(group.groups, group.groups);
    })
});