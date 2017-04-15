var catalog = {};

catalog.constructCatalogListlView = function(catalogList, parentCatalogList) {
    var catalogListView = new Array();
    var offset = 0;
    for (var i = 0; i < catalogList.length; ++i) {
        var curCatalog = catalogList[i];
        var catalogNames = "";
        for (var idx = 0; idx < parentCatalogList.length; ++idx) {
            var curParentCatalog = parentCatalogList[idx];
            for (var j = 0; j < curCatalog.Catalog; j++) {
                var val = curCatalog.Catalog[j];
                if (curParentCatalog.ID == val) {
                    if (catalogNames.length > 0) {
                        catalogNames += ", ";
                    }
                    catalogNames += curParentCatalog.Name;
                }
            }
        }
        var view = {
            ID: curCatalog.ID,
            Name: curCatalog.Name,
            Catalog: catalogNames,
            CreateDate: curCatalog.CreateDate
        };

        catalogListView[offset++] = view;
    }

    return catalogListView;
};

catalog.constructCatalogEditView = function(catalogList, catalogList) {
    var catalogListView = new Array();
    var ii = 0;
    for (var catalogIdx = 0; catalogIdx < catalogList.length; ++catalogIdx) {
        var curCatalog = catalogList[catalogIdx];

        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curModule = catalogList[idx];

            if (curCatalog.Module == curModule.ID) {
                var view = {
                    ID: curCatalog.ID,
                    URL: curCatalog.URL,
                    Method: curCatalog.Method,
                    Status: curCatalog.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                catalogListView[ii++] = view;
            }
        }
    }

    return catalogListView;
}

catalog.updateListCatalogVM = function(catalogList) {
    catalog.listVM.catalogs = catalogList;
}

catalog.updateEditModuleVM = function(catalogList) {
    catalog.editVM.modules = catalogList;
};

catalog.updateEditCatalogVM = function(catalogList) {
    catalog.editVM.catalogs = catalogList;

    // 将已经enable的catalog设置上checked标示
    for (var offset = 0; offset < catalog.editVM.catalogs.length; ++offset) {
        var curCatalog = catalog.editVM.catalogs[offset];
        if (curCatalog.Status > 0) {
            $("#selectCatalog-List .catalog_" + curCatalog.ID).prop("checked", true);
        }
    }
};

// 加载全部Catalog
catalog.getAllCatalogsAction = function(callBack) {
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

catalog.loadData = function(callBack) {
    var getAllCatalogsCallBack = function(errCode, catalogList) {
        if (errCode != 0) {
            return;
        }

        catalog.catalogs = catalogList;
        if (callBack != null) {
            callBack()
        }
    };

    // 加载完成
    catalog.getAllCatalogsAction(getAllCatalogsCallBack);
}

catalog.refreshCatalogListView = function(catalogList, catalogList) {
    var catalogsView = catalog.constructCatalogListlView(catalogList, catalogList);
    catalog.updateListCatalogVM(catalogsView);
};

catalog.refreshCatalogEditView = function(catalog, catalogList) {};

$(document).ready(function() {
    catalog.listVM = avalon.define({
        $id: "catalog-List",
        catalogs: []
    });

    catalog.editVM = avalon.define({
        $id: "catalog-Edit",
        modules: [],
        catalogs: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        catalog.updateEditModuleVM(catalog.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < catalog.modules.length; idx++) {
                    var curModule = catalog.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                    }
                }
            }
        );
        catalog.refreshCatalogEditView(catalog.catalogs, selectModuleArray);
    });

    $("#selectCatalog-button").click(
        function() {
            var selectCatalogList = new Array();
            var offset = 0;
            $("#selectCatalog-List .catalog_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectCatalogList[offset++] = id;
                }
            );

            var unSelectCatalogList = new Array();
            offset = 0;
            $("#selectCatalog-List .catalog_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectCatalogList[offset++] = id;
                }
            );

            catalog.statusCatalogsAction(
                selectCatalogList,
                unSelectCatalogList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < catalog.modules.length; idx++) {
                                var curModule = catalog.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    catalog.loadData(function() {
                        catalog.refreshCatalogListView(catalog.catalogs, catalog.catalogs);

                        catalog.refreshCatalogEditView(catalog.catalogs, selectModuleArray);
                    })
                });
        }
    );

    catalog.loadData(function() {
        catalog.refreshCatalogListView(catalog.catalogs, catalog.catalogs);
    })
});