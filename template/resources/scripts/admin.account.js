
var account = {
		user :{
			userInfo :{}
		},
		group :{
			groupInfo :{}
		}
};

account.initialize = function(accessCode, view) {
	account.accessCode = accessCode;
	
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
    		$("#user-Edit div.notification").hide();
    		
        	if (result.ErrCode > 0) {
        		$("#user-Edit div.error div").html(result.Reason);
        		$("#user-Edit div.error").show();
        	} else {
        		$("#user-Edit div.success div").html(result.Reason);
        		$("#user-Edit div.success").show();
        		
        		account.refreshUser();
        	}
        	
        	return false;
        }
        
        function validate() {
        	var result = true
        	
        	$("#user-content .user-Form .user-account").parent().find("span").remove();
        	var account = $("#user-content .user-Form .user-account").val();
        	if (account.length == 0) {
        		$("#user-content .user-Form .user-account").parent().append("<span class=\"input-notification error png_bg\">请输入账号</span>");
        		result = false;
        	}
        	
        	var id = $("#user-content .user-Form .user-id").val();
        	if (id == -1) {
            	$("#user-content .user-Form .user-password").parent().find("span").remove();
            	var password = $("#user-content .user-Form .user-password").val();
            	if (password.length == 0) {
            		$("#user-content .user-Form .user-password").parent().append("<span class=\"input-notification error png_bg\">请输入密码</span>");
            		result = false;
            	}
            	
            	$("#user-content .user-Form .user-repassword").parent().find("span").remove();
            	var repassword = $("#user-content .user-Form .user-repassword").val();
            	if (repassword.length == 0) {
            		$("#user-content .user-Form .user-repassword").parent().append("<span class=\"input-notification error png_bg\">请输入确认密码</span>");
            		result = false;
            	}
            	if (password != repassword && repassword.length > 0) {
            		$("#user-content .user-Form .user-repassword").parent().append("<span class=\"input-notification error png_bg\">两次密码不一致</span>");
            		result = false;        		
            	}
        	}
        	
        	$("#user-content .user-Form .user-nickname").parent().find("span").remove();
        	var nickname = $("#user-content .user-Form .user-nickname").val();
        	if (nickname.length == 0) {
        		$("#user-content .user-Form .user-nickname").parent().append("<span class=\"input-notification error png_bg\">请输入昵称</span>");
        		result = false;
        	}
        	
        	$("#user-content .user-Form .user-email").parent().find("span").remove();
        	var email = $("#user-content .user-Form .user-email").val();
        	if (email.length == 0) {
        		$("#user-content .user-Form .user-email").parent().append("<span class=\"input-notification error png_bg\">请输入合法的邮箱</span>");
        		result = false;
        	} else if (email.search(/^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$/) == -1) {
        		$("#user-content .user-Form .user-email").parent().append("<span class=\"input-notification error png_bg\">请输入合法的邮箱</span>");
        		result = false;        		
        	}
        	
        	$("#user-content .user-Form .user-group").parent().find("span").remove();
        	var group = $("#user-content .user-Form .user-group").val();
        	if (group == -1) {
        		$("#user-content .user-Form .user-group").parent().append("<span class=\"input-notification error png_bg\">请选择分组</span>");
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
        	$("#group-Edit div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#group-Edit div.error div").html(result.Reason);
        		$("#group-Edit div.error").show();
        	} else {
        		$("#group-Edit div.success div").html(result.Reason);
        		$("#group-Edit div.success").show();
        		
        		account.refreshGroup();        		
        	}
        	
        	return false;
        }
        
        function validate() {
        	var result = true
        	
        	$("#group-content .group-Form .group-name").parent().find("span").remove();
        	var name = $("#group-content .group-Form .group-name").val();
        	if (name.length == 0) {
        		$("#group-content .group-Form .group-name").parent().append("<span class=\"input-notification error png_bg\">请输入分组名</span>");
        		result = false;
        	}
        	
        	$("#group-content .group-Form .group-parent").parent().find("span").remove();
        	var group = $("#group-content .group-Form .group-parent").val();
        	if (group == -1) {
        		$("#group-content .group-Form .group-parent").parent().append("<span class=\"input-notification error png_bg\">请选择父分组</span>");
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

account.refreshUser = function() {
	$.post("/account/admin/queryAllUser/", {
		accesscode: account.accessCode
	}, function(result){
		account.errCode = result.ErrCode;
		account.reason = result.Reason;
		
		account.user.userInfo = result.User;		
		
		account.fillUserView();
		
		return false;
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
		
		return false;
	}, "json");	
}

account.fillUserView = function() {
	$("#user-List div.notification").hide()
	
	if (account.errCode > 0) {
		$("#user-List table").hide()
		
		$("#user-List div.error div").html(account.reason);
		$("#user-List div.error").show();
		return;
	}
	
	$("#user-List table tbody tr").remove();
	var userInfoList = account.user.userInfo;
	for (var ii =0; ii < userInfoList.length; ++ii) {
		var userInfo = userInfoList[ii];
		var trContent = account.constructUserItem(userInfo);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		$("#user-List table tbody").append(trContent);
	}
	
	$("#user-Edit div.notification").hide()
	$("#user-Edit .user-Form .user-id").val(-1);
	$("#user-Edit .user-Form .user-account").val("");
	
	$("#user-Edit .user-Form .user-password").val("");
	$("#user-Edit .user-Form .user-repassword").val("");
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
	accountLink.setAttribute("onclick","account.editUser('/account/admin/queryUser/?id=" + userInfo.Id + "'); return false;" );
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
	editLink.setAttribute("onclick","account.editUser('/account/admin/queryUser/?id=" + userInfo.Id + "'); return false;" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteUser" );
	deleteLink.setAttribute("onclick","account.deleteUser('/account/admin/deleteUser/?id=" + userInfo.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
}

account.fillGroupView = function() {
	$("#group-List div.notification").hide()
	if (account.errCode > 0) {
		$("#group-List table").hide()
		
		$("#group-List div.error div").html(account.reason);
		$("#group-List div.error").show();
		return;
	}
	
	$("#group-List table tbody tr").remove();
	var groupList = account.group.groupInfo;
	for (var ii =0; ii < groupList.length; ++ii) {
		var group = groupList[ii];
		var trContent = account.constructGroupItem(group);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		$("#group-List table tbody").append(trContent);
	}
	
	$("#group-Edit div.notification").hide()
	$("#group-Edit .group-Form .group-id").val(-1);
	$("#group-Edit .group-Form .group-name").val("");
	$("#group-Edit .group-Form .group-parent").empty();		
	$("#group-Edit .group-Form .group-parent").append("<option value=-1>请选择父分组</option>");
	$("#group-Edit .group-Form .group-parent").append("<option value=0>无父分组</option>");
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
	nameLink.setAttribute("onclick","account.editGroup('/account/admin/queryGroup/?id=" + group.Id + "'); return false;" );
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
	editLink.setAttribute("onclick","account.editGroup('/account/admin/queryGroup/?id=" + group.Id + "'); return false;" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteGroup" );
	deleteLink.setAttribute("onclick","account.deleteGroup('/account/admin/deleteGroup/?id=" + group.Id + "'); return false;" );
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
		$("#user-List div.notification").hide()
		if (result.ErrCode > 0) {
			$("#user-List div.error div").html(result.Reason);
			$("#user-List div.error").show();
			return
		}
		
		$("#user-Edit .user-Form .user-id").val(result.User.Id);
		$("#user-Edit .user-Form .user-account").val(result.User.Account);
		$("#user-Edit .user-Form .user-nickname").val(result.User.NickName);
		$("#user-Edit .user-Form .user-email").val(result.User.Email);
		
		$("#user-Edit .user-Form .user-password").parent().hide();
		$("#user-Edit .user-Form .user-repassword").parent().hide();
		
		$("#user-Edit .user-Form .user-group").empty();		
		$("#user-Edit .user-Form .user-group").append("<option value=-1>请选择分组</option>");
		for (var ii =0; ii < account.group.groupInfo.length; ++ii) {
			group = account.group.groupInfo[ii];
			
			$("#user-Edit .user-Form .user-group").append("<option value="+  group.Id + ">" + group.Name + "</option>");
			if (group.Id == result.User.Group.Id) {
				$("#user-Edit .user-Form .user-group").get(0).selectedIndex = ii +1;				
			}
		}

		$("#user-content .content-box-tabs li a").removeClass('current');
		$("#user-content .content-box-tabs li a.user-Edit-tab").addClass('current');
		$("#user-Edit").siblings().hide();
		$("#user-Edit").show();
	}, "json");
}

account.deleteUser = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: account.accessCode
	}, function(result) {
		$("#user-List div.notification").hide()
		if (result.ErrCode > 0) {
			$("#user-List div.error div").html(result.Reason);
			$("#user-List div.error").show();
			return
		}
		
		$("#user-List div.success div").html(result.Reason);
		$("#user-List div.success").show();

		account.refreshUser();
}, "json");
}

account.editGroup = function(editUrl) {
	$.post(editUrl, {
		accesscode: account.accessCode
	}, function(result) {
		$("#group-List div.notification").hide()
		
		if (result.ErrCode > 0) {
			$("#group-List div.error div").html(result.Reason);
			$("#group-List div.error").show();
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
				
		$("#group-content .content-box-tabs li a").removeClass('current');
		$("#group-content .content-box-tabs li a.group-Edit-tab").addClass('current');
		$("#group-Edit").siblings().hide();
		$("#group-Edit").show();		
	}, "json");
}

account.deleteGroup = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: account.accessCode
	}, function(result) {
		$("#group-List div.notification").hide()
		
		if (result.ErrCode > 0) {
			$("#group-List div.error div").html(result.Reason);
			$("#group-List div.error").show();
			return false;
		}
		
		$("#group-List div.success div").html(result.Reason);
		$("#group-List div.success").show();
		
		account.refreshGroup();		
		return false;
	}, "json");
}





