
var content = {
	moduleList:{},
	currentModule:{},
};

$(document).ready(function() {

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
	nameTd.innerHTML = block.Name
	tr.appendChild(nameTd);
	
	var numberTd = document.createElement("td");
	if (block.Items) {
		numberTd.innerHTML = block.Items.length;		
	} else {
		numberTd.innerHTML = 0;
	}
	tr.appendChild(numberTd);
	
	tr.setAttribute("onclick","content.editModuleBlock('" + block.Id + "'); return false;" );
	
	return tr;
};

content.editModuleBlock = function(block) {
	
};

content.refreshModuleView = function() {
	$("#block-list table tbody tr").remove();
	
	if (!content.currentModule) {
		return;
	}
	
	if (content.currentModule.Blocks) {
		console.log(content.currentModule);
		
		for (var ii =0; ii < content.currentModule.Blocks.length; ++ii) {
			var block = content.currentModule.Blocks[ii];
			var trContent = content.constructBlockItem(block);
			$("#block-list table tbody").append(trContent);
		}
		
		$("#block-list table tbody tr:even").addClass("alt-row");
	}
	
	//$("#block-list table tbody tr:even").addClass("alt-row");
};




