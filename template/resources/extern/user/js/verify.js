    var verify = {};


    $(document).ready(function() {

        // 绑定表单提交事件处理器
        $('#user-Form').submit(function() {
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
                $("#user-Form div.notification").hide();
                if (result.ErrCode > 0) {
                    $("#user-Form div.error div").html(result.Reason);
                    $("#user-Form div.error").show();
                } else {
                    $("#user-Form div.success div").html(result.Reason);
                    $("#user-Form div.success").show();

                    location.href = "/account/userProfile/";
                }
            }

            function validate() {
                var result = true

                $("#user-Form .user-nickname").parent().find("span").remove();
                var name = $("#user-Form .user-nickname").val();
                if (name.length == 0) {
                    $("#user-Form .user-nickname").parent().append("<span class=\"input-notification error png_bg\">请输入昵称</span>");
                    result = false;
                    return result;
                }

                $("#user-Form .user-password").parent().find("span").remove();
                var password = $("#user-Form .user-password").val();
                if (password.length == 0) {
                    $("#user-Form .user-password").parent().append("<span class=\"input-notification error png_bg\">请输入密码</span>");
                    result = false;
                    return result;
                }

                $("#user-Form .user-repassword").parent().find("span").remove();
                var repassword = $("#user-Form .user-repassword").val();
                if (repassword.length == 0) {
                    $("#user-Form .user-repassword").parent().append("<span class=\"input-notification error png_bg\">请输入密码</span>");
                    result = false;
                    return result;
                }
                if (repassword != password) {
                    $("#user-Form .user-repassword").parent().append("<span class=\"input-notification error png_bg\">两次输入的密码不一致</span>");
                    result = false;
                    return result;
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