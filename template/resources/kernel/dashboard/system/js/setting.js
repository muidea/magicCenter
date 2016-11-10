var setting = {
    errCode: 0,
    reason: '',
    name: '',
    logo: '',
    domain: '',
    emailServer: '',
    emailAccount: '',
    emailPassword: ''
};


$(document).ready(function() {
    // 绑定表单提交事件处理器
    $('#system-Edit .system-Form').submit(function() {
        var options = {
            beforeSubmit: showRequest, // pre-submit callback
            success: showResponse, // post-submit callback
            dataType: 'json' // 'xml', 'script', or 'json' (expected server response type) 
        };

        // pre-submit callback
        function showRequest() {
            //return false;
        }
        // post-submit callback
        function showResponse(result) {
            if (result.ErrCode > 0) {
                $("#system-Edit .alert-Info .content").html(result.Reason);
                $("#system-Edit .alert-Info").modal();
            } else {
                // noting todo
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
});

setting.initialize = function() {

    setting.fillSettingView();

};

setting.fillSettingView = function() {
    $("#system-Edit .system-Form .system-name").val(setting.name);
    $("#system-Edit .system-Form .system-logo").val(setting.logo);
    $("#system-Edit .system-Form .system-domain").val(setting.domain);
    $("#system-Edit .system-Form .email-server").val(setting.emailServer);
    $("#system-Edit .system-Form .email-account").val(setting.emailAccount);
    $("#system-Edit .system-Form .email-password").val(setting.emailPassword);
};