
var account = {
		user :{
			view :{},
			userInfo :{}
		},
		group :{
			view :{},
			groupInfo :{}
		}
};

account.initialize = function(accessCode, view) {
	var userView = view.find("#user-content");
	var groupView = view.find("#group-content");
	
	account.accessCode = accessCode;
	account.user.view = userView;		
	account.group.view = groupView;
	
	$.post("/account/admin/queryAllInfo/", {
		accesscode: accessCode
	}, function(result){
		account.errCode = result.ErrCode;
		account.reason = result.Reason;
		
		account.user.userInfo = result.User;		
		account.group.groupInfo = result.Group;
		
		account.fillUserView();
		account.fillGroupView();
		
	}, "json");
	
    // 绑定表单提交事件处理器
    $('#user-content .user-Form').submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //alert("before submit");
        } 
        // post-submit callback
        function showResponse(result) {
        	if (result.ErrCode > 0) {
        		var notificationDiv = $(account.user.view).find("#user-Edit div.error");
        		$(notificationDiv).children("div").html(result.Reason);
        		$(notificationDiv).siblings(".notification").hide();
        		$(notificationDiv).show();
        	} else {
        		var notificationDiv = $(account.user.view).find("#user-Edit div.success");
        		$(notificationDiv).children("div").html(result.Reason);
        		$(notificationDiv).siblings(".notification").hide();
        		$(notificationDiv).show();
        		
        		account.refreshUser();
        	}
        }
        //提交表单
        $(this).ajaxSubmit(options);	
    	
        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });
    
    // 绑定表单提交事件处理器
    $('#group-content .group-Form').submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //alert("before submit");
        } 
        // post-submit callback
        function showResponse(result) {
        	if (result.ErrCode > 0) {
        		var notificationDiv = $(account.group.view).find("#group-Edit div.error");
        		$(notificationDiv).children("div").html(result.Reason);
        		$(notificationDiv).siblings(".notification").hide();
        		$(notificationDiv).show();
        	} else {
        		var notificationDiv = $(account.group.view).find("#group-Edit div.success");
        		$(notificationDiv).children("div").html(result.Reason);
        		$(notificationDiv).siblings(".notification").hide();
        		$(notificationDiv).show();
        		
        		account.refreshGroup();
        	}
        }
        //提交表单
        $(this).ajaxSubmit(options);	
    	
        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });	    
};

account.refreshUser = function() {
	$.post("/account/admin/queryAllUser/", {
		accesscode: account.accessCode
	}, function(result){
		account.errCode = result.ErrCode;
		account.reason = result.Reason;
		
		account.user.userInfo = result.User;		
		
		account.fillUserView();
	}, "json");	
}

account.refreshGroup = function() {
	$.post("/account/admin/queryAllGroup/", {
		accesscode: account.accessCode
	}, function(result){
		account.errCode = result.ErrCode;
		account.reason = result.Reason;
		
		account.group.groupInfo = result.Group;		
		
		account.fillGroupView();		
	}, "json");	
}

account.fillUserView = function() {
	var userTable = account.user.view.find("#user-List").children("table");
	var notificationDiv = account.user.view.find("#user-List").children("div");
	if (account.errCode > 0) {
		$(notificationDiv).children("div").html(result.Reason);
		$(notificationDiv).siblings(".notification").hide();
		$(notificationDiv).show();
		
		userTable.hide();
		return;
	}
	
	$("#user-List .notification").hide();
	var userListBody = userTable.children("tbody");
	userListBody.children("tr").remove();
	
	var userInfoList = account.user.userInfo;
	for (var ii =0; ii < userInfoList.length; ++ii) {
		var userInfo = userInfoList[ii];
		var trContent = account.constructUserItem(userInfo);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		userListBody.append(trContent);
	}
	
	$("#user-Edit .user-Form .user-id").val(-1);
	$("#user-Edit .user-Form .user-account").val("");
	$("#user-Edit .user-Form .user-password").val("");
	$("#user-Edit .user-Form .user-nickname").val("");
	$("#user-Edit .user-Form .user-email").val("");
	$("#user-Edit .user-Form .user-group").empty();	
	$("#user-Edit .user-Form .user-group").append("<option value=-1>选择分组</option>");
	for (var ii =0; ii < account.group.groupInfo.length; ++ii) {
		group = account.group.groupInfo[ii];
		
		$("#user-Edit .user-Form .user-group").append("<option value="+  group.Id + ">" + group.Name + "</option>");
	}	
}

