var media = {};

media.constructMediaListlView = function(mediaList, catalogList) {
    var mediaListView = new Array();
    var offset = 0;
    for (var i = 0; i < mediaList.length; ++i) {
        var curMedia = mediaList[i];
        var catalogNames = "";
        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curCatalog = catalogList[idx];
            for (var j = 0; j < curMedia.Catalog; j++) {
                var val = curMedia.Catalog[j];
                if (curCatalog.ID == val) {
                    if (catalogNames.length > 0) {
                        catalogNames += ", ";
                    }
                    catalogNames += curCatalog.Name;
                }
            }
        }
        var view = {
            ID: curMedia.ID,
            Name: curMedia.Name,
            Catalog: catalogNames,
            CreateDate: curMedia.CreateDate
        };

        mediaListView[offset++] = view;
    }

    return mediaListView;
};

media.constructMediaEditView = function(mediaList, catalogList) {
    var mediaListView = new Array();
    var ii = 0;
    for (var mediaIdx = 0; mediaIdx < mediaList.length; ++mediaIdx) {
        var curMedia = mediaList[mediaIdx];

        for (var idx = 0; idx < catalogList.length; ++idx) {
            var curModule = catalogList[idx];

            if (curMedia.Module == curModule.ID) {
                var view = {
                    ID: curMedia.ID,
                    URL: curMedia.URL,
                    Method: curMedia.Method,
                    Status: curMedia.Status,
                    Module: curModule.Name,
                    ModuleID: curModule.ID
                }

                mediaListView[ii++] = view;
            }
        }
    }

    return mediaListView;
}

media.updateListMediaVM = function(mediaList) {
    media.listVM.medias = mediaList;
}

media.updateEditModuleVM = function(catalogList) {
    media.editVM.modules = catalogList;
};

media.updateEditMediaVM = function(mediaList) {
    media.editVM.medias = mediaList;

    // 将已经enable的media设置上checked标示
    for (var offset = 0; offset < media.editVM.medias.length; ++offset) {
        var curMedia = media.editVM.medias[offset];
        if (curMedia.Status > 0) {
            $("#selectMedia-List .media_" + curMedia.ID).prop("checked", true);
        }
    }
};

// 加载全部的Media
media.getAllMediasAction = function(callBack) {
    $.ajax({
        type: "GET",
        url: "/content/media/",
        data: {},
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Media);
            }
        }
    });
};

// 加载全部Catalog
media.getAllCatalogsAction = function(callBack) {
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

media.loadData = function(callBack) {
    var getAllCatalogsCallBack = function(errCode, catalogList) {
        if (errCode != 0) {
            return;
        }

        media.catalogs = catalogList;
        if (callBack != null) {
            callBack()
        }
    };

    var getAllMediasCallBack = function(errCode, mediaList) {
        if (errCode != 0) {
            return;
        }

        media.medias = mediaList;
        media.getAllCatalogsAction(getAllCatalogsCallBack);
    };

    // 加载完成
    media.getAllMediasAction(getAllMediasCallBack);
}

media.refreshMediaListView = function(mediaList, catalogList) {
    var mediasView = media.constructMediaListlView(mediaList, catalogList);
    media.updateListMediaVM(mediasView);
};

media.refreshMediaEditView = function(media, catalogList) {};

$(document).ready(function() {
    media.listVM = avalon.define({
        $id: "media-List",
        medias: []
    });

    media.editVM = avalon.define({
        $id: "media-Edit",
        modules: [],
        medias: []
    });

    $('#moduleListModal').on('show.bs.modal', function(e) {
        media.updateEditModuleVM(media.modules);

        $("#moduleListModal .module").prop("checked", false);
    });

    $('#moduleListModal').on('hidden.bs.modal', function(e) {
        var selectModuleArray = new Array()
        var offset = 0;
        $("#moduleListModal .module:checked").each(
            function() {
                var id = $(this).val();
                for (var idx = 0; idx < media.modules.length; idx++) {
                    var curModule = media.modules[idx];
                    if (curModule.ID == id) {
                        selectModuleArray[offset++] = curModule;
                    }
                }
            }
        );
        media.refreshMediaEditView(media.medias, selectModuleArray);
    });

    $("#selectMedia-button").click(
        function() {
            var selectMediaList = new Array();
            var offset = 0;
            $("#selectMedia-List .media_status_0:checked").each(
                function() {
                    var id = $(this).val();
                    selectMediaList[offset++] = id;
                }
            );

            var unSelectMediaList = new Array();
            offset = 0;
            $("#selectMedia-List .media_status_1:not(:checked)").each(
                function() {
                    var id = $(this).val();
                    unSelectMediaList[offset++] = id;
                }
            );

            media.statusMediasAction(
                selectMediaList,
                unSelectMediaList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    var selectModuleArray = new Array()
                    var offset = 0;
                    $("#moduleListModal .module:checked").each(
                        function() {
                            var id = $(this).val();
                            for (var idx = 0; idx < media.modules.length; idx++) {
                                var curModule = media.modules[idx];
                                if (curModule.ID == id) {
                                    selectModuleArray[offset++] = curModule;
                                }
                            }
                        }
                    );

                    media.loadData(function() {
                        media.refreshMediaListView(media.medias, media.catalogs);

                        media.refreshMediaEditView(media.medias, selectModuleArray);
                    })
                });
        }
    );

    media.loadData(function() {
        media.refreshMediaListView(media.medias, media.catalogs);
    })
});