var index = {
		view:{}
};

$(document).ready(function() {
	var pageView = index.view;
	console.log(pageView);
	var menu = magic.findBlock(pageView, "nav");
	var menuView = magic.listView(menu);
	$("#menu ul").remove();
	$("#menu").append(menuView);
	
	var catalog = magic.findBlock(pageView, "list");
	var catalogView = magic.listView(catalog);
	$("#sidebar .catalog ul").remove();
	$("#sidebar .catalog").append(catalogView);

	var link = magic.findBlock(pageView, "link");
	var linkView = magic.listView(link);
	$("#sidebar .link ul").remove();
	$("#sidebar .link").append(linkView);
	
	$("#content .post").remove();
	var view = magic.findBlock(pageView, "view");
	if (view && view.Items) {
		for(var ii=0; ii < view.Items.length; ++ii) {
			var item = view.Items[ii];
			var post = magic.findContent(pageView, item.Id);
			var postView = magic.contentView(post);
			
			$("#content").append(postView);
		}
	}
});	

