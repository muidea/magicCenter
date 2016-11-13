var maintain = {};

$(document).ready(function() {

    // 绑定表单提交事件处理器
    $("#content .blog-Form").submit(function() {
        var options = {
            beforeSubmit: showRequest,
            success: showResponse,
            dataType: "json"
        };

        function showRequest() {}

        function showResponse(result) {
            if (result.ErrCode > 0) {
                $("#content .alert-Info .content").html(result.Reason);
                $("#content .alert-Info").modal();
            } else {
                maintain.title = $("#content .blog-Form .blog-title").val();
                maintain.description = $("#content .blog-Form .blog-description").val();

                maintain.fillMaintain();
            }
        }

        function validate() {
            var result = true;
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

    maintain.fillMaintain();
});

maintain.initialize = function() {

};

maintain.refreshMaintain = function() {

};

maintain.fillMaintain = function() {
    $("#content .blog-Form .blog-title").val(maintain.title);
    $("#content .blog-Form .blog-description").val(maintain.description);
};