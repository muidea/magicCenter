
var module = {
	moduleList:{},
	defaultModule:'',
	currentModule:{},
	currentPage:{}
};

$(document).ready(function() {
	
	$("#module-List .button").click(
		function() {
			var enableList = "";
			var defaultModule = "";
			var radioArray = $("#module-List table tbody tr td :radio:checked");
			for (var ii =0; ii < radioArray.length; ++ii) {
				var radio = radioArray[ii];
				if ($(radio).val() == 1) {
					enableList += $(radio).attr("name");
					enableList += ",";
				}				
			}
			var defaultArray = $("#module-List .module input:checkbox:checked");
			if (defaultArray.length > 0) {
				var checkBox = defaultArray[0];
				defaultModule = $(checkBox).attr("name");
			}
			
			$.post("/admin/system/applyModuleSetting/", {
				"module-enableList":enableList,
				"module-defaultModule":defaultModule
			}, function(result) {

				module.moduleList = result.Modules;
				module.defaultModule = result.DefaultModule;
								
				$("#module-List label").addClass("hidden")
	        	if (result.ErrCode > 0) {
	        		$("#module-List .danger").html(result.Reason);
	        		$("#module-List .danger").removeClass("hidden");
	        	} else {
	        		$("#module-List .success").html(result.Reason);
	        		$("#module-List .success").removeClass("hidden");
	        	}
				
	        	//module.fillModuleView();	        	
			}, "json");
		}
	);
	
    // 绑定表单提交事件处理器
    $('#module-Maintain .block .block-Form').submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //return false;
        } 
        // post-submit callback
        function showResponse(result) {
        	$("#module-maintain div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#module-maintain div.error div").html(result.Reason);
        		$("#module-maintain div.error").show();        		
        	} else {
        		$("#module-maintain .block .block-Form .module-block").val("");
        		$("#module-maintain div.success div").html(result.Reason);
        		$("#module-maintain div.success").show();
        		
        		module.currentModule = result.Module;
        		module.refreshModuleView();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#module-maintain .block .block-Form .block-name").parent().find("span").remove();
        	var name = $("#module-maintain .block .block-Form .block-name").val();
        	if (name.length == 0) {
        		$("#module-maintain .block .block-Form .block-name").parent().append("<span class=\"input-notification error png_bg\">请输入功能块名</span>");
        		result = false;
        	}
        	        	
        	return result;
        }
        
        if (!validate()) {
        	return false;
        }
        
        //提交表单
        $(this).ajaxSubmit(options);	
    	
        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });	
    
    // 绑定表单提交事件处理器
    $('#module-maintain .page .page-Form').submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //return false;
        } 
        // post-submit callback
        function showResponse(result) {
        	$("#module-maintain div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#module-maintain div.error div").html(result.Reason);
        		$("#module-maintain div.error").show();        		
        	} else {
        		$("#module-maintain div.success div").html(result.Reason);
        		$("#module-maintain div.success").show();
        		
        		console.log(module);

        		module.currentModule = result.Module;
        		if (module.currentPage) {
        			for (var ii=0; ii < module.currentModule.Pages.length; ++ii) {
        				var page = module.currentModule.Pages[ii];
        				if (page.Url == module.currentPage.Url) {
        					module.currentPage = page;
        					break;
        				}
        			}
        		}
        		
        		module.refreshModuleView();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#module-maintain .page .page-Form .page-url").parent().find("span").remove();
        	var url = $("#module-maintain .page .page-Form .page-url").val();
        	if (url.length == 0) {
        		$("#module-maintain .page .page-Form .page-url").parent().append("<span class=\"input-notification error png_bg\">路由不能为空</span>");
        		result = false;
        	}
        	        	
        	return result;
        }
        
        if (!validate()) {
        	return false;
        }
        
        //提交表单
        $(this).ajaxSubmit(options);	
    	
        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });	    
});

module.getModuleListView = function() {
	return $("#module-List table")	
}

module.getModuleBlockListView = function() {
	return  $("#module-Maintain .block table")
}

module.getModulePageListView = function() {
	return $("#module-Maintain .page table")
}

module.initialize = function() {
	module.fillModuleListView();
};


module.fillModuleListView = function() {
	var moduleListView = module.getModuleListView();
	$(moduleListView).find("tbody tr").remove();
	for (var ii =0; ii < module.moduleList.length; ++ii) {
		var info = module.moduleList[ii];
		var trContent = module.constructModuleItem(info);

		$(moduleListView).find("tbody").append(trContent);
	}
};

