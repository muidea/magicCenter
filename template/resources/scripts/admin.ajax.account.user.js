

var user = {
	accesscode:'',
	errCode:0,
	reason:'',
	userInfo:{},
	groupInfo:{}
};

user.initialize = function() {
	
	user.fillUserView();
		
	$('#user-content .user-Form .user-account').keyup(function() {
		var account = $("#user-content .user-Form .user-account").val();
		$.post("/admin/account/checkAccount/", {
			"accesscode": user.accesscode,
			"user-account": account,
		}, function(result) {
			
			$("#user-content .user-Form .user-account").parent().find("span").remove();
        	if (result.ErrCode > 0) {
        		$("#user-content .user-Form .user-account").parent().append("<span class=\"input-notification error png_bg\">" + result.Reason +"</span>");
        	}
		}, "json");		
	});
	
    // 绑定表单提交事件处理器
    $('#user-content .user-Form').submit(function() {
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
        	$("#user-Edit div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#user-Edit div.error div").html(result.Reason);
        		$("#user-Edit div.error").show();
        	} else {
        		$("#user-Edit div.success div").html(result.Reason);
        		$("#user-Edit div.success").show();

        		$("#user-Edit .user-Form .button").val("创建");
        		user.refreshUser();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#user-content .user-Form .user-account").parent().find("span").remove();
        	var name = $("#user-content .user-Form .user-account").val();
        	if (name.length == 0) {
        		$("#user-content .user-Form .user-account").parent().append("<span class=\"input-notification error png_bg\">请输入账号</span>");
        		result = false;
        	}
        	
        	var email = $("#user-content .user-Form .user-email").val();
        	if (email.length == 0) {
        		$("#user-content .user-Form .user-email").parent().append("<span class=\"input-notification error png_bg\">请输入合法的邮箱</span>");
        		result = false;
        	} else if (email.search(/^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$/) == -1) {
        		$("#user-content .user-Form .user-email").parent().append("<span class=\"input-notification error png_bg\">请输入合法的邮箱</span>");
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

user.refreshUser = function() {
	$.post("/admin/account/queryAllUser/", {
		accesscode: user.accessCode
	}, function(result){
		article.errCode = result.ErrCode;
		article.reason = result.Reason;
		
		user.userInfo = result.User;		
		
		user.fillUserView();
	}, "json");	
};


user.fillUserView = function() {
	
	$("#user-List div.notification").hide();
	
	if ( user.errCode > 0) {
		$("#user-List div.error div").html(user.reason);
		$("#user-List div.error").show();
		
		$("#user-List table").hide();
		return;
	}
	
	$("#user-List table tbody tr").remove();
	for (var ii =0; ii < user.userInfo.length; ++ii) {
		var info = user.userInfo[ii];
		var trContent = user.constructUserItem(info);
		$("#user-List table tbody").append(trContent);
	}
	$("#user-List table tbody tr:even").addClass("alt-row");
	$("#user-List table").show();
	
	$("#user-Edit div.notification").hide();		
	$("#user-Edit .user-Form .user-id").val(-1);
	$("#user-Edit .user-Form .user-account").val("");
	$("#user-Edit .user-Form .user-nickname").val("");
	$("#user-Edit .user-Form .user-email").val("");
		
	$("#user-Edit .user-Form .user-group").children().remove();
	for (var ii =0; ii < user.groupInfo.length; ++ii) {
		var ca = user.groupInfo[ii];
		$("#user-Edit .user-Form .user-group").append("<input type='checkbox' name='user-group' value=" +  ca.Id + "> </input> <span>" + ca.Name + "</span> ");
	}
	if (ii == 0) {
		$("#user-Edit .user-Form .user-group").append("<input type='checkbox' name='user-group' readonly='readonly' value='-1' onclick='return false'> </input> <span>-</span> ");
	}
};


user.constructUserItem = function(user) {
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
	accountLink.setAttribute("href","#editLink" );
	accountLink.setAttribute("onclick","user.editUser('/admin/account/editUser/?id=" + user.Id + "'); return false;" );
	accountLink.innerHTML = user.Account;
	accountTd.appendChild(accountLink);
	tr.appendChild(accountTd);

	var nickNameTd = document.createElement("td");
	nickNameTd.innerHTML = user.NickName;
	tr.appendChild(nickNameTd);

	var emailTd = document.createElement("td");
	emailTd.innerHTML = user.Email;
	tr.appendChild(emailTd);
	
	var groupTd = document.createElement("td");
	groupTd.innerHTML = user.Group;
	tr.appendChild(groupTd);
	
	var statusTd = document.createElement("td");
	switch (user.Status)
	{
	case 0:
		statusTd.innerHTML = "正常";
		break;
	default:
		break;
	}
	tr.appendChild(statusTd);
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editUser" );
	editLink.setAttribute("onclick","user.editUser('/admin/account/editUser/?id=" + user.Id + "'); return false;" );
	var editUrl = document.createElement("img");
	editUrl.setAttribute("src","/resources/images/icons/pencil.png");
	editUrl.setAttribute("alt","Edit");
	editLink.appendChild(editUrl);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteUser" );
	deleteLink.setAttribute("onclick","user.deleteUser('/admin/account/deleteUser/?id=" + user.Id + "'); return false;" );
	var deleteUrl = document.createElement("img");
	deleteUrl.setAttribute("src","/resources/images/icons/cross.png");
	deleteUrl.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteUrl);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};

user.editUser = function(editUrl) {
	$.post(editUrl, {
		accesscode: user.accessCode
	}, function(result) {
		$("#user-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#user-List div.error div").html(result.Reason);
			$("#user-List div.error").show();
			return
		}
		
		console.log(result)
		
		
		$("#user-Edit .user-Form .button").val("保存");
		$("#user-Edit .user-Form .user-id").val(result.Id);
		$("#user-Edit .user-Form .user-account").val(result.Account);
		$("#user-Edit .user-Form .user-account").prop("readonly", true);
		$("#user-Edit .user-Form .user-email").val(result.Email);
		
		$("#user-Edit .user-Form .user-group input").prop("checked", false);
		if (result.Group) {
			for (var ii =0; ii < result.Group.length; ++ii) {
				var ca = result.Group[ii];
				$("#user-Edit .user-Form .user-group input").filter("[value="+ ca +"]").prop("checked", true);			
			}			
		}
						
		$("#user-content .content-box-tabs li a").removeClass('current');
		$("#user-content .content-box-tabs li a.user-Edit-tab").addClass('current');
		$("#user-Edit").siblings().hide();
		$("#user-Edit").show();		
	}, "json");
};

user.deleteUser = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: user.accessCode
	}, function(result) {
		$("#user-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#user-List div.error div").html(result.Reason);
			$("#user-List div.error").show();
			return;
		}
		
		$("#user-List div.success div").html(result.Reason);
		$("#user-List div.success").show();
		
		user.refreshUser();
	}, "json");
};


