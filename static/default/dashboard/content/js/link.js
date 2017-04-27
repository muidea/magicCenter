var link = {};

link.constructLinkListlView = function(linkList, catalogList) {
    var linkListView = new Array();
    var offset = 0;
    for (var i = 0; i < linkList.length; ++i) {
        var curLink = linkList[i];
        var catalogNames = "";
        if (curLink.Catalog) {
            for (var idx = 0; idx < catalogList.length; ++idx) {
                var curCatalog = catalogList[idx];
                for (var j = 0; j < curLink.Catalog.length; j++) {
                    var val = curLink.Catalog[j];
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
            ID: curLink.ID,
            Name: curLink.Name,
            Catalog: catalogNames,
            CreateDate: curLink.CreateDate
        };

        linkListView[offset++] = view;
    }

    return linkListView;
};

link.constructLinkEditView = function(catalogList) {
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

link.updateListLinkVM = function(linkList) {
    link.listVM.links = linkList;
}

link.updateEditLinkVM = function(linkView, catalogView) {
    link.editVM.link = linkView;
    link.editVM.catalogs = catalogView;
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

link.saveLinkAction = function(name, url, logo, catalogs, callBack) {
    $.ajax({
        type: "POST",
        url: "/content/link/",
        data: { "link-name": name, "link-url": url, "link-logo": logo, "link-catalog": catalogs },
        dataType: "json",
        success: function(data) {
            if (callBack != null) {
                callBack(data.ErrCode, data.Link);
            }
        }
    });
};

link.loadData = function(callBack) {
    var getAllCatalogsCallBack = function(errCode, catalogList) {
        if (errCode != 0) {
            return;
        }

        link.curLink = { ID: -1, Name: "", URL: "", Logo: "", Catalog: [] };
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

link.refreshLinkEditView = function(curLink, catalogList) {
    var linkView = {
        ID: curLink.ID,
        Name: curLink.Name,
        URL: curLink.URL,
        Logo: curLink.Logo,
        Catalog: curLink.Catalog
    };

    var catalogView = link.constructLinkEditView(catalogList);
    link.updateEditLinkVM(linkView, catalogView)
};

$(document).ready(function() {
    link.listVM = avalon.define({
        $id: "link-List",
        links: []
    });

    link.editVM = avalon.define({
        $id: "link-Edit",
        link: {},
        catalogs: []
    });

    $("#link-Edit .submit").click(
        function() {
            var linkID = $("#link-Edit .link-id").val();
            var linkName = $("#link-Edit .link-name").val();
            var linkURL = $("#link-Edit .link-url").val();
            var linkLogo = $("#link-Edit .link-logo").val();
            var selectCatalog = "";
            $("#link-Edit .link-catalog .catalog-item:checked").each(
                function() {
                    var id = $(this).val();
                    if (selectCatalog.length > 0) {
                        selectCatalog += ",";
                    }
                    selectCatalog += id;
                }
            );

            var callBack = function(errCode, catalogItem) {
                if (errCode != 0) {
                    return;
                }

                $("#link-Edit .link-name").val("");
                $("#link-Edit .link-url").val("");
                $("#link-Edit .link-logo").val("");
                $("#link-Edit .link-catalog .catalog-item").prop("checked", false);

                link.loadData(function() {
                    link.refreshLinkListView(link.links, link.catalogs);
                    link.refreshLinkEditView(link.curLink, link.catalogs);
                });
            };
            link.saveLinkAction(linkName, linkURL, linkLogo, selectCatalog, callBack);
        }
    );

    link.loadData(function() {
        link.refreshLinkListView(link.links, link.catalogs);
        link.refreshLinkEditView(link.curLink, link.catalogs);
    });
});