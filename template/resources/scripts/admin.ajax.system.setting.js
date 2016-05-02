

var setting = {
	errCode:0,
	reason:'',
	name:'',
	logo:'',
	domain:'',
	emailServer:'',
	emailAccount:'',
	emailPassword:''
};

setting.initialize = function() {
	
	setting.fillSettingView();
		
    // 绑定表单提交事件处理器
    $('#system-content .content-box-content .system-edit .system-Form').submit(function() {
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
        	$("#system-content .content-box-content .system-edit div.notification").hide();
        	if (result.ErrCode > 0) {
        		$("#system-content .content-box-content .system-edit div.error div").html(result.Reason);
        		$("#system-content .content-box-content .system-edit div.error").show();        		
        	} else {
        		$("#system-content .content-box-content .system-edit div.success div").html(result.Reason);
        		$("#system-content .content-box-content .system-edit div.success").show();        		
        	}
        }
        
        function validate() {
        	var result = true
        	
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

setting.fillSettingView = function() {
	
	if ( setting.errCode > 0) {
		return;
	}
	
	$("#system-content .content-box-content .system-edit div.notification").hide();
	$("#system-content .content-box-content .system-edit .system-Form .system-name").val(setting.name);
	$("#system-content .content-box-content .system-edit .system-Form .system-logo").val(setting.logo);
	$("#system-content .content-box-content .system-edit .system-Form .system-domain").val(setting.domain);
	$("#system-content .content-box-content .system-edit .system-Form .email-server").val(setting.emailServer);
	$("#system-content .content-box-content .system-edit .system-Form .email-account").val(setting.emailAccount);
	$("#system-content .content-box-content .system-edit .system-Form .email-password").val(setting.emailPassword);
};


