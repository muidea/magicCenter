var link = {};

link.constructLinkListlView = function(linkList, catalogList) {
    var linkListView = new Array();
    var offset = 0;
    for (var i = 0; i < linkList.length; ++i) {
        var curLink = linkList[i];
        var catalogNames = "";
        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curCatalog = catalogList[idx];
            for (var j = 0; j < curLink.Catalog; j++) {
                var val = curLink.Catalog[j];
                if (curCatalog.ID == val) {
                    if (catalogNames.length > 0) {
                        catalogNames += ", ";
                    }
                    catalogNames += curCatalog.Name;
                }
            }
        }
        var view = {
            ID: curLink.ID,
            Name: curLink.Name,
            Catalog: catalogNames,
            CreateDate: curLink.CreateDate
        };

        linkListView[offset++] = view;
    }

    return linkListView;
};

link.constructLinkEditView = function(linkList, catalogList) {
    var linkListView = new Array();
    var ii = 0;
    for (var linkIdx = 0; linkIdx < linkList.length; ++linkIdx) {
        var curLink = linkList[linkIdx];

        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curModule = catalogList[idx];

            if (curLink.Module == curModule.ID) {
                var view = {
                    ID: curLink.ID,
                    URL: curLink.URL,
                    Method: curLink.Method,
                    Status: curLink.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                linkListView[ii++] = view;
            }
        }
    }

    return linkListView;
}

link.updateListLinkVM = function(linkList) {
    link.listVM.links = linkList;
}

link.updateEditModuleVM = function(catalogList) {
    link.editVM.modules = catalogList;
};

link.updateEditLinkVM = function(linkList) {
    link.editVM.links = linkList;

    // 将已经enable的link设置上checked标示
    for (var offset = 0; offset < link.editVM.links.length; ++offset) {
        var curLink = link.editVM.links[offset];
        if (curLink.Status > 0) {
            $("#selectLink-List .link_" + curLink.ID).prop("checked", true);
        }
    }
};

// 加载全部的Link
link.getAllLinksAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/content/link/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Link);
            }
        }
    });
};

// 加载全部Catalog
link.getAllCatalogsAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/content/catalog/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Catalog);
            }
        }
    });
};

link.loadData = function(callBack) {
    var getAllCatalogsCallBack = function(errCode, catalogList) {
        if (errCode != 0) {
            return;
        }

        link.catalogs = catalogList;
        if (callBack != null) {
            callBack()
        }
    };

    var getAllLinksCallBack = function(errCode, linkList) {
        if (errCode != 0) {
            return;
        }

        link.links = linkList;
        link.getAllCatalogsAction(getAllCatalogsCallBack);
    };

    // 加载完成
    link.getAllLinksAction(getAllLinksCallBack);
}

link.refreshLinkListView = function(linkList, catalogList) {
    var linksView = link.constructLinkListlView(linkList, catalogList);
    link.updateListLinkVM(linksView);
};

link.refreshLinkEditView = function(link, catalogList) {};

$(document).ready(function() {
    link.listVM = avalon.define({
        $id: "link-List",
        links: []
    });

    link.editVM = avalon.define({
        $id: "link-Edit",
        modules: [],
        links: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        link.updateEditModuleVM(link.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < link.modules.length; idx++) {
                    var curModule = link.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                    }
                }
            }
        );
        link.refreshLinkEditView(link.links, selectModuleArray);
    });

    $("#selectLink-button").click(
        function() {
            var selectLinkList = new Array();
            var offset = 0;
            $("#selectLink-List .link_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectLinkList[offset++] = id;
                }
            );

            var unSelectLinkList = new Array();
            offset = 0;
            $("#selectLink-List .link_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectLinkList[offset++] = id;
                }
            );

            link.statusLinksAction(
                selectLinkList,
                unSelectLinkList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < link.modules.length; idx++) {
                                var curModule = link.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    link.loadData(function() {
                        link.refreshLinkListView(link.links, link.catalogs);

                        link.refreshLinkEditView(link.links, selectModuleArray);
                    })
                });
        }
    );

    link.loadData(function() {
        link.refreshLinkListView(link.links, link.catalogs);
    })
});