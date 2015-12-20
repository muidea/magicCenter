

var group = {
	accesscode:'',
	errCode:0,
	reason:'',
	groupInfo:{}
};

group.initialize = function() {
	
	group.fillGroupView();
		
    // 绑定表单提交事件处理器
    $('#group-content .content-box-content .content-edit .group-Form').submit(function() {
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
        	$("#group-content .content-box-content .content-edit div.notification").hide();
        	if (result.ErrCode > 0) {
        		$("#group-content .content-box-content .content-edit div.error div").html(result.Reason);
        		$("#group-content .content-box-content .content-edit div.error").show();        		
        	} else {
        		$("#group-content .content-box-content .content-edit div.success div").html(result.Reason);
        		$("#group-content .content-box-content .content-edit div.success").show();        		
        		group.refreshGroup();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#group-content .content-box-content .content-edit .group-Form .group-name").parent().find("span").remove();
        	var name = $("#group-content .content-box-content .content-edit .group-Form .group-name").val();
        	if (name.length == 0) {
        		$("#group-content .content-box-content .content-edit .group-Form .group-name").parent().append("<span class=\"input-notification error png_bg\">请输入分组名</span>");
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
};

group.refreshGroup = function() {
	$.post("/admin/account/queryAllGroup/", {
		accesscode: group.accessCode
	}, function(result){
		article.errCode = result.ErrCode;
		article.reason = result.Reason;
		
		group.groupInfo = result.Group;		
		
		group.fillGroupView();
	}, "json");	
};


group.fillGroupView = function() {
	
	if ( group.errCode > 0) {
		return;
	}
	
	$("#group-content .content-box-content .content-edit div.notification").hide();
	$("#group-content .content-box-content .content-edit .group-Form .group-id").val("-1");
	$("#group-content .content-box-content .content-edit .group-Form .group-name").val("");
	
	
	$("#group-content .content-list table tbody tr").remove();
	for (var ii =0; ii < group.groupInfo.length; ++ii) {
		var info = group.groupInfo[ii];
		var trContent = group.constructGroupItem(info);
		$("#group-content .content-list table tbody").append(trContent);
	}
	$("#group-content .content-list table tbody tr:even").addClass("alt-row");
	$("#group-content .content-list table").show();	
};


group.constructGroupItem = function(group) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","group");
	
	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("class","edit");
	nameLink.setAttribute("href","#editGroup" );
	nameLink.innerHTML = group.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);

	var userCountlTd = document.createElement("td");
	userCountlTd.innerHTML = group.UserCount;
	tr.appendChild(userCountlTd);
	
	var editTd = document.createElement("td");
	if (group.Catalog > 0) {
		var deleteLink = document.createElement("a");
		deleteLink.setAttribute("class","delete");
		deleteLink.setAttribute("href","#deleteGroup" );
		deleteLink.setAttribute("onclick","group.deleteGroup('/admin/account/deleteGroup/?id=" + group.Id + "'); return false;" );
		var deleteUrl = document.createElement("img");
		deleteUrl.setAttribute("src","/resources/images/icons/cross.png");
		deleteUrl.setAttribute("alt","Delete");
		deleteLink.appendChild(deleteUrl);
		editTd.appendChild(deleteLink);
	} else {
		editTd.innerHTML = "-";
	}
	
	tr.appendChild(editTd);
	tr.setAttribute("onclick","group.editGroup('/admin/account/editGroup/?id=" + group.Id + "'); return false;" );
	
	return tr;
};

group.editGroup = function(editUrl) {
	$.post(editUrl, {
		accesscode: group.accessCode
	}, function(result) {
    	$("#group-content .content-box-content .content-edit div.notification").hide();
    	if (result.ErrCode > 0) {
    		$("#group-content .content-box-content .content-edit div.error div").html(result.Reason);
    		$("#group-content .content-box-content .content-edit div.error").show();        		
    	} else {
    		$("#group-content .content-box-content .content-edit .group-Form .group-id").val(result.Id);
    		$("#group-content .content-box-content .content-edit .group-Form .group-name").val(result.Name);
    	}
	}, "json");
};

group.deleteGroup = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: group.accessCode
	}, function(result) {
    	$("#group-content .content-box-content .content-edit div.notification").hide();
    	if (result.ErrCode > 0) {
    		$("#group-content .content-box-content .content-edit div.error div").html(result.Reason);
    		$("#group-content .content-box-content .content-edit div.error").show();        		
    	} else {
    		$("#group-content .content-box-content .content-edit div.success div").html(result.Reason);
    		$("#group-content .content-box-content .content-edit div.success").show();        		
    		group.refreshGroup();
    	}		

	}, "json");
};


