
var magic = {};

/*
 * 查找View中指定的Block
 * 
 * */
magic.findBlock = function(view, blockName) {
	var val = null;
	if (view.Blocks) {
		for (var ii =0; ii < view.Blocks.length; ++ii) {
			var block = view.Blocks[ii];
			if (block.Tag == blockName) {
				val = block;
				break;
			}
		}
	}
	
	return val;
};

magic.listView = function(view) {
	var ul = document.createElement("ul");
	if (view.Items) {
		for (var ii =0; ii < view.Items.length; ++ii) {
			var item = view.Items[ii];
			
			var li = document.createElement("li");
			var a = document.createElement("a");
			a.innerHTML = item.Name;
			a.setAttribute("href", item.Url);
			li.appendChild(a);
			ul.appendChild(li);
		}
	}
	
	return ul;
};