account.constructUserItem = function(userInfo) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","user");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);

	var accountTd = document.createElement("td");
	var accountLink = document.createElement("a");
	accountLink.setAttribute("class","edit");
	accountLink.setAttribute("href","#queryUser");
	accountLink.setAttribute("onclick","account.editUser('/account/admin/queryUser/?id=" + userInfo.Id + "')" );
	accountLink.innerHTML = userInfo.Account;
	accountTd.appendChild(accountLink);
	tr.appendChild(accountTd);

	var nickNameTd = document.createElement("td");
	nickNameTd.innerHTML = userInfo.NickName;
	tr.appendChild(nickNameTd);

	var eMailTd = document.createElement("td");
	eMailTd.innerHTML = userInfo.Email;
	tr.appendChild(eMailTd);
	
	var groupTd = document.createElement("td");
	groupTd.innerHTML = userInfo.Group.Name;
	tr.appendChild(groupTd);

	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#queryUser");
	editLink.setAttribute("onclick","account.editUser('/account/admin/queryUser/?id=" + userInfo.Id + "')" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteUser" );
	deleteLink.setAttribute("onclick","account.deleteUser('/account/admin/deleteUser/?id=" + userInfo.Id + "')" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
}

account.fillGroupView = function() {
	var groupTable = account.group.view.find("#group-List").children("table");
	var notificationDiv = account.group.view.find("#group-List").children("div");
	if (account.errCode > 0) {
		$(notificationDiv).children("div").html(result.Reason);
		$(notificationDiv).siblings(".notification").hide();
		$(notificationDiv).show();
		
		groupTable.hide();
		return;
	}
	
	$("#group-List .notification").hide();
	var groupListBody = groupTable.children("tbody");
	groupListBody.children("tr").remove();
	
	var groupList = account.group.groupInfo;
	for (var ii =0; ii < groupList.length; ++ii) {
		var group = groupList[ii];
		var trContent = account.constructGroupItem(group);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		groupListBody.append(trContent);
	}
	
	$("#group-Edit .group-Form .group-id").val(-1);
	$("#group-Edit .group-Form .group-name").val("");
	$("#group-Edit .group-Form .group-parent").empty();		
	$("#group-Edit .group-Form .group-parent").append("<option value=-1>请选择父类</option>");
	$("#group-Edit .group-Form .group-parent").append("<option value=0>无父类</option>");
	for (var ii =0; ii < account.group.groupInfo.length; ++ii) {
		group = account.group.groupInfo[ii];
		
		$("#group-Edit .group-Form .group-parent").append("<option value="+  group.Id + ">" + group.Name + "</option>");				
	}
};


account.constructGroupItem = function(group) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","group");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);

	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("class","edit");
	nameLink.setAttribute("href","#editGroup" );
	nameLink.setAttribute("onclick","account.editGroup('/account/admin/queryGroup/?id=" + group.Id + "')" );
	nameLink.innerHTML = group.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);

	var parentName = "-";
	for (var ii = 0; ii < account.group.groupInfo.length; ++ii) {
		var g = account.group.groupInfo[ii]; 
		if (g.Id == group.Catalog) {
			parentName = g.Name;
			break;
		}
	}
	var parentTd = document.createElement("td");
	parentTd.innerHTML = parentName;
	tr.appendChild(parentTd);
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editGroup" );
	editLink.setAttribute("onclick","account.editGroup('/account/admin/queryGroup/?id=" + group.Id + "')" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteGroup" );
	deleteLink.setAttribute("onclick","account.deleteGroup('/account/admin/deleteGroup/?id=" + group.Id + "')" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};

account.editUser = function(editUrl) {
	$.post(editUrl, {
		accesscode: account.accessCode
	}, function(result) {
		if (result.ErrCode > 0) {
			$("#user-List div.error div").html(result.Reason);
			$("#user-List .notification").show();
			return
		}
		
		$("#user-Edit .user-Form .user-id").val(result.User.Id);
		$("#user-Edit .user-Form .user-account").val(result.User.Account);
		$("#user-Edit .user-Form .user-nickname").val(result.User.NickName);
		$("#user-Edit .user-Form .user-email").val(result.User.Email);
		
		$("#user-Edit .user-Form .user-group").empty();		
		$("#user-Edit .user-Form .user-group").append("<option value=-1>请选择分组</option>");
		for (var ii =0; ii < account.group.groupInfo.length; ++ii) {
			group = account.group.groupInfo[ii];
			
			$("#user-Edit .user-Form .user-group").append("<option value="+  group.Id + ">" + group.Name + "</option>");
			if (group.Id == result.User.Group.Id) {
				$("#user-Edit .user-Form .user-group").get(0).selectedIndex = ii +1;				
			}
		}

		$(account.user.view).find(".content-box-tabs li a").removeClass('current');
		$(account.user.view).find(".content-box-tabs li a.user-Edit-tab").addClass('current');
		$("#user-Edit").siblings().hide();
		$("#user-Edit").show();
	}, "json");
}

account.deleteUser = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: account.accessCode
	}, function(result) {
		if (result.ErrCode > 0) {
			$("#user-List div.error div").html(result.Reason);
			$("#user-List .notification").show();
			return
		}
		
		account.refreshUser();
	}, "json");
}

account.editGroup = function(editUrl) {
	$.post(editUrl, {
		accesscode: account.accessCode
	}, function(result) {
		if (result.ErrCode > 0) {
			$("#group-List div.error div").html(result.Reason);
			$("#group-List .notification").show();
			return;
		}
		
		$("#group-Edit .group-Form .group-id").val(result.Group.Id);
		$("#group-Edit .group-Form .group-name").val(result.Group.Name);
		
		$("#group-Edit .group-Form .group-parent").empty();		
		$("#group-Edit .group-Form .group-parent").append("<option value=-1>请选择分组</option>");
		$("#group-Edit .group-Form .group-parent").append("<option value=0>无分组</option>");
		
		var index = 1;
		for (var ii =0; ii < account.group.groupInfo.length; ++ii) {
			group = account.group.groupInfo[ii];
			
			if (group.Id != result.Group.Id) {
				$("#group-Edit .group-Form .group-parent").append("<option value="+  group.Id + ">" + group.Name + "</option>");				
			}
			if (group.Id == result.Group.Catalog) {
				index = ii + 2;
			}
		}
		$("#group-Edit .group-Form .group-parent").get(0).selectedIndex = index;				
		
		$(account.group.view).find(".content-box-tabs li a").removeClass('current');
		$(account.group.view).find(".content-box-tabs li a.group-Edit-tab").addClass('current');
		$("#group-Edit").siblings().hide();
		$("#group-Edit").show();
	}, "json");
}

account.deleteGroup = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: account.accessCode
	}, function(result) {
		if (result.ErrCode > 0) {
			$("#group-List div.error div").html(result.Reason);
			$("#group-List .notification").show();
			return;
		}
		
		account.refreshGroup();
	}, "json");
}





