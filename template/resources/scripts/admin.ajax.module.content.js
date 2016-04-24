
var content = {
	moduleList:{},
	articleList:{},
	catalogList:{},
	linkList:{},
	currentModule:{},
	currentBlock:-1,
};

$(document).ready(function() {
	
	$("#block-form .form .button").click(
			function() {
				var articleList = "";
				var catalogList = "";
				var linkList = "";
				var block_id = $("#block-form .block-id").val();
				var blockArray = $("#block-list table tbody tr td :checkbox:checked");
				var articleArray = $("#block-form .article-list table tbody tr td :checkbox:checked");
				var catalogArray = $("#block-form .catalog-list table tbody tr td :checkbox:checked");
				var linkArray = $("#block-form .link-list table tbody tr td :checkbox:checked");
				
				if (blockArray.length == 0) {
	        		$("#module-content div.error div").html("请选择一个功能块");
	        		$("#module-content div.error").show();
	        		return;
				}
				
				for (var ii =0; ii < articleArray.length; ++ii) {
					var chk = articleArray[ii];
					articleList += $(chk).attr("value");
					articleList += ",";
				}
				
				for (var ii =0; ii < catalogArray.length; ++ii) {
					var chk = catalogArray[ii];
					catalogList += $(chk).attr("value");
					catalogList += ",";
				}

				for (var ii =0; ii < linkArray.length; ++ii) {
					var chk = linkArray[ii];
					linkList += $(chk).attr("value");
					linkList += ",";
				}
				
				$.post("/admin/system/ajaxBlockItem/", {
					'module-id':content.currentModule.Id,
					"block-id":block_id,
					"article-list":articleList,
					"catalog-list":catalogList,
					"link-list":linkList
				}, function(result) {

					content.currentModule = result.Module;
					
					$("#module-content div.notification").hide();
		        	if (result.ErrCode > 0) {
		        		$("#module-content div.error div").html(result.Reason);
		        		$("#module-content div.error").show();
		        	} else {
		        		$("#module-content div.success div").html(result.Reason);
		        		$("#module-content div.success").show();
			        	content.refreshModuleView();
		        	}
		        			        	
				}, "json");
			}
		);
});

content.initialize = function() {
	content.refreshModuleView();
	
	content.fillModuleView();
};

content.fillModuleView = function() {
		
	$("#module-list div.notification").hide();
	$("#module-list table tbody tr").remove();
	for (var ii =0; ii < content.moduleList.length; ++ii) {
		var info = content.moduleList[ii];
		var trContent = content.constructModuleItem(info);		
		$("#module-list table tbody").append(trContent);
	}
	
	$("#module-list table tbody tr:even").addClass("alt-row");	
};

content.constructModuleItem = function(mod) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","module");
		
	var nameTd = document.createElement("td");	
	nameTd.innerHTML = mod.Name;
	tr.appendChild(nameTd);

	var descriptionTd = document.createElement("td");
	descriptionTd.innerHTML = mod.Description
	tr.appendChild(descriptionTd);
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","button");
	editLink.setAttribute("href","#");
	editLink.setAttribute("onclick","content.maintainContent('/admin/system/queryModuleContent/?id=" + mod.Id + "'); return false;" );
	editLink.innerHTML = "内容管理";
	editTd.appendChild(editLink);
	tr.appendChild(editTd);	
	return tr;
};

content.maintainContent = function(maintainUrl) {
	$.get(maintainUrl, {
	}, function(result) {
		
	if (result.ErrCode > 0) {
		$("#module-List div.error div").html(result.Reason);
		$("#module-List div.error").show();
		return
	}

	content.currentModule = result.Module;
	content.articleList = result.Articles;
	content.catalogList = result.Catalogs;
	content.linkList = result.Links;
	
	$("#content .content-box-header .content-box-tabs li a").removeClass('current');
	$("#content .content-box-header .content-box-tabs li a.module-Content-tab").addClass('current');
	$("#module-content").siblings().hide();
	
	content.refreshModuleView();
	
	$("#module-content").show();	
	}, "json");	
};

content.constructBlockItem = function(block) {
	
	var tr = document.createElement("tr");
	tr.setAttribute("class","block");
	
	var nameTd = document.createElement("td");
	var chkBox = document.createElement("input");
	chkBox.setAttribute("type","checkbox");
	chkBox.setAttribute("value",block.Id);
	nameTd.appendChild(chkBox);
	var label = document.createElement("span");
	label.innerHTML = block.Name;
	nameTd.appendChild(label);
	tr.appendChild(nameTd);
	
	var numberTd = document.createElement("td");
	var numVal = 0;
	if (block.Article) {
		numVal += block.Article.length;
	}
	if (block.Catalog) {
		numVal += block.Catalog.length;
	}
	if (block.Link) {
		numVal += block.Link.length;
	}	
	numberTd.innerHTML = numVal;
	
	tr.appendChild(numberTd);
	
	tr.setAttribute("onclick","content.editModuleBlock(" + block.Id + "); return false;" );
	
	return tr;
};

