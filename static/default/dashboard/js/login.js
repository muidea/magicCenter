$(document).ready(function() {
    var options = {
        beforeSubmit: showRequest, // pre-submit callback
        success: showResponse, // post-submit callback
        dataType: 'json' // 'xml', 'script', or 'json' (expected server response type) 
    };

    $('#login_form').attr("action", "/cas/user/");

    // 绑定表单提交事件处理器
    $('#login_form').submit(function() {
        //提交表单
        $(this).ajaxSubmit(options);

        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });
    // pre-submit callback
    function showRequest() {

    }
    // post-submit callback
    function showResponse(re) {
        console.log(re);
        if (re.ErrCode > 0) {
            $("#alertInfo").removeClass('hidden');
            $("#alertInfo").html(re.Reason);
        } else {
            location.href = "/static/dashboard/index.html";
        }
    }
});