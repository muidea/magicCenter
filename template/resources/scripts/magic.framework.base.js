
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

magic.findContent = function(view, id) {
	var val = null;
	if (view.Contents) {
		for (var ii =0; ii < view.Contents.length; ++ii) {
			var content = view.Contents[ii];
			if (content.Id == id) {
				val = content;
				break;
			}			
		}
	}
	
	return val;
};

magic.contentView = function(view) {
	var div = document.createElement("div");
	div.setAttribute("class","post");
	
	// meta
	var meta = document.createElement("p");
	meta.setAttribute("class","meta");
	
	var date = document.createElement("span");
	date.setAttribute("class","date");
	date.innerHTML = view.CreateDate;
	meta.appendChild(date);
	
	var author = document.createElement("span");
	author.setAttribute("class","author");
	author.innerHTML = view.Author.Name;
	meta.appendChild(author);	
	
	div.appendChild(meta);
	
	// title
	var title = document.createElement("h2");
	title.setAttribute("class","title");
	title.innerHTML = view.Title;
	
	div.appendChild(title);
	
	// entry
	var entry = document.createElement("div");
	entry.setAttribute("class","entry");
	entry.innerHTML = view.Content;
	
	div.appendChild(entry);
	
	// readmore
	var readMore = document.createElement("div");
	var a = document.createElement("a");
	a.innerHTML = "阅读全文。。.";
	a.setAttribute("href", view.Url);
	readMore.appendChild(a);
	
	div.appendChild(readMore);
	
	return div;
};


