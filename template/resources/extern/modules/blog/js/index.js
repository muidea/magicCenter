var index = {
    view: {}
};

$(document).ready(function() {
    var pageView = index.view;
    console.log(pageView);

    $("#menu ul").remove();
    var menu = magic.findBlock(pageView, "nav");
    if (menu) {
        var menuView = magic.listView(menu);
        $("#menu").append(menuView);
    }

    $("#sidebar .catalog ul").remove();
    var catalog = magic.findBlock(pageView, "list");
    if (catalog) {
        var catalogView = magic.listView(catalog);
        $("#sidebar .catalog").append(catalogView);
    }

    $("#sidebar .link ul").remove();
    var link = magic.findBlock(pageView, "link");
    if (link) {
        var linkView = magic.listView(link);
        $("#sidebar .link").append(linkView);
    }

    $("#content .post").remove();
    var view = magic.findBlock(pageView, "post");
    if (view) {
        if (view && view.Items) {
            for (var ii = 0; ii < view.Items.length; ++ii) {
                var item = view.Items[ii];
                var post = magic.findPost(pageView, item.Id);
                var postView = magic.postView(post);
                $("#content").append(postView);
            }
        }
    }
});