var view = {
    view: {}
};

$(document).ready(function() {
    var pageView = view.view;
    console.log(pageView);

    var content = pageView.Content;
    if (content) {
        $("#content div").remove();
        var contentView = magic.contentView(content);
        $("#content").append(contentView);
    } else {
        $("#content").remove();
    }

    var menu = magic.findBlock(pageView, "nav");
    if (menu) {
        $("#menu ul").remove();
        var menuView = magic.listView(menu);
        $("#menu").append(menuView);
    } else {
        $("#menu").remove();
    }

    var catalog = magic.findBlock(pageView, "list");
    if (catalog) {
        $("#sidebar .catalog ul").remove();
        var catalogView = magic.listView(catalog);
        $("#sidebar .catalog").append(catalogView);
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
});