content.editModuleBlock = function(block) {
	content.currentBlock = block;
	
	$("#block-list table tbody tr td input").prop("checked", false);
	$("#block-form table tbody tr td input").prop("checked", false);
	
	$("#block-list table tbody tr td input").filter("[value="+ block +"]").prop("checked", true);
	$("#block-form .block-id").val(block);	
	
	content.refreshModuleView();
};

content.constructArticleItem = function(article) {
	
	var tr = document.createElement("tr");
	
	var nameTd = document.createElement("td");
	var chkBox = document.createElement("input");
	chkBox.setAttribute("type","checkbox");
	chkBox.setAttribute("value",article.Id);
	chkBox.setAttribute("class","article");
	nameTd.appendChild(chkBox);
	var label = document.createElement("span");
	label.innerHTML = article.Title;
	nameTd.appendChild(label);
	tr.appendChild(nameTd);
	
	return tr;
};

content.constructCatalogItem = function(catalog) {
	
	var tr = document.createElement("tr");
	
	var nameTd = document.createElement("td");
	var chkBox = document.createElement("input");
	chkBox.setAttribute("type","checkbox");
	chkBox.setAttribute("value",catalog.Id);
	chkBox.setAttribute("class","catalog");
	nameTd.appendChild(chkBox);
	var label = document.createElement("span");
	label.innerHTML = catalog.Name;
	nameTd.appendChild(label);
	tr.appendChild(nameTd);
	
	return tr;
};

content.constructLinkItem = function(link) {
	
	var tr = document.createElement("tr");
	
	var nameTd = document.createElement("td");
	var chkBox = document.createElement("input");
	chkBox.setAttribute("type","checkbox");
	chkBox.setAttribute("value",link.Id);
	chkBox.setAttribute("class","link");
	nameTd.appendChild(chkBox);
	var label = document.createElement("span");
	label.innerHTML = link.Name;
	nameTd.appendChild(label);
	tr.appendChild(nameTd);
	
	return tr;
};

content.refreshModuleView = function() {
	$("#block-list table tbody tr").remove();
	$("#block-form .article-list table tbody tr").remove();
	$("#block-form .catalog-list table tbody tr").remove();
	$("#block-form .link-list table tbody tr").remove();
	
	if (!content.currentModule) {
		return;
	}
	
	var currentBlock = null;
	if (content.currentModule.Blocks) {
		
		for (var ii =0; ii < content.currentModule.Blocks.length; ++ii) {
			var block = content.currentModule.Blocks[ii];
			
			if (block.Id == content.currentBlock) {
				currentBlock = block;
			}
			
			var trContent = content.constructBlockItem(block);
			$("#block-list table tbody").append(trContent);
		}
		
		$("#block-list table tbody tr:even").addClass("alt-row");
		
		$("#block-list table tbody tr td input").filter("[value="+ content.currentBlock +"]").prop("checked", true);
	}
	
	if (content.articleList) {
		for (var ii =0; ii < content.articleList.length; ++ii) {
			var article = content.articleList[ii];
			var trContent = content.constructArticleItem(article);
			$("#block-form .article-list table tbody").append(trContent);
		}
		
		if(currentBlock) {
			for (var ii=0; ii < currentBlock.Article.length; ++ii) {
				var item = currentBlock.Article[ii];
				$("#block-form .article-list table tbody tr td input").filter("[value="+ item.Rid +"]").prop("checked", true);
			}
		}
		
		$("#block-form .article-list table tbody tr:odd").addClass("alt-row");
	}
	
	if (content.catalogList) {
		for (var ii =0; ii < content.catalogList.length; ++ii) {
			var catalog = content.catalogList[ii];
			var trContent = content.constructCatalogItem(catalog);
			$("#block-form .catalog-list table tbody").append(trContent);
		}
		
		if(currentBlock) {
			for (var ii=0; ii < currentBlock.Catalog.length; ++ii) {
				var item = currentBlock.Catalog[ii];
				$("#block-form .catalog-list table tbody tr td input").filter("[value="+ item.Rid +"]").prop("checked", true);
			}
		}
		
		$("#block-form .catalog-list table tbody tr:odd").addClass("alt-row");
	}
	
	if (content.linkList) {
		for (var ii =0; ii < content.linkList.length; ++ii) {
			var link = content.linkList[ii];
			var trContent = content.constructLinkItem(link);
			$("#block-form .link-list table tbody").append(trContent);
		}
		
		if(currentBlock) {
			for (var ii=0; ii < currentBlock.Link.length; ++ii) {
				var item = currentBlock.Link[ii];
				$("#block-form .link-list table tbody tr td input").filter("[value="+ item.Rid +"]").prop("checked", true);
			}
		}
		
		$("#block-form .link-list table tbody tr:odd").addClass("alt-row");
	}
	
	//$("#block-list table tbody tr:even").addClass("alt-row");
};




