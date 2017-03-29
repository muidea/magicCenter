var acl = {};

acl.acl2AclView = function(acls) {
    var ret = new Array();
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

        ret[i] = view;
    }

    return ret;
};

$(document).ready(function() {

    acl.listVM = avalon.define({
        $id: "acl-List",
        acl: []
    });

    $.ajax({
        type: "GET",
        url: "/dashboard/module/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (data.ErrCode == 0) {
                acl.module = data.Module;
            }
        }
    });

    $.ajax({
        type: "GET",
        url: "/cas/acl/?module=all",
        data: {},
        dataType: "json",
        success: function(data) {
            if (data.ErrCode == 0) {
                acl.listVM.acl = acl.acl2AclView(data.ACLs);
            }
        }
    });
});