module.fillModuleMaintainView = function() {
	var blockListView = module.getModuleBlockListView();
	$(blockListView).find("tbody tr").remove();
	if (module.currentModule) {
		for (var ii =0; ii < module.currentModule.Blocks.length; ++ii) {
			var block = module.currentModule.Blocks[ii];
			var trContent = module.constructBlockItem(block);
			$(blockListView).find("tbody").append(trContent);
		}
	}

	var pageListView = module.getModulePageListView();
	$(pageListView).find("tbody tr").remove();
	if (module.currentModule.Pages) {
		for (var ii =0; ii < module.currentModule.Pages.length; ++ii) {
			var page = module.currentModule.Pages[ii];
			var trContent = module.constructPageItem(page);
			$(pageListView).find("tbody").append(trContent);
		}			
	}	
}

module.constructModuleItem = function(mod) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","module");
	
	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("href","#");
	nameLink.setAttribute("onclick","module.maintainModule('/admin/system/queryModuleDetail/?id=" + mod.Id + "'); return false;" );
	nameLink.innerHTML = mod.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);

	var descriptionTd = document.createElement("td");
	descriptionTd.innerHTML = mod.Description
	tr.appendChild(descriptionTd);
	
	var editTd = document.createElement("td");
	var radioGroup = document.createElement("radiobox");
	var enable_radio = document.createElement("input");
	enable_radio.setAttribute("type","radio");
	enable_radio.setAttribute("name",mod.Id);
	enable_radio.setAttribute("value","1");
	radioGroup.appendChild(enable_radio);	
	var enable_span = document.createElement("span");
	enable_span.innerHTML ="启用";
	radioGroup.appendChild(enable_span);
	
	var disable_radio = document.createElement("input");
	disable_radio.setAttribute("type","radio");
	disable_radio.setAttribute("name",mod.Id);
	disable_radio.setAttribute("value","0");
	radioGroup.appendChild(disable_radio);
	if (mod.EnableFlag) {
		enable_radio.checked = true;
		disable_radio.checked = false;
	} else {
		enable_radio.checked = false;
		disable_radio.checked = true;		
	}
	
	var disable_span = document.createElement("span");
	disable_span.innerHTML ="禁用";
	radioGroup.appendChild(disable_span);
	
	editTd.appendChild(radioGroup);
		
	var checkGroup = document.createElement("checkbox");
	var default_check = document.createElement("input");
	default_check.setAttribute("type","checkbox");
	default_check.setAttribute("name",mod.Id);
	default_check.setAttribute("onclick","module.selectDefaultModule('"+ mod.Id +"');");	
	checkGroup.appendChild(default_check);
	if (module.defaultModule == mod.Id) {
		default_check.checked = true;
	} else {
		default_check.checked = false;
	}
	
	
	var default_span = document.createElement("span");
	default_span.innerHTML ="设为默认 ";
	checkGroup.appendChild(default_span);
	
	editTd.appendChild(checkGroup);
	
	tr.appendChild(editTd);	
	return tr;
};

module.selectDefaultModule = function(defaultModule) {
	$("#module-list .module input:checkbox").prop("checked", false);
	$("#module-list .module input:checkbox[name='"+ defaultModule +"']").prop("checked", true);
};

module.maintainModule = function(maintainUrl) {
	$.get(maintainUrl, {
	}, function(result) {
		if (result.ErrCode > 0) {
			return
		}

		module.currentModule = result.Module;
		module.fillModuleMaintainView();		
		$("#module-content .content-header .nav .module-Maintain").find("a").trigger("click");	
	}, "json");	
};

module.constructBlockItem = function(block) {
	
	var tr = document.createElement("tr");
	tr.setAttribute("class","block");
	
	var nameTd = document.createElement("td");
	nameTd.innerHTML = block.Name
	tr.appendChild(nameTd);
	
	var styleTd = document.createElement("td");
	if (block.Style == 0) {
		styleTd.innerHTML = "链接";
	} else {
		styleTd.innerHTML = "内容";
	}
	tr.appendChild(styleTd);
	
	var editTd = document.createElement("td");
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteBlock" );
	deleteLink.setAttribute("onclick","module.deleteBlock('/admin/system/deleteModuleBlock/?id=" + block.Id + "&owner="+ module.currentModule.Id +"'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	tr.appendChild(editTd);	
	return tr;
};

module.constructPageItem = function(page) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","block");
	
	var nameTd = document.createElement("td");
	nameTd.innerHTML = page.Url
	tr.appendChild(nameTd);
	
	var blocks = "";
	var blocksTd = document.createElement("td");
	if (page.Blocks) {
		for (var ii =0; ii < page.Blocks.length;) {
			var block = page.Blocks[ii++];
			blocks += block.Name;
			if (ii < page.Blocks.length) {
				blocks += ",";
			}
		}
	}
	blocksTd.innerHTML = blocks
	tr.appendChild(blocksTd);
	tr.setAttribute("onclick","module.editPageBlock('" + page.Url + "'); return false;" );
		
	return tr;
};

