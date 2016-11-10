var index = {
    view: {}
};

$(document).ready(function() {
    var pageView = index.view;
    console.log(pageView);

    $("#menu ul").remove();
    var menu = cms.findBlock(pageView, "nav");
    if (menu) {
        var menuView = cms.listView(menu);
        $("#menu").append(menuView);
    }

    $("#sidebar .catalog ul").remove();
    var catalog = cms.findBlock(pageView, "list");
    if (catalog) {
        var catalogView = cms.listView(catalog);
        $("#sidebar .catalog").append(catalogView);
    }

    $("#sidebar .link ul").remove();
    var link = cms.findBlock(pageView, "link");
    if (link) {
        var linkView = cms.listView(link);
        $("#sidebar .link").append(linkView);
    }

    $("#content .post").remove();
    var view = cms.findBlock(pageView, "post");
    if (view) {
        if (view && view.Items) {
            for (var ii = 0; ii < view.Items.length; ++ii) {
                var item = view.Items[ii];
                var post = cms.findPost(pageView, item.Id);
                var postView = cms.postView(post);
                $("#content").append(postView);
            }
        }
    }
});