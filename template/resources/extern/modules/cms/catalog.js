    var catalog = {
        view: {}
    };

    $(document).ready(function() {
        var pageView = catalog.view;
        var menu = magic.findBlock(pageView, "nav");
        if (menu) {
            $("#menu ul").remove();
            var menuView = magic.listView(menu);
            $("#menu").append(menuView);
        } else {
            $("#menu").remove();
        }

        var list = magic.findBlock(pageView, "list");
        if (list) {
            $("#sidebar .catalog ul").remove();
            var listView = magic.listView(list);
            $("#sidebar .catalog").append(listView);
        } else {
            $("#sidebar .catalog").remove();
        }

        var link = magic.findBlock(pageView, "link");
        if (link) {
            $("#sidebar .link ul").remove();
            var linkView = magic.listView(link);
            $("#sidebar .link").append(linkView);
        } else {
            $("#sidebar .link").remove();
        }

        var catalogs = pageView.Catalogs;
        if (catalogs) {
            $("#content").children().remove();
            var catalogsView = magic.catalogView(catalogs);
            $("#content").append(catalogsView);
        }
    });