module.editPageBlock = function(pageUrl) {
	if (module.currentModule.Pages) {
		for (var ii =0; ii < module.currentModule.Pages.length; ++ii) {
			var page = module.currentModule.Pages[ii];
			if (page.Url == pageUrl) {
				module.currentPage = page;
				break;
			}
		}
	}
	
	$("#module-maintain .page .page-Form .page-url").val(pageUrl);
	$("#module-maintain .page .page-Form .page-block input").prop("checked", false);
	if (module.currentPage) {
		if (module.currentPage.Blocks) {
			for (var jj =0; jj < module.currentPage.Blocks.length; ++jj) {
				var block = module.currentPage.Blocks[jj];
				$("#module-maintain .page .page-Form .page-block input").filter("[value="+ block.Id +"]").prop("checked", true);
			}
		}
	}
};

module.deleteBlock = function(deleteUrl) {
	$.get(deleteUrl, {
	}, function(result) {
		
		$("#module-maintain div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#module-maintain div.error div").html(result.Reason);
			$("#module-maintain div.error").show();
			return
		} else {
			$("#module-maintain div.success div").html(result.Reason);
			$("#module-maintain div.success").show();
			
			module.resetStatus();
    		module.currentModule = result.Module;
    		module.refreshModuleView();			
		}
	}, "json");	
};

module.refreshModuleView = function() {
	$("#module-Maintain .block table tbody tr").remove();
	$("#module-Maintain .page .page-Form .page-block").children().remove();
	$("#module-Maintain .page table tbody tr").remove();
	
	if (!module.currentModule) {
		return;
	}
	
	if (module.currentModule.Blocks) {
		for (var ii =0; ii < module.currentModule.Blocks.length; ++ii) {
			var block = module.currentModule.Blocks[ii];
			var trContent = module.constructBlockItem(block);
			$("#module-maintain .block table tbody").append(trContent);
			
			$("#module-maintain .page .page-Form .page-block").append("<input type='checkbox' name='page-block' value=" +  block.Id + "> </input> <span>" + block.Name + "</span> ");
		}			
	}
	$("#module-maintain .block table tbody tr:even").addClass("alt-row");
	
	$("#module-maintain .block .block-Form input").prop("checked", false);
	$("#module-maintain .block .block-Form input").filter("[value=0]").prop("checked", true);
	$("#module-maintain .block .block-Form .module-id").val(module.currentModule.Id);
	
	if (module.currentModule.Pages) {
		for (var ii =0; ii < module.currentModule.Pages.length; ++ii) {
			var page = module.currentModule.Pages[ii];
			var trContent = module.constructPageItem(page);
			$("#module-maintain .page table tbody").append(trContent);
		}			
	}
	$("#module-maintain .page table tbody tr:even").addClass("alt-row");
	
	$("#module-maintain .page .page-Form .page-owner").val(module.currentModule.Id);
	if (module.currentPage) {
		if (module.currentPage.Blocks) {
			for (var jj =0; jj < module.currentPage.Blocks.length; ++jj) {
				var block = module.currentPage.Blocks[jj];
				$("#module-maintain .page .page-Form .page-block input").filter("[value="+ block.Id +"]").prop("checked", true);
			}
		}
	}
	
};

module.refreshPageView = function() {	
	if (module.pageList && module.currentPage) {
		for (var ii =0; ii < module.pageList.length; ++ii) {
			var page = module.pageList[ii];
			if (page.Url == module.currentPage) {
				if (page.Blocks) {
					for (var jj =0; jj < page.Blocks.length; ++jj) {
						var block = page.Blocks[jj];
						$("#module-maintain .page .page-Form .page-block input").filter("[value="+ block +"]").prop("checked", true);
					}
				}
				
				break;
			}
		}
	}	
};

module.resetStatus = function() {
	module.currentModule = null;
	module.currentPage = null;
	
	$("#module-Maintain .page .page-Form .page-url").val('');
	$("#module-Maintain .page .page-Form .page-block input").prop("checked", false);	
};




