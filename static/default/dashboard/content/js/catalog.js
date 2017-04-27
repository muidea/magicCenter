var catalog = {};

catalog.constructCatalogListlView = function(catalogList, parentCatalogList) {
    var catalogListView = new Array();
    var offset = 0;
    for (var ii = 0; ii < catalogList.length; ++ii) {
        var curCatalog = catalogList[ii];
        var catalogNames = "";
        if (curCatalog.Catalog) {
            for (var idx = 0; idx < parentCatalogList.length; ++idx) {
                var curParentCatalog = parentCatalogList[idx];
                for (var jj = 0; jj < curCatalog.Catalog.length; jj++) {
                    var val = curCatalog.Catalog[jj];
                    if (curParentCatalog.ID == val) {
                        if (catalogNames.length > 0) {
                            catalogNames += ", ";
                        }
                        catalogNames += curParentCatalog.Name;
                    }
                }
            }
        }
        var view = {
            ID: curCatalog.ID,
            Name: curCatalog.Name,
            Catalog: catalogNames,
            CreateDate: curCatalog.CreateDate
        };

        catalogListView[ii] = view;
    }

    return catalogListView;
};

catalog.constructCatalogEditView = function(catalogList, filterFun) {
    var catalogListView = new Array();
    var ii = 0;
    for (var idx = 0; idx < catalogList.length; ++idx) {
        var curCatalog = catalogList[idx];
        if (!filterFun(curCatalog)) {
            continue;
        }
        var view = {
            ID: curCatalog.ID,
            Name: curCatalog.Name
        };

        catalogListView[ii++] = view;
    }

    return catalogListView;
};

catalog.updateListCatalogVM = function(catalogList) {
    catalog.listVM.catalogs = catalogList;
};

catalog.updateEditCatalogVM = function(curCatalog, parentCatalogList) {
    catalog.editVM.catalog = curCatalog;
    catalog.editVM.parentCatalogs = parentCatalogList;
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

catalog.saveCatalogAction = function(name, description, parents, callBack) {
    $.ajax({
        type: "POST",
        url: "/content/catalog/",
        data: { "catalog-name": name, "catalog-description": description, "catalog-parent": parents },
        dataType: "json",
        success: function(data) {
            console.log(data);

            if (callBack != null) {
                callBack(data.ErrCode, data.Catalog);
            }
        }
    });
}

catalog.loadData = function(callBack) {
    var getAllCatalogsCallBack = function(errCode, catalogList) {
        if (errCode != 0) {
            return;
        }
        catalog.curCatalog = { ID: -1, Name: "", Description: "", ParentCatalog: [] };
        catalog.catalogs = catalogList;
        if (callBack != null) {
            callBack()
        }
    };

    // 加载完成
    catalog.getAllCatalogsAction(getAllCatalogsCallBack);
};

catalog.refreshCatalogListView = function(catalogList, catalogList) {
    var catalogsView = catalog.constructCatalogListlView(catalogList, catalogList);
    catalog.updateListCatalogVM(catalogsView);
};

catalog.refreshCatalogEditView = function(curCatalog, catalogList) {
    var parentCatalogView = catalog.constructCatalogEditView(catalogList, function(catalog) {
        return true;
    });
    catalog.updateEditCatalogVM({ ID: curCatalog.ID, Name: curCatalog.Name, Description: curCatalog.Description, ParentCatalog: curCatalog.ParentCatalog }, parentCatalogView);
};

$(document).ready(function() {
    catalog.listVM = avalon.define({
        $id: "catalog-List",
        catalogs: []
    });

    catalog.editVM = avalon.define({
        $id: "catalog-Edit",
        catalog: {},
        parentCatalogs: []
    });

    $("#submitCatalog-button").click(
        function() {
            var catalogID = $("#catalog-Edit .catalog-id").val();
            var catalogName = $("#catalog-Edit .catalog-name").val();
            var catalogDescription = $("#catalog-Edit .catalog-description").val();
            var selectParent = "";
            $("#catalog-Edit .catalog-parent .item:checked").each(
                function() {
                    var id = $(this).val();
                    if (selectParent.length > 0) {
                        selectParent += ",";
                    }
                    selectParent += id;
                }
            );

            var callBack = function(errCode, catalogItem) {
                if (errCode != 0) {
                    return;
                }

                $("#catalog-Edit .catalog-name").val("");
                $("#catalog-Edit .catalog-description").val("");
                $("#catalog-Edit .catalog-parent .item").prop("checked", false);

                catalog.loadData(function() {
                    catalog.refreshCatalogListView(catalog.catalogs, catalog.catalogs);

                    catalog.refreshCatalogEditView(catalog.curCatalog, catalog.catalogs);
                });
            };

            catalog.saveCatalogAction(catalogName, catalogDescription, selectParent, callBack);
        }
    );

    catalog.loadData(function() {
        catalog.refreshCatalogListView(catalog.catalogs, catalog.catalogs);

        catalog.refreshCatalogEditView(catalog.curCatalog, catalog.catalogs);
    });
});