var view = {
		view:{}
};

$(document).ready(function() {
	var pageView = view.view;
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
});	

