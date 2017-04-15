var user = {};

user.constructUserListlView = function(userList) {
    var userListView = new Array();
    var offset = 0;
    for (var i = 0; i < userList.length; ++i) {
        var curUser = userList[i];

        var view = {
            ID: curUser.ID,
            Account: curUser.Account,
            Email: curUser.Email,
            Status: curUser.Status
        };

        userListView[offset++] = view;
    }

    return userListView;
};

user.constructUserEditView = function(userList, groupList) {
    var userListView = new Array();
    var ii = 0;
    for (var userIdx = 0; userIdx < userList.length; ++userIdx) {
        var curUser = userList[userIdx];

        for (var idx = 0; idx < groupList.length; ++idx) {
            var curModule = groupList[idx];

            if (curUser.Module == curModule.ID) {
                var view = {
                    ID: curUser.ID,
                    URL: curUser.URL,
                    Method: curUser.Method,
                    Status: curUser.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                userListView[ii++] = view;
            }
        }
    }

    return userListView;
}

user.updateListUserVM = function(userList) {
    user.listVM.users = userList;
}

user.updateEditModuleVM = function(groupList) {
    user.editVM.modules = groupList;
};

user.updateEditUserVM = function(userList) {
    user.editVM.users = userList;

    // 将已经enable的user设置上checked标示
    for (var offset = 0; offset < user.editVM.users.length; ++offset) {
        var curUser = user.editVM.users[offset];
        if (curUser.Status > 0) {
            $("#selectUser-List .user_" + curUser.ID).prop("checked", true);
        }
    }
};

// 加载全部的User
user.getAllUsersAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/account/user/",
        data: {},
        dataType: "json",
        success: function(data) {
            console.log(data);
            if (callBack != null) {
                callBack(data.ErrCode, data.User);
            }
        }
    });
};

user.loadData = function(callBack) {
    var getAllUsersCallBack = function(errCode, userList) {
        if (errCode != 0) {
            return;
        }

        user.users = userList;
        if (callBack != null) {
            callBack()
        }
    };

    // 加载完成
    user.getAllUsersAction(getAllUsersCallBack);
}

user.refreshUserListView = function(userList, groupList) {
    var usersView = user.constructUserListlView(userList, groupList);
    user.updateListUserVM(usersView);
};

user.refreshUserEditView = function(user, groupList) {};

$(document).ready(function() {
    user.listVM = avalon.define({
        $id: "user-List",
        users: []
    });

    user.editVM = avalon.define({
        $id: "user-Edit",
        modules: [],
        users: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        user.updateEditModuleVM(user.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < user.modules.length; idx++) {
                    var curModule = user.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                    }
                }
            }
        );
        user.refreshUserEditView(user.users, selectModuleArray);
    });

    $("#selectUser-button").click(
        function() {
            var selectUserList = new Array();
            var offset = 0;
            $("#selectUser-List .user_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectUserList[offset++] = id;
                }
            );

            var unSelectUserList = new Array();
            offset = 0;
            $("#selectUser-List .user_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectUserList[offset++] = id;
                }
            );

            user.statusUsersAction(
                selectUserList,
                unSelectUserList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < user.modules.length; idx++) {
                                var curModule = user.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    user.loadData(function() {
                        user.refreshUserListView(user.users, user.groups);

                        user.refreshUserEditView(user.users, selectModuleArray);
                    })
                });
        }
    );

    user.loadData(function() {
        user.refreshUserListView(user.users, user.groups);
    })
});