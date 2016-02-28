
var module = {
	accesscode:'',
	moduleList:{}
};

$(document).ready(function() {
	
	$("#module-list .button").click(
		function() {
			var enableList = "";
			var disableList = "";
			var defaultModule = "";
			var radioArray = $("#module-list table tbody tr td :radio:checked");
			for (var ii =0; ii < radioArray.length; ++ii) {
				var radio = radioArray[ii];
				if ($(radio).val() == 1) {
					enableList += $(radio).attr("name");
					enableList += ",";
				} else {
					disableList += $(radio).attr("name");
					disableList += ",";
				}
				
			}
			
			var checkboxArray = $("#module-list table tbody tr td :checkbox:checked");
			for (var ii =0; ii < checkboxArray.length;) {
				var checkbox = checkboxArray[ii++];
				defaultModule += $(checkbox).attr("name");
				if (ii < checkboxArray.length) {
					defaultModule += ",";
				}
			}
			
			$.post("/admin/system/applyModule/", {
				enableList:enableList,
				disableList:disableList,
				defaultModule:defaultModule
			}, function(result) {

				$("#module-list div.notification").hide();
	        	if (result.ErrCode > 0) {
	        		$("#module-list div.error div").html(result.Reason);
	        		$("#module-list div.error").show();
	        	} else {
	        		$("#module-list div.success div").html(result.Reason);
	        		$("#module-list div.success").show();
	        	}				
				
			}, "json");
		}
	);
	
    // 绑定表单提交事件处理器
    $('#module-maintain .block .block-Form').submit(function() {
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
        		
				var trContent = module.constructBlockItem(result.Owner, result.Block);
				$("#module-maintain .block table tbody").append(trContent);
				$("#module-maintain .block table tbody tr:even").addClass("alt-row");
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#module-maintain .block .block-Form .module-block").parent().find("span").remove();
        	var name = $("#module-maintain .block .block-Form .module-block").val();
        	if (name.length == 0) {
        		$("#module-maintain .block .block-Form .module-block").parent().append("<span class=\"input-notification error png_bg\">请输入功能块名</span>");
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

module.initialize = function() {
	module.fillModuleView();
};

module.fillModuleView = function() {
	
	$("#module-list div.notification").hide();
	$("#module-list table tbody tr").remove();
	for (var ii =0; ii < module.moduleList.length; ++ii) {
		var info = module.moduleList[ii];
		var trContent = module.constructModuleItem(info);		
		$("#module-list table tbody").append(trContent);
	}
	
	$("#module-list table tbody tr:even").addClass("alt-row");	
};

module.constructModuleItem = function(module) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","module");
	
	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("class","view");
	nameLink.setAttribute("href","#");
	nameLink.setAttribute("onclick","module.maintainModule('/admin/system/queryModuleInfo/?id=" + module.Id + "'); return false;" );
	nameLink.innerHTML = module.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);

	var descriptionTd = document.createElement("td");
	descriptionTd.innerHTML = module.Description
	tr.appendChild(descriptionTd);
	
	var editTd = document.createElement("td");
	var radioGroup = document.createElement("radiobox");
	var enable_radio = document.createElement("input");
	enable_radio.setAttribute("type","radio");
	enable_radio.setAttribute("name",module.Id);
	enable_radio.setAttribute("value","1");
	radioGroup.appendChild(enable_radio);	
	var enable_span = document.createElement("span");
	enable_span.innerHTML ="启用";
	radioGroup.appendChild(enable_span);
	
	var disable_radio = document.createElement("input");
	disable_radio.setAttribute("type","radio");
	disable_radio.setAttribute("name",module.Id);
	disable_radio.setAttribute("value","0");
	radioGroup.appendChild(disable_radio);
	if (module.Enable) {
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
	default_check.setAttribute("name",module.Id);
	checkGroup.appendChild(default_check);
	if (module.Default) {
		default_check.checked = true;
	} else {
		default_check.checked = false;
	}
	
	
	var default_span = document.createElement("span");
	default_span.innerHTML ="设为默认 ";
	checkGroup.appendChild(default_span);
	
	editTd.appendChild(checkGroup);	
	
	if(module.Internal == 0) {
		var uninstall = document.createElement("input");
		uninstall.setAttribute("type","button");
		uninstall.setAttribute("class","button");
		uninstall.setAttribute("value","卸载模块");
		editTd.appendChild(uninstall);
	}
	
	tr.appendChild(editTd);	
	return tr;
};

module.maintainModule = function(maintainUrl) {
	$.get(maintainUrl, {
	}, function(result) {
		$("#module-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#module-List div.error div").html(result.Reason);
			$("#module-List div.error").show();
			return
		}
		
		$("#module-content .content-box-tabs li a").removeClass('current');
		$("#module-content .content-box-tabs li a.module-Maintain-tab").addClass('current');
		$("#module-maintain").siblings().hide();
		$("#module-maintain .block table tbody tr").remove();
		if (result.Blocks) {
			for (var ii =0; ii < result.Blocks.length; ++ii) {
				var block = result.Blocks[ii];
				var trContent = module.constructBlockItem(result.Module.Id, block);
				$("#module-maintain .block table tbody").append(trContent);
			}			
		}
		$("#module-maintain .block table tbody tr:even").addClass("alt-row");
		$("#module-maintain .block .block-Form .module-id").val(result.Module.Id);

		console.log(result);
		$("#module-maintain .page table tbody tr").remove();
		if (result.Pages) {
			for (var ii =0; ii < result.Pages.length; ++ii) {
				var page = result.Pages[ii];
				var trContent = module.constructPageItem(result.Module.Id, page);
				$("#module-maintain .page table tbody").append(trContent);
			}			
		}
		$("#module-maintain .page table tbody tr:even").addClass("alt-row");
		
		$("#module-maintain").show();	
	}, "json");	
};

module.constructBlockItem = function(owner, block) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","block");
	
	var nameTd = document.createElement("td");
	nameTd.innerHTML = block.Name
	tr.appendChild(nameTd);
	
	var editTd = document.createElement("td");
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteBlock" );
	deleteLink.setAttribute("onclick","module.deleteBlock('/admin/system/deleteBlock/?id=" + block.Id + "&owner="+ owner +"'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	tr.appendChild(editTd);	
	return tr;
};

module.constructPageItem = function(owner, page) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","block");
	
	var nameTd = document.createElement("td");
	nameTd.innerHTML = page.Url
	tr.appendChild(nameTd);
	
	var blocks = "";
	var blocksTd = document.createElement("td");
	if (page.Blocks) {
		for (var ii =0; ii < page.Blocks.length; ++ii) {
			blocks += page.Blocks[ii];
			blocks += ","
		}		
	}
	blocksTd.innerHTML = blocks
	tr.appendChild(blocksTd);	
	return tr;
};

module.deleteBlock =function(deleteUrl) {
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
			
			$("#module-maintain .block table tbody tr").remove();
			if (result.Blocks) {
				for (var ii =0; ii < result.Blocks.length; ++ii) {
					var info = result.Blocks[ii];
					var trContent = module.constructBlockItem(result.Owner, info);
					$("#module-maintain .block table tbody").append(trContent);
				}			
			}
			$("#module-maintain .block table tbody tr:even").addClass("alt-row");			
		}
	}, "json");	
};

