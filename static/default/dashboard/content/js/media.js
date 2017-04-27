var media = {};

media.constructMediaListlView = function(mediaList, catalogList) {
    var mediaListView = new Array();
    var offset = 0;
    for (var i = 0; i < mediaList.length; ++i) {
        var curMedia = mediaList[i];
        var catalogNames = "";
        if (curMedia.Catalog) {
            for (var idx = 0; idx < catalogList.length; ++idx) {
                var curCatalog = catalogList[idx];
                for (var j = 0; j < curMedia.Catalog.length; j++) {
                    var val = curMedia.Catalog[j];
                    if (curCatalog.ID == val) {
                        if (catalogNames.length > 0) {
                            catalogNames += ", ";
                        }
                        catalogNames += curCatalog.Name;
                    }
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

media.constructMediaEditView = function(catalogList) {
    var catalogListView = new Array();
    var ii = 0;

    for (var idx = 0; idx < catalogList.length; ++idx) {
        var curCatalog = catalogList[idx];
        var view = {
            ID: curCatalog.ID,
            Name: curCatalog.Name
        }

        catalogListView[ii++] = view;
    }

    return catalogListView;
}

media.updateListMediaVM = function(mediaList) {
    media.listVM.medias = mediaList;
}

media.updateEditMediaVM = function(mediaView, catalogView) {
    media.editVM.media = mediaView;
    media.editVM.catalogs = catalogView;
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

        media.curMedia = { ID: -1, Name: "", Catalog: [] };
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

media.refreshMediaEditView = function(curMedia, catalogList) {
    var mediaView = {
        ID: curMedia.ID,
        Name: curMedia.Name,
        Catalog: curMedia.Catalog
    };

    var catalogView = media.constructMediaEditView(catalogList);
    media.updateEditMediaVM(mediaView, catalogView)
};

$(document).ready(function() {
    media.listVM = avalon.define({
        $id: "media-List",
        medias: []
    });

    media.editVM = avalon.define({
        $id: "media-Edit",
        media: {},
        catalogs: []
    });

    $("#selectMedia-button").click(
        function() {
            media.statusMediasAction(
                selectMediaList,
                unSelectMediaList,
                function(errCode) {
                    if (errCode != 0) {
                        return;
                    }

                    media.loadData(function() {
                        media.refreshMediaListView(media.medias, media.catalogs);

                        media.refreshMediaEditView(media.curMedia, media.catalogs);
                    })
                });
        }
    );

    media.loadData(function() {
        media.refreshMediaListView(media.medias, media.catalogs);

        media.refreshMediaEditView(media.curMedia, media.catalogs);
    })
});